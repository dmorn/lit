// see https://www.economics.utoronto.ca/osborne/latex/BIBTEX.HTM
package bibtex

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type EntryType string

const (
	// any article published in a periodical like a journal article or
	// magazine article
	EntryTypeArticle EntryType = "article"
	EntryTypeBook              = "book"
	// A book w/o a designated publisher
	EntryTypeBooklet = "booklet"
	// A conference paper
	EntryTypeConference = "conference"
	// A section or chapter of a book
	EntryTypeInBook = "inbook"
	// An article in a collection
	EntryTypeInCollection = "incollection"
	// A conference paper (same as EntryTypeConference)
	EntryTypeInProceedings = "inproceedings"
	// A technical manual
	EntryTypeManual       = "manual"
	EntryTypeMasterThesis = "masterthesis"
	// Used if nothing else fits
	EntryTypeMisc      = "misc"
	EntryTypePhDThesis = "phdthesis"
	// A technical report, government report or white paper
	EntryTypeTechReport = "techreport"
	// A work that has not yet been officially published
	EntryTypeUnpublished = "unpublished"
)

type Reference interface {
	EntryType() EntryType
	CommonInfo() *Entry
	CiteKey() string
	Fields() map[string]string
}

func MarshalBibTeXReference(w io.Writer, ref Reference) error {
	maxKeyLen := 0
	for k := range ref.Fields() {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}
	keyFmt := fmt.Sprintf("%%%ds", maxKeyLen)

	fmt.Fprintf(w, "@%s{%s,\n", string(ref.EntryType()), ref.CiteKey())
	for k, v := range ref.Fields() {
		fmt.Fprintf(w, "\t"+keyFmt+" = {%s}\n", k, strings.ReplaceAll(v, "\n", ""))
	}
	fmt.Fprintf(w, "}\n")
	return nil
}

func MarshalBibTeXReferenceList(w io.Writer, refs []Reference) error {
	for i, v := range refs {
		if err := MarshalBibTeXReference(w, v); err != nil {
			return fmt.Errorf("marshal BibTeX, item %d: %w", i, err)
		}
	}
	return nil
}

type Entry struct {
	Title  string
	Author string
	Year   int

	// optional fields
	DOI      *string // e.g. 10.1038/d41586-018-07848-2
	Issn     *string // e.g. 1476-4687
	Isbn     *string // e.g. 9780201896831
	Url      *string
	Abstract *string
}

func (e Entry) AuthorShort() string {
	return takeMaxRunes(e.Author, 3)
}

func (e Entry) TitleShort() string {
	return takeMaxRunes(e.Title, 5)
}

func (e Entry) CiteKey() string {
	return fmt.Sprintf("%s%d%s", e.AuthorShort(), e.Year, e.TitleShort())
}

func (e Entry) CommonInfo() *Entry {
	return &e
}

func (e Entry) Fields() map[string]string {
	return map[string]string{
		"title":  e.Title,
		"author": e.Author,
		"year":   fmt.Sprintf("%d", e.Year),
	}
}

type Article struct {
	Entry
	Journal string

	// optional fields
	Volume    *string
	PageRange *string
}

func (a Article) Fields() map[string]string {
	m := a.Entry.Fields()
	m["journal"] = a.Journal

	if vol := a.Volume; vol != nil {
		m["volume"] = *vol
	}
	if r := a.PageRange; r != nil {
		m["pages"] = *r
	}
	return m
}

func (a Article) EntryType() EntryType {
	return EntryTypeArticle
}

type Book struct {
	Entry
	Publisher string

	// optional fields
	Address *string
}

func (a Book) Fields() map[string]string {
	m := a.Entry.Fields()
	m["publisher"] = a.Publisher

	if addr := a.Address; addr != nil {
		m["address"] = *addr
	}
	return m
}

func (a Book) EntryType() EntryType {
	return EntryTypeBook
}

type InCollection struct {
	Entry
	BookTitle string
	Publisher string

	// optional fields
	Editor    *string
	PageRange *string
	Address   *string
}

func (a InCollection) Fields() map[string]string {
	m := a.Entry.Fields()
	m["booktitle"] = a.BookTitle
	m["publisher"] = a.Publisher

	if ed := a.Editor; ed != nil {
		m["editor"] = *ed
	}
	if r := a.PageRange; r != nil {
		m["pages"] = *r
	}
	if addr := a.Address; addr != nil {
		m["address"] = *addr
	}
	return m
}

func (a InCollection) EntryType() EntryType {
	return EntryTypeInCollection
}

type InProceedings struct {
	Entry
	BookTitle string

	// optional fields
	PageRange *string
	Publisher *string
	Address   *string
}

func (a InProceedings) Fields() map[string]string {
	m := a.Entry.Fields()
	m["booktitle"] = a.BookTitle

	if r := a.PageRange; r != nil {
		m["pages"] = *r
	}
	if p := a.Publisher; p != nil {
		m["publisher"] = *p
	}
	if addr := a.Address; addr != nil {
		m["address"] = *addr
	}
	return m
}

func (a InProceedings) EntryType() EntryType {
	return EntryTypeInProceedings
}

type TechReport struct {
	Entry
	Institution string

	// optional fields
	Number *string
	Type   *string
}

func (a TechReport) Fields() map[string]string {
	m := a.Entry.Fields()
	m["institution"] = a.Institution

	if num := a.Number; num != nil {
		m["number"] = *num
	}
	if t := a.Type; t != nil {
		m["type"] = *t
	}
	return m
}

func (a TechReport) EntryType() EntryType {
	return EntryTypeTechReport
}

type Misc struct {
	Entry

	// optional fields
	Note *string
}

func (a Misc) Fields() map[string]string {
	m := a.Entry.Fields()

	if note := a.Note; note != nil {
		m["note"] = *note
	}
	return m
}

func (a Misc) EntryType() EntryType {
	return EntryTypeMisc
}

func takeMaxRunes(s string, n int) string {
	s = strings.TrimSpace(s)
	r := strings.NewReader(s)
	take := 0
	for i := 0; i < n; i++ {
		ru, size, err := r.ReadRune()
		if err != nil || !unicode.IsLetter(ru) {
			break
		}
		take += size
	}
	return string([]byte(s)[:take])
}
