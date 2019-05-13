package wikigopher_test

import (
	"strings"
	"testing"

	"github.com/ishaangandhi/wikigopher"
)

func TestPage(t *testing.T) {
	cases := []struct {
		query, title string
		pageid       uint64
	}{
		{"nyc", "New York City", 645042},
	}
	for _, c := range cases {
		got, err := wikigopher.Page(c.query)
		if err != nil {
			t.Errorf("Page errored with query %q and err %q", c.query, err)
		}
		if got.PageID != c.pageid {
			t.Errorf("Page failed with query %q and page id %d", c.query, got.PageID)
		}
	}
}

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
	for _, c := range cases {
		got := wikigopher.Search(c.query)
		if got[0] != c.search[0] {
			t.Errorf("Search(%q) == %q, wanted %q", c.query, got, c.search)
		}
	}
}
