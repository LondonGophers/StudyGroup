package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

var protocolPrefix = regexp.MustCompile(`^https?:\/\/`)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	fileName := fileNameFromUrl(url)
	file, err := createFile(fileName)
	if err != nil {
		ch <- fmt.Sprintf("Error creating file %s: %v\n", fileName, err)
	}
	defer file.Close()
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func fileNameFromUrl(url string) string {
	fileName := protocolPrefix.ReplaceAll([]byte(url), []byte(""))
	return string(fileName)
}

func createFile(fileName string) (*os.File, error) {
	err := os.MkdirAll("fetchall", os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Error creating dir: %v\n", err)
	}
	filePath := "fetchall/" + fileName
	file, err := os.Create(filePath)
	return file, err
}
