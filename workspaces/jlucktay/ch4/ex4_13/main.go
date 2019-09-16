package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	const queryBase = "http://www.omdbapi.com/?apikey=%s&s=%s&type=movie"

	s := url.QueryEscape(strings.Join(os.Args[1:], " "))
	searchURL := fmt.Sprintf(queryBase, APIkey, s)

	fmt.Printf("Search URL: %s\n", strings.ReplaceAll(searchURL, APIkey, "<redacted>"))

	result, errSearch := search(searchURL)
	if errSearch != nil {
		panic(errSearch)
	}
	fmt.Printf("Total results: %d\n", result.TotalResults)

	posters(result.Movies)

	countResults := len(result.Movies)
	page := 2

	for countResults < result.TotalResults {
		searchPage := fmt.Sprintf("%s&page=%d", searchURL, page)
		resultPage, errS := search(searchPage)
		if errS != nil {
			panic(errS)
		}

		posters(resultPage.Movies)

		countResults += len(resultPage.Movies)
		page++
	}
}

func posters(movies []Movie) {
	for _, movie := range movies {
		fmt.Printf("%s (%d)\n\t%s\n\n", movie.Title, movie.Year, movie.Poster)
	}
}
