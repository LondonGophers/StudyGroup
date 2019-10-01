// Write a function to populate a mapping from element names—`p`, `div`, `span`, and so on—to the number of elements
// with that name in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mapnodes: %v\n", err)
		os.Exit(1)
	}

	nodeCount := make(map[string]int)

	mapnodes(nodeCount, doc)

	for name, count := range nodeCount {
		fmt.Printf("%10s: %d\n", name, count)
	}
}

// mapnodes tallies a count of element node occurrences by name in the given map.
func mapnodes(nodeCount map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		nodeCount[n.Data]++
	}

	if n.FirstChild != nil {
		mapnodes(nodeCount, n.FirstChild)
	}

	if n.NextSibling != nil {
		mapnodes(nodeCount, n.NextSibling)
	}
}
