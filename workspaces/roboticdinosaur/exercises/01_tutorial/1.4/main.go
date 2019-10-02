package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineCountsInFiles := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineCountsInFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineCountsInFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
			for filename, count := range lineCountsInFiles[line] {
				fmt.Printf("%s\t, %v\n", filename, count)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineCountsInFiles map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if lineCountsInFiles[input.Text()] == nil {
			lineCountsInFiles[input.Text()] = make(map[string]int)
		}
		lineCountsInFiles[input.Text()][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
