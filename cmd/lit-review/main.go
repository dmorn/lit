package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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

const CommandsHelp = "Accept=[a] Accept+Highligh=[A] Reject=[r] Quit=[q]"

type model struct {
	db     *edb.Db
	query  string
	client lit.Library

	cursor int
	pubs   []lit.Publication
	err    error
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return m, tea.Quit
		case "a":
			return m, makeReview(m.pubs[m.cursor], lit.Review{
				IsAccepted: true,
			}, m.cursor+1)
		case "A":
			return m, makeReview(m.pubs[m.cursor], lit.Review{
				IsAccepted:    true,
				IsHighlighted: true,
			}, m.cursor+1)
		case "r":
			return m, makeReview(m.pubs[m.cursor], lit.Review{
				IsAccepted:   false,
				RejectReason: "rejected with no particular reason",
			}, m.cursor+1)
		}
	case errMsg:
		// TODO: we actually don't want to quit here but rather show the
		// error to the user. It may want to retry.
		m.err = msg
		return m, tea.Quit
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

func (m model) View() string {
	p := m.pubs[m.cursor]
	var b strings.Builder

	b.WriteString(fmt.Sprintf("%s\n", p.Title))
	switch {
	case m.err != nil:
		b.WriteString(fmt.Sprintf("error: %v\n", m.err))
	case p.Abstract != nil:
		b.WriteString(fmt.Sprintf("\n"))

		r := strings.NewReader(p.Abstract.Text)
		buf := make([]byte, 80)
		for {
			n, _ := r.Read(buf)
			if n < len(buf) {
				break
			}
			b.WriteString(fmt.Sprintf("%s\n", string(buf)))
		}
		b.WriteString(fmt.Sprintf("%s\n\n", string(buf)))

	default:
		b.WriteString(fmt.Sprintf("downloading abstract...\n"))
	}
	b.WriteString(fmt.Sprintf("(%d/%d)\n", m.cursor+1, len(m.pubs)))
	b.WriteString(fmt.Sprintf("%s\n", CommandsHelp))
	return b.String()
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

func Main() error {
	db, err := edb.Open(*edbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := ""
	pubs := []lit.Publication{}
	cursor := 0
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

	return tea.NewProgram(model{
		db:     db,
		client: scopus.NewClient(scopusKey),
		cursor: cursor,
		query:  query,
		pubs:   pubs,
	}, tea.WithAltScreen()).Start()
}

func main() {
	flag.Parse()

	if err := Main(); err != nil {
		log.Fatale(err)
	}
}
