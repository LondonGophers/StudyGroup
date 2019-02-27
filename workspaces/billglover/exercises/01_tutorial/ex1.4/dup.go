package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)
	locs := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", locs)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %s: %v\n", arg, err)
				continue
			}
			countLines(f, counts, arg, locs)
			err = f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %s: %v\n", arg, err)
				continue
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for foundIn := range locs[line] {
				fmt.Printf("  %s", foundIn)
			}
			fmt.Println()
		}
	}
}

// CountLines counts the number of times a line appears in an io.Reader. It
// stores the line count in the provided counts map. It keeps track of the
// location a line was found (the loc parameter). If location tracking is not
// required, pass a nil locs map.
func countLines(r io.Reader, counts map[string]int, loc string, locs map[string]map[string]bool) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		counts[line]++

		if locs == nil {
			continue
		}

		if locs[line] == nil {
			locs[line] = make(map[string]bool)
		}
		locs[line][loc] = true
	}
}
