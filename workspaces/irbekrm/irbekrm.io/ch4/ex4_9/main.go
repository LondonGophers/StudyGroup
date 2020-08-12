package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wf := WordFreq("example.txt")
	for word, f := range wf {
		fmt.Printf("%s: %d\n", word, f)
	}
}

// WordFreq accepts file with whitespace separated words and returns a map of word frequencies
func WordFreq(path string) map[string]int {
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
	return wf
}
