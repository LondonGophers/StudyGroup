// The `strings.NewReader` function returns a value that satisfies the `io.Reader`
// interface (and others) by reading from its argument, a string. Implement a
// simple version of `NewReader` yourself, and use it to make the HTML parser (§5.2)
// take input from a string.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {

	h := "<!DOCTYPE html><html></html>"

	r := NewSimpleReader(h)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for element, count := range visit(nil, doc) {
		fmt.Printf("%s %d\n", element, count)
	}
}

// NewSimpleReader returns a simple `io.Reader`
func NewSimpleReader(s string) *SimpleReader {
	return &SimpleReader{s, 0, -1}
}

// SimpleReader is a simple `io.Reader`
type SimpleReader struct {
	s        string // string to read
	i        int64  // current reading index
	prevRune int    // previous rune index
}

// Read reads string 's' into byte slice 'buf' rune by rune and returns the number
// bytes read from the string and any error.
func (r *SimpleReader) Read(buf []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(buf, r.s[r.i:])
	r.i += int64(n)
	return
}

func visit(elements map[string]int, n *html.Node) map[string]int {
	if elements == nil {
		elements = make(map[string]int)
	}

	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = visit(elements, c)
	}
	return elements
}
