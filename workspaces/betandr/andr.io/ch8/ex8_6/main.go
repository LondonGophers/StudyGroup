// Add depth-limiting to the concurrent crawler. That is, if the user sets
// `-depth=3`, then only URLs reachable by at most three links will be fetched.
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"gopl.io/ch5/links"
)

var depth = flag.Int("depth", 3, "Only URLs reachable by this number of links will be fetched")
var url = flag.String("url", "", "The URL to start crawling from.")

type link struct {
	URL   string
	Depth int
}

func crawl(url link) []link {
	fmt.Println(url)
	list, err := links.Extract(url.URL)
	if err != nil {
		log.Print(err)
	}

	links := make([]link, len(list))
	for i, l := range list {
		links[i] = link{URL: l, Depth: url.Depth}
	}
	return links
}

func main() {
	worklist := make(chan []link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan link) // de-duplicated URLs
	var wg sync.WaitGroup

	flag.Parse()

	// Start from specified url
	start := make([]link, 1)
	start[0] = link{URL: *url, Depth: 0}
	go func() { worklist <- start }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for unseenLink := range unseenLinks {
				go func(l link) {
					if l.Depth <= *depth {
						foundLinks := crawl(l)
						worklist <- foundLinks
					}
				}(unseenLink)
			}
		}()
	}

	// wait for done and close work channel
	go func() {
		wg.Wait()
		close(worklist)
	}()

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, l := range list {
			if !seen[l.URL] {
				seen[l.URL] = true
				wg.Add(1)
				unseenLinks <- link{l.URL, l.Depth + 1}
			}
		}
	}
}
