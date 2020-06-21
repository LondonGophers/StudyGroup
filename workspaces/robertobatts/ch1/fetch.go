package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch(url string) {
	adHTTPPrefix(&url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch reading: %v\n", err)
		os.Exit(1)
	}
}

func adHTTPPrefix(url *string) {
	prefix := "http://"
	if !strings.HasPrefix(*url, prefix) {
		*url = prefix + *url
	}
}
