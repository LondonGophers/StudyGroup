package main

import (
	"bufio"
	"fmt"
	"os"
)

func spotDuplicateLines1() {
	counts := make(map[string]int)
	filePaths := os.Args[1:]
	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "filePath: %s, error: %v\n", filePath, err)
			continue
		}
		countLines(file, counts)
		file.Close()
	}
	printDuplicates(counts)
}

func spotDuplicateLines2() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	printDuplicates(counts)
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printDuplicates(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
