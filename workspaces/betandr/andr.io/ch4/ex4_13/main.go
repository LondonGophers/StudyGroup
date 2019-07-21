// The JSON-based web service of the Open Movie Database lets you search
// https://omdbapi.com/ for a movie by name and download its poster image.
// Write a tool poster that downloads the poster image for the movie named
// on the command line.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// OMDBURL is the OMDB URL
const OMDBURL = "http://www.omdbapi.com"

type Search struct {
	Films []*Film `json:"Search"`
}

// Film is returned by an OMDB search
type Film struct {
	Title  string
	Year   string
	ID     string `json:"imdbID"`
	Type   string
	Poster string
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: poster {name}")
		os.Exit(1)
	}

	key := os.Getenv("OMDB_API_KEY")
	if len(key) <= 0 {
		fmt.Println("No API key found. Run `export OMDB_API_KEY=xxxxx`")
		os.Exit(1)
	}

	params := url.Values{}
	params.Set("s", os.Args[1])
	params.Set("apikey", key)

	url := fmt.Sprintf("%s?%s", OMDBURL, params.Encode())

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Search failed: %s", res.Status)
	}

	var search Search
	if err := json.NewDecoder(res.Body).Decode(&search); err != nil {
		res.Body.Close()
		log.Fatal(err)
	}

	for _, film := range search.Films {
		if strings.HasPrefix(film.Poster, "http") {
			response, e := http.Get(film.Poster)
			if e != nil {
				log.Fatal(e)
			}

			file, err := os.Create(fmt.Sprintf("%s.jpg", film.ID))
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}

			response.Body.Close()
			file.Close()
		}
	}
}
