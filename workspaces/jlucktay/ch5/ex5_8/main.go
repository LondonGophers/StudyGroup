// Exercise 5.8
//
// Modify `forEachNode` so that the `pre` and `post` functions return a boolean result indicating whether to continue
// the traversal. Use it to write a function `ElementByID` with the following signature that finds the first HTML
// element with the specified `id` attribute. The function should stop the traversal as soon as a match is found.
//
//		func ElementById(doc *html.Node, id string) *html.Node

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func ElementById(doc *html.Node, id string) *html.Node {
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node x in the tree rooted at n.
// Both functions are optional.
// pre is called before the children are visited (preorder) and post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if post(n) {
			return
		}
	}
}

var depth int

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}

	return
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}

	return
}
