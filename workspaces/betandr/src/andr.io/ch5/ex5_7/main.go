// Develop `startElement` and `endElement` into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element
// (`<a href='...'>`). Use short forms like `<img/>` instead of `<img></img>`
// when an element has no children. Write a test to ensure that the output can
// be parsed successfully.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Run with either a list of URLs, such as `outline https://foo.com https://bar.com`
// or with standard input such as `cat index.html | outline`
func main() {
	if len(os.Args) > 1 {
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
			outline(doc)
		}
	} else {
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			os.Exit(1)
		}
		outline(doc)
	}
}

func outline(doc *html.Node) error {
	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*4, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Printf(">\n")
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*4, "", n.Data)
	}
}
