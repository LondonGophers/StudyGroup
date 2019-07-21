// Change the findlinks program to traverse the `n.FirstChild` linked list using
// recursive calls to `visit` instead of a loop.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for element, count := range visit(nil, doc) {
		fmt.Printf("%s %d\n", element, count)
	}
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
