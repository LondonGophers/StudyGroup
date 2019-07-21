// The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571,
// one of many favorites. Download each URL (once!) and build an offline index.
// Write a tool xkcd that, using this index, prints the URL and transcript of
// each comic that matches a search term provided on the command line.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	day        int
}

// Exceptionally rudimentary tool that scans JSON every time it's invoked.
// Obviously this has horrible performance compared to a more appropriate index
// but it's a JSON exercise so I'm not over-complicating it. :)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: xkcd {index_dir} {search term}")
	}

	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.FromSlash(fmt.Sprintf("%s/%s", os.Args[1], file.Name()))

		comicFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		var comic Comic
		if err := json.NewDecoder(comicFile).Decode(&comic); err != nil {
			log.Fatal(err)
		}

		if strings.Contains(
			strings.ToLower(comic.Transcript),
			strings.ToLower(os.Args[2])) {
			fmt.Printf("Comic #%d:\n%s\n%s\n\n",
				comic.Num,
				comic.Img,
				comic.Transcript)
		}
	}
}
