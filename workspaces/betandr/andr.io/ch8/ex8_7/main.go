// Write a concurrent program that creates a local mirror of a website, fetching
// each reachable page and writing it to a directory on the local disk. Only pages
// within the original domain (for instance `golang.org`) should be fetched. URLs
// within mirrored pages should be altered as needed so that they refer to the
// mirrored page, not the original.
//
// This program is pretty brittle and doesn't really work that well but it does
// explore the ideas of using channels to do some work so I'll leave it there for now.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"gopl.io/ch5/links"
)

var url = flag.String("url", "", "The URL to mirror.")
var out = flag.String("out", "", "The location to write mirror to.")

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

// domain in general cases returns the domain name from a URL
// but also is very brittle so won't stand up to much use.
// Given the string:
// https://www.example.com/something/else.html
// ...it will return www.example.com
func domain(link string) string {
	return strings.Split(link, "/")[2]
}

// download a link given a path
func download(l link, path string) error {
	response, err := http.Get(l.URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	u := strings.Split(l.URL, "/")
	filename := fmt.Sprintf("%s/%s", path, u[len(u)-1])

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}

func main() {
	worklist := make(chan []link)  // lists of URLs, may have duplicates
	unseenPages := make(chan link) // de-duplicated URLs
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
			for unseenPage := range unseenPages {
				go func(l link) {
					if l.Depth <= 1 {
						foundLinks := crawl(l)
						worklist <- foundLinks
					}
					wg.Done()
				}(unseenPage)
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

	// if depth == 1
	// if path not seen
	// send to unseenPages
	seen := make(map[string]bool)
	for list := range worklist {
		for _, l := range list {
			if strings.Contains(domain(l.URL), domain(*url)) {
				fmt.Println("downloading: ", l.URL)
				download(l, *out)
			} else {
				fmt.Println("ignoring: ", l.URL)
			}

			if !seen[l.URL] {
				seen[l.URL] = true
				unseenPages <- link{l.URL, l.Depth + 1}
				wg.Add(1)
			}
		}
	}
}
