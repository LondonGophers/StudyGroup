// Modify `fetch` to also print the HTTP status code, found in `resp.Status`.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") { // ignoring https ;)
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("GET finished with status: " + resp.Status)

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}
	}
}
