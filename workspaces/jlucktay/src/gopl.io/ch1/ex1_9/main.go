package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, errGet := http.Get(url)
		if errGet != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", errGet)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Printf("HTTP response status: (%d) %s\n", resp.StatusCode, resp.Status)

		_, errCopy := io.Copy(os.Stdout, resp.Body)
		if errCopy != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, errCopy)
			os.Exit(1)
		}
	}
}
