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

#### Search command
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

#### Summary command
Summary returns the Wikipedia summary for a given search query

```Go
func main() {
	var summary string

	summary = wikigopher.Summary("barack obama")
	fmt.Println(summary)
	// prints: "Barack Hussein Obama II ( (listen); born August 4, 1961) is
	// an American attorney..."

	summary = wikigopher.Search("golang")
	fmt.Println(summary)
	// prints: "Go (also referred to as Golang) is a statically typed,
	// compiled programming language..."
}
```
