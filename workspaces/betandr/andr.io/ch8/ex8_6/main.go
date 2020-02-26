// Add depth-limiting to the concurrent crawler. That is, if the user sets
// `-depth=3`, then only URLs reachable by at most three links will be fetched.
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"gopl.io/ch5/links"
)

var depth = flag.Int("depth", 3, "Only URLs reachable by this number of links will be fetched")
var url = flag.String("url", "", "The URL to start crawling from.")

type link struct {
	URL   string
	Depth int
}

func crawl(url link) []link {
	fmt.Println(url.URL)
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
	wg.Add(1)

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
					wg.Done()
				}(unseenLink)
			}
		}()
	}

	// wait for done and close work channel
	go func() {
		wg.Wait()
		close(worklist)
	}()

	// this feels like a kludge; because of the wg.Add(1) at the start
	// that gives the goroutines time to start adding waitgroup items
	// before the worklist channel is closed we need to remove this
	// waitgroup item. So, this goroutine sleeps for a couple of seconds,
	// to give the others time to start working, then removes the "guard"
	// goroutine.
	go func() {
		time.Sleep(2 * time.Second)
		wg.Done()
	}()

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, l := range list {
			if !seen[l.URL] {
				seen[l.URL] = true
				unseenLinks <- link{l.URL, l.Depth + 1}
				wg.Add(1)
			}
		}
	}
}
