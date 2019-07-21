// Write a variadic function `ElementsByTagName` that, given an HTML node tree
// and zero or more names, returns all the elements that match one of those
// names. Here are two example calls:
//     `func ElementsByTagName(doc *html.Node, name ...string) []*html.Node`
//     `images := ElementsByTagName(doc, "img")``
//     `headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")``
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Run with a list of URLs, such as `ebtn https://foo.com https://bar.com`
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("get %s failed: %s", url, err)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatalf("parse failed: %s", err)
		}

		// find image tags
		images := ElementsByTagName(doc, "img")
		fmt.Printf("found %d images\n", len(images))
		for _, image := range images {
			fmt.Printf("%v\n", renderNode(image))
		}

		// find heading tags
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
		fmt.Printf("found %d headings\n", len(headings))
		for _, heading := range headings {
			fmt.Printf("%v\n", renderNode(heading))
		}
	}
}

// renderNode returns a HTML string representation of the `html.Node` supplied
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

// ElementsByTagName returns all the elements of an an HTML node tree that match
// one of zero or more names given.
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var ids = make(map[string]bool)
	for _, n := range name { // turn our slice of strings into a map for speed
		ids[n] = true
	}

	return visit([]*html.Node{}, doc, ids)
}

// visit checks if the current node matches any supplied `ids` then adds it to
// nodes slice which is passed back into the visit function as it traverses the
// `n` node tree.
func visit(nodes []*html.Node, n *html.Node, ids map[string]bool) []*html.Node {

	if n.Type == html.ElementNode {
		if ids[n.Data] {
			nodes = append(nodes, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = visit(nodes, c, ids)
	}

	return nodes
}
