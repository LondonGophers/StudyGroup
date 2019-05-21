//If a website does not respond, the program hangs for 30s then throws a timeout error
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	//"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	sites := []string{"google.com","facebook.com","amazon.com","steph-dmts.tech","www.codecademy.com"}
	for _, url := range  sites{
		if !strings.HasPrefix(url,"http://") {
			url = "http://"+url
		}
		go fetch(url, ch) 
	}
	for range sites {
		fmt.Println(<-ch) 
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
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
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
