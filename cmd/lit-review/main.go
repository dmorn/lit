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
	Left      key.Binding
	Right     key.Binding
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
		{k.Left, k.Right, k.Accept, k.Highlight, k.Reject},
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
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
		key.WithKeys("H", "?"),
		key.WithHelp("?/H", "toggle help"),
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

	style style

	cursor int
	pubs   []lit.Publication

	rejectedCount int
	acceptedCount int
	err           error

	help     help.Model
	progress progress.Model
}

func (m model) Init() tea.Cmd {
	return getAbstract(m.client, m.cursor, m.pubs[m.cursor])
}

type cursorMsg int

func moveCursor(n int) tea.Cmd {
	return func() tea.Msg {
		return cursorMsg(n)
	}
}

func makeReview(cursor int, p lit.Publication, r lit.Review) tea.Cmd {
	return func() tea.Msg {
		if p.Review != nil {
			return nil
		}

		p.Review = &r
		return publicationMsg{
			cursor: cursor,
			pub:    p,
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
	cursor int
	pub    lit.Publication
}

func getAbstract(client lit.Library, cursor int, p lit.Publication) tea.Cmd {
	return func() tea.Msg {
		if p.Abstract != nil {
			return nil
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := p.GetAbstract(ctx, client); err != nil {
			return errMsg{err}
		}
		return publicationMsg{
			pub:    p,
			cursor: cursor,
		}
	}
}

type quitMsg struct{}

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
			return m, func() tea.Msg {
				return quitMsg{}
			}
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, keys.Left):
			return m, moveCursor(m.cursor - 1)
		case key.Matches(msg, keys.Right):
			return m, moveCursor(m.cursor + 1)
		case key.Matches(msg, keys.Accept):
			return m, tea.Sequentially(
				makeReview(m.cursor, m.pubs[m.cursor], lit.Review{
					IsAccepted: true,
				}),
				moveCursor(m.cursor+1),
			)
		case key.Matches(msg, keys.Highlight):
			return m, tea.Sequentially(
				makeReview(m.cursor, m.pubs[m.cursor], lit.Review{
					IsAccepted:    true,
					IsHighlighted: true,
				}),
				moveCursor(m.cursor+1),
			)
		case key.Matches(msg, keys.Reject):
			// TODO: promt the user for a reason, then make the review.
			return m, tea.Sequentially(
				makeReview(m.cursor, m.pubs[m.cursor], lit.Review{
					IsAccepted:   false,
					RejectReason: "rejected with no particular reason",
				}),
				moveCursor(m.cursor+1),
			)
		}
	case errMsg:
		m.err = msg
		return m, nil
	case publicationMsg:
		data, err := msg.pub.Marshal()
		if err != nil {
			m.err = err
			return m, nil
		}
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "reviewer",
			Scope:  "lit",
			Action: "update_lit",
			Data:   []string{msg.pub.BibId(), data, fmt.Sprintf("%d", msg.cursor)},
		}); err != nil {
			m.err = err
			return m, nil
		}
		m.pubs[msg.cursor] = msg.pub
		if rev := msg.pub.Review; rev != nil {
			if rev.IsAccepted {
				m.acceptedCount++
			} else {
				m.rejectedCount++
			}
		}
		return m, nil
	case cursorMsg:
		cursor := int(msg)
		switch {
		case cursor < 0:
			cursor = len(m.pubs) + cursor
		case cursor > len(m.pubs)-1:
			cursor = cursor - len(m.pubs)
		}
		m.cursor = cursor
		m.err = nil
		return m, getAbstract(m.client, m.cursor, m.pubs[m.cursor])
	case quitMsg:
		// NOTE: if an error occurs here we won't catch it.
		m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "reviewer",
			Scope:  "lit",
			Action: "move_cursor",
			Data:   []string{fmt.Sprintf("%d", int(m.cursor))},
		})
		return m, tea.Quit
	}
	return m, nil
}

type style struct {
	body lipgloss.Style
	bold lipgloss.Style

	abstract lipgloss.Style
	err      lipgloss.Style

	rejected lipgloss.Style
	accepted lipgloss.Style
	todo     lipgloss.Style
}

var defaultStyle = style{
	body: lipgloss.NewStyle().Width(MaxWidth).Margin(1),
	bold: lipgloss.NewStyle().Bold(true),

	abstract: lipgloss.NewStyle(),
	err:      lipgloss.NewStyle().Foreground(lipgloss.Color("5")), // TODO: change this color

	rejected: lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Bold(true),
	accepted: lipgloss.NewStyle().Foreground(lipgloss.Color("#EE6FF8")).Bold(true),
	todo: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	}).Bold(true),
}

func (m model) titleView() string {
	return m.style.bold.Height(1).Render(fmt.Sprintf("#%d: %s", m.cursor+1, m.pubs[m.cursor].Title))
}

func (m model) creatorView() string {
	p := m.pubs[m.cursor]
	return m.style.abstract.Render(fmt.Sprintf("%s, %d (%s)", p.Creator, p.CoverDate.Year, p.BibId()))
}

func (m model) statusView() string {
	rev := m.pubs[m.cursor].Review

	var statusView string
	switch {
	case rev == nil:
		statusView = m.style.todo.Render("to be reviewed")
	case rev.IsAccepted && rev.IsHighlighted:
		statusView = m.style.accepted.Render("accepted (+highlight)")
	case rev.IsAccepted:
		statusView = m.style.accepted.Render("accepted")
	default:
		statusView = m.style.rejected.Render("rejected: " + rev.RejectReason)
	}
	return statusView
}

func (m model) progressView() string {
	return m.progress.ViewAs(float64(m.acceptedCount+m.rejectedCount) / float64(len(m.pubs)))
}

func (m model) abstractView() string {
	p := m.pubs[m.cursor]
	abstractView := m.style.todo.Render("downloading abstract...")
	switch {
	case m.err != nil:
		abstractView = m.style.err.Render(fmt.Sprintf("error: %v", m.err))
	case p.Abstract != nil:
		abstractView = m.style.abstract.Render(p.Abstract.GetText())
	}
	return abstractView
}

func (m model) statsView() string {
	return m.style.abstract.Render(fmt.Sprintf("[total=%d todo=%d accepted=%d rejected=%d]",
		len(m.pubs),
		len(m.pubs)-(m.acceptedCount+m.rejectedCount),
		m.acceptedCount,
		m.rejectedCount,
	))
}

func (m model) helpView() string {
	return m.help.View(keys)
}

func (m model) View() string {
	header := lipgloss.NewStyle().Render(fmt.Sprintf("%s\n%s\n%s",
		m.titleView(),
		m.creatorView(),
		m.statusView(),
	))

	progress := lipgloss.NewStyle().MarginBottom(1).MarginTop(1).Render(
		m.progressView(),
	)

	footer := lipgloss.NewStyle().MarginTop(1).Render(fmt.Sprintf("%s\n%s",
		m.statsView(),
		m.helpView(),
	))

	return m.style.body.Render(fmt.Sprintf("%s\n%s\n%s\n%s",
		header,
		progress,
		m.abstractView(),
		footer,
	))
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
			if err := p.Unmarshal(e.Data[1]); err != nil {
				return fmt.Errorf("add_lit: %w", err)
			}
			pubs = append(pubs, *p)
		case "update_lit":
			p := new(lit.Publication)
			if err := p.Unmarshal(e.Data[1]); err != nil {
				return err
			}
			if rev := p.Review; rev != nil {
				if rev.IsAccepted {
					acceptedCount++
				} else {
					rejectedCount++
				}
			}
			index, err := strconv.Atoi(e.Data[2])
			if err != nil {
				return fmt.Errorf("update_lit: index conversion: %w", err)
			}
			pubs[index] = *p // Yes I'm buying myself a panic
		case "move_cursor":
			var err error
			cursor, err = strconv.Atoi(e.Data[0])
			if err != nil {
				return fmt.Errorf("move_cursor: cursor conversion: %w", err)
			}
		default:
		}
		return nil
	}); err != nil {
		return err
	}

	return tea.NewProgram(model{
		db:            db,
		client:        scopus.NewClient(scopusKey),
		style:         defaultStyle,
		cursor:        cursor,
		acceptedCount: acceptedCount,
		rejectedCount: rejectedCount,
		query:         query,
		pubs:          pubs,
		help:          help.NewModel(),
		progress:      progress.NewModel(progress.WithDefaultGradient()),
	}).Start()
}

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
