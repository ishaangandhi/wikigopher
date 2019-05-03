package wikigopher_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ishaangandhi/wikigopher"
)

func TestSummary(t *testing.T) {
	cases := []struct {
		query, summary string
	}{
		{"ishaan", "Ishaan Khatter (born 1 November"},
		{"wikipedia", "Wikipedia ( (listen),  "},
	}
	for _, c := range cases {
		got := wikigopher.Summary(c.query)
		if !strings.HasPrefix(got, c.summary) {
			t.Errorf("Summary(%q) == %q, wanted %q", c.query, got, c.summary)
		}
	}
}

func TestSearch(t *testing.T) {
	cases := []struct {
		query  string
		search []string
	}{
		{"wikipedia", []string{"Wikipedia"}},
	}
	var summary string
	summary = wikigopher.Summary("golang")
	fmt.Println(summary)
	for _, c := range cases {
		got := wikigopher.Search(c.query)
		if got[0] != c.search[0] {
			t.Errorf("Search(%q) == %q, wanted %q", c.query, got, c.search)
		}
	}
}
