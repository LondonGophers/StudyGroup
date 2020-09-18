package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type movie struct {
	Id string `json:"imdbID"`
}

func main() {
	m := movie{}
	key, ok := os.LookupEnv("OMD_API_KEY")
	if !ok {
		fmt.Fprintln(os.Stderr, "API key not found")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./omdbgo MOVIE")
		os.Exit(1)
	}
	title := os.Args[1]

	fmt.Printf("Retrieving poster for %q\n", title)

	v := url.Values{}
	v.Add("t", title)
	v.Add("apikey", key)
	// Encode values (in case there are space characters in the title)
	ev := v.Encode()

	// First retrieve extra movie data as it seems like the poster API only allows searching with imdbID not title
	url := fmt.Sprintf("http://www.omdbapi.com/?%s", ev)
	b, err := get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving movie data: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling response from OMDb: %v\n", err)
		os.Exit(1)
	}
	url = fmt.Sprintf("http://img.omdbapi.com/?i=%s&apikey=%s", m.Id, key)
	b, err = get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving the poster: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("poster.png", b, 0666)
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("Non-successful status code %d\n", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	resp.Body.Close()
	return b, nil
}
