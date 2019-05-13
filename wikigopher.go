package wikigopher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

const (
	apiURL           = "http://en.wikipedia.org/w/api.php"
	searchResultsCap = 10
)

// WikipediaPage represents a Wikipedia Page
type WikipediaPage struct {
	PageID     uint64
	Title      string
	Content    string
	Categories []string
	Links      []string
}

// Page returns the WikipediaPage object for a given search query
func Page(query string) (*WikipediaPage, error) {
	// first run Search() to get most relevant result for query
	var searchResults = Search(query)
	v := url.Values{
		"titles":      []string{searchResults[0]},
		"action":      []string{"query"},
		"prop":        []string{"categories|extracts|revisions|links"},
		"rvprop":      []string{"ids"},
		"format":      []string{"json"},
		"explaintext": []string{""},
	}

	var page WikipediaPage
	respJSON, err := wikiRequest(v)
	if err != nil {
		return &page, err
	}

	pageJSON := gjson.GetBytes(respJSON, "query.pages.*")
	page.PageID = pageJSON.Get("pageid").Uint()
	page.Title = pageJSON.Get("title").String()
	page.Content = pageJSON.Get("extract").String()
	pageLinks := pageJSON.Get("links.#.title").Array()
	for _, link := range pageLinks {
		page.Links = append(page.Links, link.String())
	}
	pageCategories := pageJSON.Get("categories.#.title").Array()
	for _, category := range pageCategories {
		page.Categories = append(page.Categories, category.String())
	}

	return &page, nil
}

// Summary returns the Wikipedia summary for a given search query
func Summary(query string) string {
	// first run Search() to get most relevant result for query
	var searchResults = Search(query)
	v := url.Values{
		"titles":      []string{searchResults[0]},
		"action":      []string{"query"},
		"prop":        []string{"extracts"},
		"exintro":     []string{""},
		"format":      []string{"json"},
		"explaintext": []string{""},
	}

	resp, err := wikiRequest(v)
	if err != nil {
		return ""
	}

	var summaryResult map[string]interface{}
	json.Unmarshal(resp, &summaryResult)

	pages := summaryResult["query"].(map[string]interface{})["pages"]
	page := firstValueInDictionary(pages)
	return page["extract"].(string)
}

// get first value in weird wikipedia pages entry
func firstValueInDictionary(page interface{}) map[string]interface{} {
	for _, v := range page.(map[string]interface{}) {
		return v.(map[string]interface{})
	}
	return nil
}

// Search does a Wikipedia search of query and returns a slice of the result
// titles
func Search(query string) []string {
	v := url.Values{
		"list":     []string{"search"},
		"action":   []string{"query"},
		"srsearch": []string{query},
		"format":   []string{"json"},
	}

	type Page struct {
		PageID uint64 `json:"pageid"`
		Title  string `json:"title"`
	}

	type SearchResult struct {
		Query struct {
			Pages []Page `json:"search"`
		} `json:"query"`
	}

	var results SearchResult

	resp, err := wikiRequest(v)
	if err != nil {
		return []string{}
	}
	json.Unmarshal(resp, &results)

	var titles []string
	for _, page := range results.Query.Pages {
		titles = append(titles, page.Title)
	}

	return titles
}

// wikiRequest hits the wikipedia api with the given request parameters
func wikiRequest(v url.Values) ([]byte, error) {
	res, err := http.Get(apiURL + "?" + v.Encode())
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
