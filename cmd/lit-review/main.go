package main

import (
	"archive/zip"
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jecoz/edb"
	"github.com/jecoz/lit"
	"github.com/jecoz/lit/bibtex"
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
	PlaceholderReject = "Rejected due to..."
	PlaceholderPrint  = "archive.zip name/path"
)

const MaxWidth = 120

type insertMode struct {
	Enter  key.Binding
	Cancel key.Binding
}

func (k insertMode) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Cancel}
}

func (k insertMode) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.ShortHelp(),
		{},
	}
}

func newInsertMode() insertMode {
	return insertMode{
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm input"),
		),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "exit insert mode"),
		),
	}
}

type normalMode struct {
	Left  key.Binding
	Right key.Binding

	Accept    key.Binding
	Highlight key.Binding
	Reject    key.Binding
	Print     key.Binding
	Inspect   key.Binding

	Help key.Binding
	Quit key.Binding
}

func (k normalMode) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k normalMode) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Right, k.Accept, k.Highlight, k.Reject, k.Print, k.Inspect},
		{k.Help, k.Quit},
	}
}

func newNormalMode() normalMode {
	return normalMode{
		Help: key.NewBinding(
			key.WithKeys("H", "?"),
			key.WithHelp("?/H", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
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
			key.WithHelp("r", "reject publication (providing a reason)"),
		),
		Print: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "print review"),
		),
		Inspect: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "inspect publication data blob"),
		),
	}
}

type model struct {
	db     *edb.Db
	query  string
	client lit.Library

	normal normalMode
	insert insertMode
	style  style

	cursor int
	pubs   []lit.Publication

	rejecting     bool
	printing      bool
	inspecting    bool
	inspection    string
	rejectedCount int
	acceptedCount int
	err           error

	help      help.Model
	progress  progress.Model
	textInput textinput.Model
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

type reviewMsg struct {
	cursor int
	pub    lit.Publication
}

func makeReview(cursor int, p lit.Publication, r lit.Review) tea.Cmd {
	return func() tea.Msg {
		if p.Review != nil {
			return nil
		}

		p.Review = &r
		return reviewMsg{
			cursor: cursor,
			pub:    p,
		}
	}
}

type inspectMsg struct {
	dump string
}

func makeInspection(client lit.Library, db *edb.Db, pub lit.Publication) tea.Cmd {
	return func() tea.Msg {
		var (
			event edb.Event
			ok    bool
		)
		key := client.ToBibTeX(pub).CiteKey()
		db.Revive(func(e edb.Event) error {
			if e.Scope == "lit" && e.Action == "add_blob" && e.Data[0] == key {
				ok = true
				event = e
				return fmt.Errorf("stop")
			}
			return nil
		})
		if !ok {
			return errMsg{
				fmt.Errorf("make inspection: blob %q not found", key),
			}
		}
		blob := new(lit.Blob)
		if err := blob.Unmarshal(event.Data[1]); err != nil {
			return errMsg{err}
		}
		var buf bytes.Buffer
		if err := client.PrettyPrint(*blob, &buf); err != nil {
			return errMsg{err}
		}
		return inspectMsg{
			dump: buf.String(),
		}
	}
}

func saveReview(client lit.Library, name string, pubs []lit.Publication) tea.Cmd {
	return func() tea.Msg {
		accepted := make([]bibtex.Reference, 0, len(pubs))
		rejected := make([]bibtex.Reference, 0, len(pubs))
		for _, v := range pubs {
			if v.Review != nil && v.Review.IsAccepted {
				accepted = append(accepted, client.ToBibTeX(v))
			}
			if v.Review != nil && !v.Review.IsAccepted {
				rejected = append(rejected, client.ToBibTeX(v))
			}
		}

		f, err := os.Create(name)
		if err != nil {
			return errMsg{err}
		}
		defer f.Close()

		archive := zip.NewWriter(f)
		buf, err := archive.Create("accepted.bib")
		if err != nil {
			return errMsg{err}
		}
		if err := bibtex.MarshalBibTeXReferenceList(buf, accepted); err != nil {
			return errMsg{err}
		}
		buf, err = archive.Create("rejected.bib")
		if err != nil {
			return errMsg{err}
		}
		if err := bibtex.MarshalBibTeXReferenceList(buf, rejected); err != nil {
			return errMsg{err}
		}

		if err := archive.Close(); err != nil {
			return errMsg{err}
		}
		return nil
	}
}

type errMsg struct {
	err error
}

func (m errMsg) Error() string {
	return m.err.Error()
}

type abstractMsg struct {
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
		return abstractMsg{
			pub:    p,
			cursor: cursor,
		}
	}
}

type quitMsg struct{}

func (m model) handleKeyNormal(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	keys := m.normal
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
		m.textInput.Focus()
		m.textInput.Placeholder = PlaceholderReject
		m.rejecting = true
		return m, nil
	case key.Matches(msg, keys.Print):
		m.textInput.Focus()
		m.textInput.Placeholder = PlaceholderPrint
		m.textInput.SetValue(fmt.Sprintf("review-%s.zip", time.Now().Format(time.RFC3339)))
		m.printing = true
		return m, nil
	case key.Matches(msg, keys.Inspect):
		return m, makeInspection(m.client, m.db, m.pubs[m.cursor])
	}
	return m, nil
}

func (m model) handleKeyInsert(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	reset := func(m *model) {
		m.textInput.Reset()
		m.textInput.Blur()
		m.rejecting = false
		m.printing = false
		m.inspecting = false
	}

	keys := m.insert
	switch {
	case key.Matches(msg, keys.Cancel):
		reset(&m)
		return m, nil
	case key.Matches(msg, keys.Enter):
		if len(m.textInput.Value()) == 0 || m.inspecting {
			reset(&m)
			return m, nil
		}

		var cmd tea.Cmd
		switch {
		case m.rejecting:
			cmd = tea.Sequentially(
				makeReview(m.cursor, m.pubs[m.cursor], lit.Review{
					IsAccepted:   false,
					RejectReason: m.textInput.Value(),
				}),
				moveCursor(m.cursor+1),
			)
		case m.printing:
			cmd = saveReview(m.client, m.textInput.Value(), m.pubs)
		}

		reset(&m)
		return m, cmd
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.rejecting || m.printing || m.inspecting {
		return m.handleKeyInsert(msg)
	}
	return m.handleKeyNormal(msg)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.progress.Width = msg.Width - 1*2 - 4
		if m.progress.Width > MaxWidth {
			m.progress.Width = MaxWidth
		}
	case tea.KeyMsg:
		return m.handleKey(msg)
	case errMsg:
		m.err = msg
		return m, nil
	case abstractMsg:
		data, err := msg.pub.Abstract.Marshal()
		if err != nil {
			m.err = err
			return m, nil
		}
		ref := m.client.ToBibTeX(msg.pub)
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "reviewer",
			Scope:  "lit",
			Action: "add_abstract",
			Data:   []string{ref.CiteKey(), data, fmt.Sprintf("%d", msg.cursor)},
		}); err != nil {
			m.err = err
			return m, nil
		}
		m.pubs[msg.cursor] = msg.pub
		return m, nil
	case reviewMsg:
		data, err := msg.pub.Review.Marshal()
		if err != nil {
			m.err = err
			return m, nil
		}
		ref := m.client.ToBibTeX(msg.pub)
		if err := m.db.Append(&edb.Event{
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Issuer: "reviewer",
			Scope:  "lit",
			Action: "add_review",
			Data:   []string{ref.CiteKey(), data, fmt.Sprintf("%d", msg.cursor)},
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
	case inspectMsg:
		m.inspecting = true
		m.inspection = msg.dump
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
	bold lipgloss.Style

	abstract lipgloss.Style
	link     lipgloss.Style
	err      lipgloss.Style

	rejected lipgloss.Style
	accepted lipgloss.Style
	todo     lipgloss.Style
}

var defaultStyle = style{
	bold: lipgloss.NewStyle().Bold(true),

	abstract: lipgloss.NewStyle(),
	link:     lipgloss.NewStyle(),
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
	ref := m.client.ToBibTeX(p)
	return m.style.abstract.Render(fmt.Sprintf("%s, %d (%s, %s)", p.Creator, p.CoverDate.Year(), ref.CiteKey(), ref.EntryType()))
}

func (m model) linkView() string {
	p := m.pubs[m.cursor]
	l := m.client.ReferenceLink(p)
	return m.style.link.Render(l)
}

func (m model) statusView() string {
	if m.rejecting {
		return m.textInput.View()
	}

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
	if m.rejecting || m.inspecting {
		return m.help.View(m.insert)
	}

	return m.help.View(m.normal)
}

func (m model) View() string {
	container := lipgloss.NewStyle().Margin(1)

	progress := lipgloss.NewStyle().MarginBottom(1).Render(
		m.progressView(),
	)
	footer := lipgloss.NewStyle().MarginTop(1).Render(fmt.Sprintf("%s\n%s\n%s",
		progress,
		m.statsView(),
		m.helpView(),
	))

	if m.printing {
		return container.Render(fmt.Sprintf("%s\n%s\n",
			m.textInput.View(),
			footer,
		))
	}
	if m.inspecting {
		return container.Render(fmt.Sprintf("%s\n%s",
			m.inspection,
			footer,
		))
	}

	header := fmt.Sprintf("%s\n%s\n%s",
		m.titleView(),
		m.creatorView(),
		m.statusView(),
	)

	abstract := lipgloss.NewStyle().MarginTop(1).MarginBottom(1).Width(MaxWidth).Render(m.abstractView())
	body := fmt.Sprintf("%s\n%s", abstract, m.linkView())

	return container.Render(fmt.Sprintf("%s\n%s\n%s",
		header,
		body,
		footer,
	))
}

func Main() error {
	db, err := edb.Open(*edbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	client := scopus.NewClient(scopusKey)
	query := ""
	pubs := []lit.Publication{}
	cursor := 0
	rejectedCount := 0
	acceptedCount := 0
	if err := db.Revive(func(e edb.Event) error {
		switch e.Action {
		case "set_query":
			query = e.Data[0]
		case "add_blob":
			b := new(lit.Blob)
			if err := b.Unmarshal(e.Data[1]); err != nil {
				return fmt.Errorf("add_blob: unmarshal: %w", err)
			}
			pub, err := client.ParsePublication(*b)
			if err != nil {
				return fmt.Errorf("add_blob: parse publication: %w", err)
			}
			pubs = append(pubs, pub)
		case "add_abstract":
			index, err := strconv.Atoi(e.Data[2])
			if err != nil {
				return fmt.Errorf("add_abstract: index conversion: %w", err)
			}
			a := new(lit.Abstract)
			if err := a.Unmarshal(e.Data[1]); err != nil {
				return err
			}
			pubs[index].Abstract = a
		case "add_review":
			index, err := strconv.Atoi(e.Data[2])
			if err != nil {
				return fmt.Errorf("add_review: index conversion: %w", err)
			}
			r := new(lit.Review)
			if err := r.Unmarshal(e.Data[1]); err != nil {
				return err
			}
			if r.IsAccepted {
				acceptedCount++
			} else {
				rejectedCount++
			}
			pubs[index].Review = r
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

	ti := textinput.NewModel()
	ti.CharLimit = 256

	return tea.NewProgram(model{
		db:            db,
		client:        client,
		style:         defaultStyle,
		insert:        newInsertMode(),
		normal:        newNormalMode(),
		cursor:        cursor,
		acceptedCount: acceptedCount,
		rejectedCount: rejectedCount,
		query:         query,
		pubs:          pubs,
		help:          help.NewModel(),
		progress:      progress.NewModel(progress.WithDefaultGradient()),
		textInput:     ti,
	}).Start()
}

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
