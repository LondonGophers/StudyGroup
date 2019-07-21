// Modify `dup2` to print the names of all files in which each duplicated line
// occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string)
	done := make(map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, locations, done)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, locations, done)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, locations[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, locations map[string][]string, done map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		file := f.Name()

		counts[line]++
		if !done[line+file] {
			locations[line] = append(locations[line], file)
			done[line+file] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
