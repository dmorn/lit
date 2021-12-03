package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jecoz/edb"
	"github.com/jecoz/lit"
	"github.com/jecoz/lit/log"
	"github.com/jecoz/lit/scopus"
)

var (
	scopusKey = os.Getenv("SCOPUS_API_KEY")
)

var (
	edbPath = flag.String("edb", "lit.edb", "Event database file. Everything will be stored here.")
)

const (
	MaxWidth = 120
	Margin   = 1
)

type keyMap struct {
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
		{},
	}
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	db     *edb.Db
	client lit.Library
	query  string
	max    int
	next   *lit.PublicationChan

	received int
	err      error
	done     bool

	progress progress.Model
	help     help.Model
}

type publicationMsg struct {
	pub lit.Publication
}

type doneMsg struct{}

func recvPublications(c <-chan lit.Publication) tea.Cmd {
	return func() tea.Msg {
		pub, ok := <-c
		if !ok {
			return doneMsg{}
		}
		return publicationMsg{
			pub: pub,
		}
	}
}

func (m model) Init() tea.Cmd {
	return recvPublications(m.next.Chan)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.progress.Width = msg.Width - Margin*2 - 4
		if m.progress.Width > MaxWidth {
			m.progress.Width = MaxWidth
		}
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	case doneMsg:
		m.err = m.next.Err
		if m.err == nil {
			if m.received < m.max {
				m.err = fmt.Errorf("download pipeline exited prematurely")
			} else {
				m.done = true
			}
		}
		return m, nil
	case publicationMsg:
		m.received++
		data, err := msg.pub.Marshal()
		if err != nil {
			m.err = err
			return m, nil
		}
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: m.client.GetName(),
			Scope:  "lit",
			Action: "add_lit",
			Data:   []string{msg.pub.BibId(), data},
		}); err != nil {
			m.err = err
			return m, nil
		}
		return m, recvPublications(m.next.Chan)
	}
	return m, nil
}

var (
	bodyStyle    = lipgloss.NewStyle().Width(MaxWidth).Margin(Margin)
	titleStyle   = bodyStyle.Copy().Bold(true).Blink(true)
	errorStyle   = bodyStyle.Copy().Foreground(lipgloss.Color("5"))
	successStyle = bodyStyle.Copy().Foreground(lipgloss.Color("#EE6FF8")).Bold(true)
	helpStyle    = bodyStyle.Copy()
)

func (m model) View() string {
	title := fmt.Sprintf("Downloading %q (%d results) from %s...", m.query, m.max, m.client.GetName())
	titleView := titleStyle.Render(title)
	progressView := bodyStyle.Render(m.progress.ViewAs(float64(m.received) / float64(m.max)))
	helpView := helpStyle.Render(m.help.View(keys))

	var statusView string
	if m.err != nil {
		statusView = errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	} else if m.done {
		statusView = successStyle.Render("Done!")
	}

	return fmt.Sprintf("%s\n%s\n%s\n%s\n",
		titleView,
		progressView,
		statusView,
		helpView,
	)
}

func Main() error {
	db, err := edb.Open(*edbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := ""
	max := 0
	if err := db.Revive(func(e edb.Event) error {
		// TODO: we can easility recover from previous download sessions by checking:
		// 1. wether the max number of items did non change.
		// 2. if yes, start counting from the last item received.
		switch e.Action {
		case "set_query":
			var err error
			query = e.Data[0]
			max, err = strconv.Atoi(e.Data[1])
			if err != nil {
				return fmt.Errorf("parse maximum literature count: %w", err)
			}
		default:
		}
		return nil
	}); err != nil {
		return err
	}

	client := scopus.NewClient(scopusKey)
	next := lit.GetLiterature(context.Background(), client, lit.Request{Query: query})
	return tea.NewProgram(model{
		db:     db,
		client: client,
		query:  query,
		// We get an initial maximum publication count from lit-max
		// stored output. In the meanwhile, this count may have
		// changed.
		max:      next.Total,
		next:     next,
		progress: progress.NewModel(progress.WithDefaultGradient()),
		help:     help.NewModel(),
	}).Start()
}

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
