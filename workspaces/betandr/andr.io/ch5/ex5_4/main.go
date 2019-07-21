// Extend the `visit` function so that it extracts other kinds of links from the
// document, such as images, scripts, and style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Links are the types of links in an HTML document
type Links struct {
	Hrefs   []string
	Images  []string
	Scripts []string
	Style   []string
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	links := Links{[]string{}, []string{}, []string{}, []string{}}
	links = visit(links, doc)

	fmt.Println("IMGs:")
	for _, link := range links.Images {
		fmt.Println(link)
	}

	fmt.Println("\nSCRIPTs:")
	for _, link := range links.Scripts {
		fmt.Println(link)
	}

	fmt.Println("\nSTYLEs:")
	for _, link := range links.Style {
		fmt.Println(link)
	}

	fmt.Println("\nHREFs:")
	for _, link := range links.Hrefs {
		fmt.Println(link)
	}
}

func isStylesheet(attrs []html.Attribute) bool {
	for _, a := range attrs {
		if a.Key == "rel" && a.Val == "stylesheet" {
			return true
		}
	}

	return false
}

func visit(links Links, n *html.Node) Links {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links.Hrefs = append(links.Hrefs, a.Val)
				}
			}
		case "img":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links.Images = append(links.Images, a.Val)
				}
			}
		case "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links.Scripts = append(links.Scripts, a.Val)
				}
			}
		case "link":
			if isStylesheet(n.Attr) {
				for _, a := range n.Attr {
					if a.Key == "href" {
						links.Style = append(links.Style, a.Val)
					}
				}
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
