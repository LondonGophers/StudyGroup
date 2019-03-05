package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "fetchall: expects a file containing a list of URLs as the only argument")
		os.Exit(1)
	}

	inf, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}
	defer inf.Close() // NOTE: ignoring error on close

	domains, err := readList(inf, 20)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}

	of, err := os.Create("results.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}
	defer of.Close() // NOTE: ignoring error on close

	fetchAll(domains, 20, of)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}
}

func fetchAll(domains []string, n int, of io.Writer) {

	start := time.Now()
	ch := make(chan string)
	for _, url := range domains {
		go fetch(url, ch)
	}

	for range domains {
		fmt.Fprintln(of, <-ch)
	}

	fmt.Fprintf(of, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if strings.HasPrefix(url, "http://") == false &&
		strings.HasPrefix(url, "https://") == false {
		url = "https://" + url
	}

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func readList(f io.Reader, n int) ([]string, error) {

	domains := make([]string, n)
	count := 0

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF || count >= n {
			break
		}
		if err != nil {
			return nil, err
		}

		domains[count] = strings.TrimRight(line, "\n")
		count++
	}

	if count < n {
		return domains, fmt.Errorf("not enough domains in input file")
	}

	return domains, nil
}
