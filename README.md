# wikigopher
A wrapper around Wikipedia's API for Golang

Inspired by the Wikipedia wrapper for Python (https://github.com/goldsmith/Wikipedia)

**********

Start using wikipedia for Go in less than 5 minutes!

Begin by installing wikigopher::

	$ go get github.com/Ishaangandhi/wikigopher

## Features
First, import wikigopher:

```go
import "github.com/ishaangandhi/wikigopher"
```

#### WikipediaPage Struct
WikipediaPage represents a Wikipedia Page

```Go
type WikipediaPage struct {
	PageID     uint64
	Title      string
	Content    string
	Categories []string
	Links      []string
}
```

To get a `WikipediaPage` just use the `Page()` function with a search query. For example:

```Go
page, err := wikigopher.Page("nyc")
fmt.Println(page.Title)
// prints "New York City"

fmt.Println(page.Content)
// prints "The City of New York, usually called either New York
// City (NYC) or simply New York (NY)"...

fmt.Println(page.Links)
// prints "[10 Hudson Yards 110th Street (Manhattan) 120th Street...]"

```

#### Search function
Search does a Wikipedia search of query and returns a slice of the result titles

```Go
func main() {
	var searchResults []string

	searchResults = wikigopher.Search("barack obama")
	fmt.Println(searchResults)
	// prints [Barack Obama Barack Obama Sr. Family of Barack Obama...]

	searchResults = wikigopher.Search("golang")
	fmt.Println(searchResults)
	// prints [Go (programming language) Data mapper pattern Template..]
}
```

#### Summary function
Summary returns the Wikipedia summary for a given search query

```Go
func main() {
	var summary string

	summary = wikigopher.Summary("barack obama")
	fmt.Println(summary)
	// prints: "Barack Hussein Obama II ( (listen); born August 4, 1961) is
	// an American attorney..."

	summary = wikigopher.Summary("golang")
	fmt.Println(summary)
	// prints: "Go (also referred to as Golang) is a statically typed,
	// compiled programming language..."
}
```
