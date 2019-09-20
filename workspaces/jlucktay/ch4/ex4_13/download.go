package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func download(m Movie) {
	xURL := strings.Split(m.Poster, ".")
	extension := xURL[len(xURL)-1]
	filename := fmt.Sprintf("%s.%d.%s", m.Title, m.Year, extension)

	if _, errStat := os.Stat(filename); errStat == nil {
		fmt.Printf("[Download] '%s' already exists!\n", filename)
		return
	} else if os.IsNotExist(errStat) {
		fmt.Printf("[Download] Started for %s (%d) to '%s'.\n", m.Title, m.Year, filename)
	} else {
		panic(errStat)
	}

	out, errCreate := os.Create(filename)
	if errCreate != nil {
		panic(errCreate)
	}
	defer out.Close()

	resp, errGet := http.Get(m.Poster)
	if errGet != nil {
		panic(errGet)
	}
	defer resp.Body.Close()

	resp.Header.Get(http.CanonicalHeaderKey("content-type"))

	n, errCopy := io.Copy(out, resp.Body)
	if errCopy != nil {
		panic(errCopy)
	}

	fmt.Printf("[Download] Finished for %s (%d) to '%s'. (%d byte(s))\n", m.Title, m.Year, filename, n)
}
