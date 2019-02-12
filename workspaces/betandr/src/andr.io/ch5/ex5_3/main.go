// Write a function to print the contents of all text nodes in an HTML document
// tree. Do not descend into `<script>` or `<style>` elements, since their
// contents are not visible in a web browser.

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(text []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			text = append(text, s)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "script" {
			text = visit(text, c)
		}
	}

	return text
}
