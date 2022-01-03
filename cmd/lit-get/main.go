package main

import (
	"context"
	"flag"
	"fmt"
	"os"
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
	next   *lit.BlobChan

	received int
	err      error
	done     bool

	progress progress.Model
	help     help.Model
}

type blobMsg struct {
	blob lit.Blob
}

type errMsg struct {
	err error
}

func handleBlob(blobChan *lit.BlobChan) tea.Cmd {
	return func() tea.Msg {
		blob, ok := <-blobChan.Recv()
		if !ok {
			return errMsg{blobChan.Err()}
		}
		return blobMsg{
			blob: blob,
		}
	}
}

func listenPublications(blobChan *lit.BlobChan, client lit.Library, query string) tea.Cmd {
	return func() tea.Msg {
		lit.GetLiterature(context.Background(), blobChan, client, lit.Request{
			Query: query,
		})
		return nil
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		handleBlob(m.next),
		listenPublications(m.next, m.client, m.query),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.progress.Width = msg.Width - Margin*2 - 4
		if m.progress.Width > MaxWidth {
			m.progress.Width = MaxWidth
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	case errMsg:
		m.done = true
		m.err = msg.err
	case blobMsg:
		m.received++
		pub, err := m.client.ParsePublication(msg.blob)
		if err != nil {
			m.err = err
			return m, nil
		}
		ref := m.client.ToBibTeX(pub)

		data, err := msg.blob.Marshal()
		if err != nil {
			m.err = err
			return m, nil
		}

		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: m.client.GetName(),
			Scope:  "lit",
			Action: "add_blob",
			Data:   []string{ref.CiteKey(), data},
		}); err != nil {
			m.err = err
			return m, nil
		}
		return m, handleBlob(m.next)
	}
	return m, nil
}

var (
	bodyStyle    = lipgloss.NewStyle().Width(MaxWidth).Margin(1)
	titleStyle   = lipgloss.NewStyle().Bold(true).Blink(true).MarginBottom(1)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EE6FF8")).Bold(true)
	helpStyle    = lipgloss.NewStyle()
)

func (m model) View() string {
	title := fmt.Sprintf("Downloading %q (%d results) from %s...", m.query, m.max, m.client.GetName())
	titleView := titleStyle.Render(title)
	progressView := lipgloss.NewStyle().MarginBottom(1).Render(m.progress.ViewAs(float64(m.received) / float64(m.max)))
	helpView := helpStyle.Render(m.help.View(keys))

	var statusView string
	if m.err != nil {
		statusView = errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	} else if m.done {
		statusView = successStyle.Render("Done!")
	}

	return bodyStyle.Render(fmt.Sprintf("%s\n%s\n%s\n%s\n",
		titleView,
		progressView,
		statusView,
		helpView,
	))
}

func Program(db *edb.Db, client lit.Library, opts ...tea.ProgramOption) (*tea.Program, error) {
	query := ""
	if err := db.Revive(func(e edb.Event) error {
		// TODO: we can easility recover from previous download sessions by checking:
		// 1. wether the max number of items did non change.
		// 2. if yes, start counting from the last item received.
		switch e.Action {
		case "set_query":
			query = e.Data[0]
			return nil
		default:
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if query == "" {
		return nil, fmt.Errorf("query not found within edb. Did you run lit-max?")
	}

	// We might read the max value from edb set_query as well. It might
	// have changed in the meanwhile though!
	max, err := client.GetMaxLiterature(context.Background(), lit.Request{
		Query: query,
	})
	if err != nil {
		return nil, err
	}

	return tea.NewProgram(model{
		db:       db,
		client:   client,
		query:    query,
		max:      max,
		next:     lit.NewBlobChan(max, 0),
		progress: progress.NewModel(progress.WithDefaultGradient()),
		help:     help.NewModel(),
	}, opts...), nil
}

func Main(db *edb.Db, client lit.Library, opts ...tea.ProgramOption) error {
	p, err := Program(db, client, opts...)
	if err != nil {
		return err
	}
	return p.Start()
}

func main() {
	flag.Parse()

	db, err := edb.Open(*edbPath)
	if err != nil {
		log.Fatale(err)
	}

	client := scopus.NewClient(scopusKey)
	err = Main(db, client, tea.WithoutCatchPanics())
	db.Close()

	if err != nil {
		log.Fatale(err)
	}
}
