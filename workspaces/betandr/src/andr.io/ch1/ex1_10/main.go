// Find a web site that produces a large amount of data. Investigate caching by
// running `fetchall` twice in succession to see whether the reported time changes
// much. Do you get the same content each time? Modify `fetchall` to print its
// output to a file so that it can be examined.

// Try `fetchall` with longer argument lists, such as samples from the top million
// web sites available at `alexa.com`. How does the program behave if a web site
// just doesn't respond? (Section 8.9 describes mechanisms for coping in such
// cases.)
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	// On no response:
	// Get https://qq.com: dial tcp 111.161.64.48:443: i/o timeout

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	addr := url[strings.LastIndex(url, "/")+1:]
	filename := strings.Replace(addr, ".", "_", -1) + ".html"

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename, b, 0644)

	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, len(b), url)
}
