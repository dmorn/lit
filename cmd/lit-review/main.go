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
	Accept    key.Binding
	Highlight key.Binding
	Reject    key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Accept, k.Highlight, k.Reject}, // first column
		{k.Help, k.Quit},                  // second column
	}
}

var keys = keyMap{
	Accept: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "accept publication"),
	),
	Highlight: key.NewBinding(
		key.WithKeys("A"),
		key.WithHelp("A", "accept publication, with highlight"),
	),
	Reject: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reject publication"),
	),
	Help: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	db     *edb.Db
	query  string
	client lit.Library

	cursor        int
	acceptedCount int
	rejectedCount int
	pubs          []lit.Publication
	err           error

	help     help.Model
	progress progress.Model
}

func (m model) Init() tea.Cmd {
	return getAbstract(m.client, m.pubs[m.cursor])
}

type cursorMsg int

func moveCursor(n int) tea.Cmd {
	return func() tea.Msg {
		return cursorMsg(n)
	}
}

func makeReview(p lit.Publication, r lit.Review, next int) tea.Cmd {
	return func() tea.Msg {
		p.Review = &r
		return publicationMsg{
			pub:  p,
			next: moveCursor(next),
		}
	}
}

type errMsg struct {
	err error
}

func (m errMsg) Error() string {
	return m.err.Error()
}

type publicationMsg struct {
	pub  lit.Publication
	next tea.Cmd
}

func getAbstract(client lit.Library, p lit.Publication) tea.Cmd {
	return func() tea.Msg {
		if p.Abstract != nil {
			return nil
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := p.GetAbstract(ctx, client); err != nil {
			return errMsg{err}
		}
		return publicationMsg{
			pub: p,
		}
	}
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
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, keys.Accept):
			return m, makeReview(m.pubs[m.cursor], lit.Review{
				IsAccepted: true,
			}, m.cursor+1)
		case key.Matches(msg, keys.Highlight):
			return m, makeReview(m.pubs[m.cursor], lit.Review{
				IsAccepted:    true,
				IsHighlighted: true,
			}, m.cursor+1)
		case key.Matches(msg, keys.Reject):
			return m, makeReview(m.pubs[m.cursor], lit.Review{
				IsAccepted:   false,
				RejectReason: "rejected with no particular reason",
			}, m.cursor+1)
		}
	case errMsg:
		// TODO: we actually don't want to quit here but rather show the
		// error to the user. It may want to retry.
		m.err = msg
		return m, nil
	case publicationMsg:
		data, err := msg.pub.Marshal()
		if err != nil {
			m.err = err
			return m, tea.Quit
		}
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "reviewer", // TODO: reviewer id?
			Scope:  "lit",
			Action: "update_lit",
			Data:   []string{data, fmt.Sprintf("%d", m.cursor)},
		}); err != nil {
			m.err = err
			return m, tea.Quit
		}
		m.pubs[m.cursor] = msg.pub
		return m, msg.next
	case cursorMsg:
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "lit",
			Scope:  "lit",
			Action: "move_cursor",
			Data:   []string{fmt.Sprintf("%d", int(msg))},
		}); err != nil {
			m.err = err
			return m, tea.Quit
		}
		m.cursor = int(msg)
		return m, getAbstract(m.client, m.pubs[m.cursor])
	}
	return m, nil
}

var (
	bodyStyle     = lipgloss.NewStyle().Width(MaxWidth).Margin(Margin)
	titleStyle    = bodyStyle.Copy().Bold(true).MarginBottom(Margin)
	abstractStyle = bodyStyle.Copy()
	loadingStyle  = bodyStyle.Copy().Blink(true)
	errorStyle    = bodyStyle.Copy().Foreground(lipgloss.Color("5"))
	helpStyle     = bodyStyle.Copy()
	statsStyle    = bodyStyle.Copy()
)

func (m model) View() string {
	p := m.pubs[m.cursor]

	title := titleStyle.Render(p.Title)
	abstract := loadingStyle.Render("downloading abstract...")
	switch {
	case m.err != nil:
		abstract = errorStyle.Render(fmt.Sprintf("error: %v", m.err))
	case p.Abstract != nil:
		abstract = abstractStyle.Render(p.Abstract.GetText())
	}
	progress := bodyStyle.Render(m.progress.ViewAs(float64(m.cursor+1) / float64(len(m.pubs))))

	stats := statsStyle.Render(fmt.Sprintf("[accepted=%d rejected=%d total=%d current=%d]",
		m.acceptedCount,
		m.rejectedCount,
		len(m.pubs),
		m.cursor+1,
	))

	body := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		title,
		abstract,
		progress,
		stats,
	)

	helpView := helpStyle.Render(m.help.View(keys))
	return body + helpView
}

func Main() error {
	db, err := edb.Open(*edbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := ""
	pubs := []lit.Publication{}
	cursor := 0
	rejectedCount := 0
	acceptedCount := 0
	if err := db.Revive(func(e edb.Event) error {
		switch e.Action {
		case "set_query":
			query = e.Data[0]
		case "add_lit":
			p := new(lit.Publication)
			if err := p.Unmarshal(e.Data[0]); err != nil {
				return err
			}
			pubs = append(pubs, *p)
		case "update_lit":
			p := new(lit.Publication)
			if err := p.Unmarshal(e.Data[0]); err != nil {
				return err
			}
			if rev := p.Review; rev != nil {
				if rev.IsAccepted {
					acceptedCount++
				} else {
					rejectedCount++
				}
			}
			index, err := strconv.Atoi(e.Data[1])
			if err != nil {
				return err
			}
			pubs[index] = *p // Yes I'm buying myself a panic
		case "move_cursor":
			var err error
			cursor, err = strconv.Atoi(e.Data[0])
			if err != nil {
				return err
			}
		default:
			fmt.Printf("unhandled event %s\n", e.Id)
		}
		return nil
	}); err != nil {
		return err
	}
	if seen := rejectedCount + acceptedCount; seen != cursor+1 {
		return fmt.Errorf("database inconsistency: saw %d reviews (accepted %d + rejected %d) with cursor @%d")
	}

	return tea.NewProgram(model{
		db:            db,
		client:        scopus.NewClient(scopusKey),
		cursor:        cursor,
		acceptedCount: acceptedCount,
		rejectedCount: rejectedCount,
		query:         query,
		pubs:          pubs,
		help:          help.NewModel(),
		progress:      progress.NewModel(progress.WithDefaultGradient()),
	}, tea.WithAltScreen()).Start()
}

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
