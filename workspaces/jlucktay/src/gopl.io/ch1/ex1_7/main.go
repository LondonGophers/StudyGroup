package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, errGet := http.Get(url)
		if errGet != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", errGet)
			os.Exit(1)
		}
		defer resp.Body.Close()

		_, errCopy := io.Copy(os.Stdout, resp.Body)
		if errCopy != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, errCopy)
			os.Exit(1)
		}
	}
}
