// Findlinks1 prints the links in an HTML document read from standard input.
//
// Extend the `visit` function so that it extracts other kinds of links from the document, such as images, scripts, and
// style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {

		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img", "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}

		case "link":
			ss := htmlAttributeSlice(n.Attr).getStylesheet()
			if len(ss) > 0 {
				links = append(links, ss)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

type htmlAttributeSlice []html.Attribute

func (has htmlAttributeSlice) getStylesheet() (href string) {
	isStylesheet := false

	for _, ha := range has {
		if ha.Key == "rel" && ha.Val == "stylesheet" {
			isStylesheet = true
		}

		if ha.Key == "href" {
			href = ha.Val
		}
	}

	if isStylesheet {
		return href
	}

	return ""
}
