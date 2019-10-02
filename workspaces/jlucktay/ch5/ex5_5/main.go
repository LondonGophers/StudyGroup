// Implement `countWordsAndImages`. (See Exercise 4.9 for word-splitting.)
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a single URL string as an argument!")
		os.Exit(1)
	}

	arg := os.Args[1]

	if !strings.HasPrefix(arg, "http://") && !strings.HasPrefix(arg, "https://") {
		arg = "http://" + arg
	}

	u, err := url.Parse(arg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("URL: %s\n", u)
	w, i, err := CountWordsAndImages(u.String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Words: %d\nImages: %d\n", w, i)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		trimmed := strings.TrimSpace(n.Data)

		in := bufio.NewScanner(strings.NewReader(trimmed))
		in.Split(bufio.ScanWords)

		for in.Scan() {
			words++
		}
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "img":
			images++
		case "script", "style":
			return
		}
	}

	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}

	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}

	return
}
