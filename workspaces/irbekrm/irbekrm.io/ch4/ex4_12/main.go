package main

import (
	"fmt"
	"os"
)

type comic struct {
	Url        string `json:"img"`
	Transcript string `json:"transcript"`
	Num        int    `json:"num"`
}

const currentUrl string = "https://xkcd.com/info.0.json"

// https://xkcd.com/json.html
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage ./xkcdgo populate|search TERM")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "populate":
		if err := populate(); err != nil {
			fmt.Fprintf(os.Stderr, "Error populating data: %v\n", err)
		}
	case "search":
		if err := search(); err != nil {
			fmt.Fprintf(os.Stderr, "Error searching: %v\n", err)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
