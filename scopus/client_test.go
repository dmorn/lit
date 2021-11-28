package scopus

import (
	"os"
	"strings"
	"testing"
)

func TestExtractsAbstract(t *testing.T) {
	wantBytes, err := os.ReadFile("abstract.response.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := strings.TrimSpace(string(wantBytes))

	html, err := os.Open("abstract.response.html")
	if err != nil {
		t.Fatal(err)
	}
	defer html.Close()

	have, err := ParseAbstract(html)
	if err != nil {
		t.Fatal(err)
	}
	have = strings.TrimSpace(have)

	if have != want {
		t.Fatalf("want:\n%q\nhave:\n%q", want, have)
	}
}
