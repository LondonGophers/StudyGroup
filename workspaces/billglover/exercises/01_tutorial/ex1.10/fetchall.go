package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	of, err := os.Create("results.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintln(of, <-ch)
	}

	fmt.Fprintf(of, "%.2fs elapsed\n", time.Since(start).Seconds())

	err = of.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
