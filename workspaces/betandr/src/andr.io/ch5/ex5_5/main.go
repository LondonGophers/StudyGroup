// Implement `countWordsAndImages`.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type count struct {
	words  int
	images int
}

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count: %v\n", err)
			continue
		}
		if res.StatusCode != http.StatusOK {
			res.Body.Close()
			fmt.Fprintf(os.Stderr, "%s returned %d\n", url, res.StatusCode)
			continue
		}
		doc, err := html.Parse(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing %s: %s\n", url, err)
			continue
		}
		words, images := CountWordsAndImages(doc)
		fmt.Printf("%s contains %d words within text tags and %d image tags.\n", url, words, images)
	}
}

// CountWordsAndImages traverses an HTML document from node `n` and counts all
// instances of `img` element nodes and counts the number of words in text nodes.
// 'words' is used loosely as some aren't _dictionary_ words. ;)
func CountWordsAndImages(n *html.Node) (words, images int) {
	counts := visit(count{words: 0, images: 0}, n)
	return counts.words, counts.images
}

func visit(wai count, n *html.Node) count {
	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			wai.words += len(strings.Fields(s))
		}
	} else if n.Type == html.ElementNode {
		if n.Data == "img" {
			wai.images++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "script" {
			wai = visit(wai, c)
		}
	}

	return wai
}
