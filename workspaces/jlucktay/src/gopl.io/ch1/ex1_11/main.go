package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	logFile, errOpenLog := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errOpenLog == nil {
		log.SetOutput(logFile)
	} else {
		log.Infof("Failed to log to file, using default stderr: %v", errOpenLog)
	}

	urls := getUrls("alexa.top50.txt")

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		log.Info(<-ch) // receive from channel ch
	}
	log.Infof("%.2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	start := time.Now()
	resp, errGet := http.Get(url)
	if errGet != nil {
		ch <- fmt.Sprint(errGet) // send to channel ch
		return
	}

	nbytes, errCopy := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if errCopy != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, errCopy)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func getUrls(filename string) []string {
	urlBytes, errRead := ioutil.ReadFile(filename)
	if errRead != nil {
		log.Fatalf("error opening '%s': %v", filename, errRead)
	}
	return strings.Split(string(urlBytes), "\n")
}
