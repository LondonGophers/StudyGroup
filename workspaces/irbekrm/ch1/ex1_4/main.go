package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	appearsIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) < 1 {
		countLines(os.Stdin, counts, appearsIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, appearsIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Counts: %d\tText: %s\tIn files: %v\n", n, line, strings.Join(appearsIn[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, appearsIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		appearsIn[input.Text()] = append(appearsIn[input.Text()], f.Name())
	}
}
