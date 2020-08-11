package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	WordFreq("example.txt", os.Stdin)
}

// WordFreq reads words from input file, calculates frequency and writes output to
// writer accepted as parameter
func WordFreq(path string, w io.Writer) {
	wf := make(map[string]int)
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		wf[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
	for word, f := range wf {
		fmt.Fprintf(w, "%s: %d\n", word, f)
	}
}
