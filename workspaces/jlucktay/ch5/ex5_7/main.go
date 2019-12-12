// Exercise 5.7
//
// Develop `startElement` and `endElement` into a general HTML pretty-printer. Print comment nodes, text nodes, and the
// attributes of each element(`<a href='...'>`). Use short forms like `<img/>` instead of `<img></img>` when an element
// has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)

// Outline prints the outline of an HTML document tree.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := outline(url); err != nil {
			panic(err)
		}
	}
}

func outline(url string) error {
	resp, errGet := http.Get(url)
	if errGet != nil {
		return errGet
	}
	defer resp.Body.Close()

	if errProcess := process(resp.Body); errProcess != nil {
		return errProcess
	}

	return nil
}

func process(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node x in the tree rooted at n.
// Both functions are optional.
// pre is called before the children are visited (preorder) and post is called after (postorder).
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
	switch n.Type {
	case html.DoctypeNode:
		fmt.Printf("<!doctype %v>\n", n.Data)
	case html.CommentNode:
		fmt.Printf("%*s<!-- %v -->\n", depth*2, "", n.Data)
	case html.TextNode:
		trimmed := strings.TrimSpace(n.Data)
		if len(trimmed) > 0 {
			scanner := bufio.NewScanner(strings.NewReader(n.Data))
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				text := scanner.Text()
				fmt.Printf("%*s%s\n", depth*2, "", text)
			}
		}
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)

		for _, a := range n.Attr {
			fmt.Printf(` %s="%s"`, a.Key, a.Val)
		}

		if n.FirstChild == nil {
			fmt.Println(" />")
		} else {
			fmt.Println(">")
			depth++
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
