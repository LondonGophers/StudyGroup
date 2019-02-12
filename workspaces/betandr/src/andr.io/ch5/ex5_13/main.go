// Modify crawl to make local copies of the pages it finds, creating directories
// as necessary. Don’t make copies of pages that come from a different domain.
// For example, if the original page comes from golang.org, save all files from
// there, but exclude ones from vimeo.com.
package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item, baseUrl string) []string, baseURL string) {
	worklist := []string{baseURL}
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, baseURL)...)
			}
		}
	}
}

// isSameDomain returns true if the domains names match or false otherwise
// https://example.com/foo/bar/baz matches http://example.com/quux, however
// https://foo.example.com _does not_ match https://example.com
// This may be an issue but for simplicity it's ok.
func isSameDomain(url1, url2 string) bool {
	u1, _ := url.Parse(url1)
	u2, _ := url.Parse(url2)

	return u1.Host == u2.Host
}

func crawl(u, baseURL string) []string {

	if !isSameDomain(u, baseURL) {
		fmt.Printf("ignoring: %s\n", u)

	} else {
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			fmt.Errorf("error creating request to %s:%s", u, err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Errorf("error requesting %s:%s", u, err)
		}

		if res.StatusCode != http.StatusOK {
			res.Body.Close()
			fmt.Errorf("GET failed: %s", res.Status)
		}

		idx := strings.Index(baseURL, "//")
		site := strings.Trim(baseURL[idx:], "/")
		dir, _ := os.Getwd()
		path := fmt.Sprintf("%s/%s%s", dir, site, html.EscapeString(req.URL.Path))

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Errorf("error reading response body: %v", err)
		}

		os.MkdirAll(path, os.ModePerm)

		fmt.Printf("saving: %s to %s\n", u, path)

		// Hard-coding index.html is a bit bad as we're assuming it's all HTML but
		// it's ok as a proof of concept.
		file := path + "index.html"
		err = ioutil.WriteFile(file, body, 0644)
		if err != nil {
			fmt.Errorf("error writing file %s: %v", path, err)
		}

		res.Body.Close()
	}

	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Crawl the web breadth-first and make a local copy of the page
	breadthFirst(crawl, os.Args[1])
}
