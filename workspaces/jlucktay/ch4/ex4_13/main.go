package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/jlucktay/golang-workbench/secrets"
)

func main() {
	const queryBase = "http://www.omdbapi.com/?apikey=%s&s=%s&type=movie"

	apiKey := secrets.ReadTokenFromSecrets("omdb-api.json")

	s := url.QueryEscape(strings.Join(os.Args[1:], " "))
	searchURL := fmt.Sprintf(queryBase, apiKey, s)

	fmt.Printf("Search URL: %s\n", strings.ReplaceAll(searchURL, apiKey, "<redacted>"))

	result, errSearch := search(searchURL)
	if errSearch != nil {
		panic(errSearch)
	}

	fmt.Printf("Total results: %d\n\n", result.TotalResults)

	var wg sync.WaitGroup

	posters(result.Movies, &wg)

	countResults := len(result.Movies)
	page := 2

	for countResults < result.TotalResults {
		searchPage := fmt.Sprintf("%s&page=%d", searchURL, page)
		resultPage, errS := search(searchPage)

		if errS != nil {
			panic(errS)
		}

		posters(resultPage.Movies, &wg)

		countResults += len(resultPage.Movies)
		page++
	}

	wg.Wait()
}

func posters(movies []Movie, wg *sync.WaitGroup) {
	for _, movie := range movies {
		fmt.Printf("%s (%d)\n", movie.Title, movie.Year)

		if strings.ToLower(movie.Poster) != "n/a" {
			wg.Add(1)

			go func(m Movie) {
				download(m)
				wg.Done()
			}(movie)
		}
	}
}
