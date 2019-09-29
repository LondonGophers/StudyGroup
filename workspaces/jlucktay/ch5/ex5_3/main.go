// Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into `<script>`
// or `<style>` elements, since their contents are not visible in a web browser.
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
		fmt.Fprintf(os.Stderr, "printnodes: %v\n", err)
		os.Exit(1)
	}

	printnodes(doc)
}

// printnodes recursively prints the content of text nodes.
func printnodes(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	if n.Type == html.TextNode {
		trimmed := strings.TrimSpace(n.Data)

		if len(trimmed) > 0 {
			fmt.Println(trimmed)
		}
	}

	if n.FirstChild != nil {
		printnodes(n.FirstChild)
	}

	if n.NextSibling != nil {
		printnodes(n.NextSibling)
	}
}
