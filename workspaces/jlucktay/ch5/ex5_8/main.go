// Exercise 5.8
//
// Modify `forEachNode` so that the `pre` and `post` functions return a boolean result indicating whether to continue
// the traversal. Use it to write a function `ElementByID` with the following signature that finds the first HTML
// element with the specified `id` attribute. The function should stop the traversal as soon as a match is found.
//
//		func ElementByID(doc *html.Node, id string) *html.Node

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("You must specify at least 2 arguments:")
		fmt.Println("1) the ID of an HTML node")
		fmt.Println("2) one or more URLs to fetch and parse")
		os.Exit(1)
	}

	for _, url := range os.Args[2:] {
		doc, errFetch := fetch(url)
		if errFetch != nil {
			log.Fatalf("Error fetching '%s': %v", url, errFetch)
		}

		ele := ElementByID(doc, os.Args[1])
		if ele != nil {
			fmt.Printf("Found element at '%s' with attributes: '%+v'\n", url, ele.Attr)
		} else {
			fmt.Printf("Element with ID '%s' not found at '%s'.\n", os.Args[1], url)
		}
	}
}

func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, checkElementID)
}

// forEachNode calls the function pre(x) for each node x in the tree rooted at n.
// This function is optional.
// pre is called before the children are visited (preorder).
func forEachNode(n *html.Node, id string, pre func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		recurse := forEachNode(c, id, pre)
		if recurse != nil {
			return recurse
		}
	}

	return nil
}

func checkElementID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if strings.ToLower(a.Key) == "id" && strings.EqualFold(a.Val, id) {
				return true
			}
		}
	}

	return false
}
