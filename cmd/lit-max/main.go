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
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jecoz/edb"
	"github.com/jecoz/lit"
	"github.com/jecoz/lit/log"
	"github.com/jecoz/lit/scopus"
)

const (
	MaxWidth = 120
	Margin   = 1
)

var (
	scopusKey = os.Getenv("SCOPUS_API_KEY")
)

var (
	edbPath = flag.String("edb", "lit.edb", "Event database file. Everything will be stored here.")
)

type keyMap struct {
	Search key.Binding
	Quit   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Search, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
		{},
	}
}

var keys = keyMap{
	Search: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "issue query"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc", "quit"),
	),
}

type errMsg struct {
	err error
}

func (m errMsg) Error() string {
	return m.err.Error()
}

type maxMsg struct {
	max   int
	query string
}

func getMaxLiterature(client lit.Library, q string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		max, err := client.GetMaxLiterature(ctx, lit.Request{
			Query: q,
		})
		if err != nil {
			return errMsg{err}
		}
		return maxMsg{
			max:   max,
			query: q,
		}
	}
}

type model struct {
	db     *edb.Db
	client lit.Library

	searching bool
	query     string
	max       int
	err       error

	help      help.Model
	textInput textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Search):
			m.searching = true
			m.err = nil
			return m, getMaxLiterature(m.client, m.textInput.Value())
		}
	case errMsg:
		m.searching = false
		m.err = msg
		return m, nil
	case maxMsg:
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "reviewer", // TODO: reviewer id?
			Scope:  "lit",
			Action: "set_query",
			Data:   []string{msg.query, fmt.Sprintf("%d", msg.max)},
		}); err != nil {
			m.err = err
			return m, nil
		}
		m.searching = false
		m.max = msg.max
		m.query = msg.query
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

var (
	bodyStyle      = lipgloss.NewStyle().Width(MaxWidth).Margin(Margin)
	resultStyle    = lipgloss.NewStyle()
	errorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	helpStyle      = lipgloss.NewStyle()
	inputStyle     = lipgloss.NewStyle().MarginBottom(1)
	searchingStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	})
)

func (m model) queryView() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("error: %v", m.err))
	}
	if m.searching {
		return searchingStyle.Render("searching...")
	}
	return resultStyle.Render(fmt.Sprintf("%q hit %d results", m.query, m.max))
}

func (m model) View() string {
	textView := inputStyle.Render(m.textInput.View())
	helpView := helpStyle.Render(m.help.View(keys))

	return bodyStyle.Render(fmt.Sprintf("%s\n%s\n%s\n",
		textView,
		lipgloss.NewStyle().MarginBottom(1).Render(m.queryView()),
		helpView,
	))
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
		switch e.Action {
		case "set_query":
			var err error
			query = e.Data[0]
			max, err = strconv.Atoi(e.Data[1])
			if err != nil {
				return fmt.Errorf("parse maximum literature count: %w", err)
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return err
	}

	ti := textinput.NewModel()
	ti.Placeholder = "(FPGA AND GPU) AND NN"
	ti.Focus()
	ti.CharLimit = 256

	return tea.NewProgram(model{
		db:        db,
		client:    scopus.NewClient(scopusKey),
		textInput: ti,
		query:     query,
		max:       max,
		help:      help.NewModel(),
	}).Start()
}

func main() {
	flag.Parse()
	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
