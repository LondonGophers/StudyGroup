package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprint(os.Stderr, "No files provided")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			hasDupl := countLines(f, counts)
			if hasDupl {
				fmt.Println(arg)
			}
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	hasDupl := false
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			hasDupl = true
		}

	}
	return hasDupl
}
