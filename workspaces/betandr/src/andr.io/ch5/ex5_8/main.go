// Modify `forEachNode` so that the `pre` and `post` functions return a boolean
// result indicating whether to continue the traversal. Use it to write a
// function `ElementByID` with the following signature that finds the first
// HTML element with the specific id attribute. The function should stop the
// traversal as soon as a match is found.
//      `func ElementByID(doc *html.Node, id string) *html.Node`
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Run with either a list of URLs, such as `element title https://foo.com https://bar.com`
// or with standard input such as `cat index.html | element title`
func main() {
	if len(os.Args) > 1 {
		for _, url := range os.Args[2:] {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("get %s failed: %s", url, err)
			}
			defer resp.Body.Close()

			doc, err := html.Parse(resp.Body)
			if err != nil {
				log.Fatalf("parse failed: %s", err)
			}
			element := ElementByID(doc, os.Args[1])
			if element != nil {
				fmt.Printf("found: %v\n", element.Data)
			}
		}
	} else {
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			os.Exit(1)
		}
		element := ElementByID(doc, os.Args[1])
		if element != nil {
			fmt.Printf("found: %v\n", element.Data)
		}
	}
}

// ElementByID finds an element `id` in an `html.Node` document and returns the
// `html.Node` matching the `id`.
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement)
}

// forEachNode calls the optional pre(x) function for each node `x` in the tree
// rooted at `n`.
// `pre` is called before the children are visited (preorder)
func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		found := pre(n, id)
		if found {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := forEachNode(c, id, pre)
		if res != nil {
			return res
		}
	}

	return nil
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		if n.Data == id {
			return true
		}
	}
	return false
}
