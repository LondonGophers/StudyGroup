package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]
	var dupFiles []string
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			hasDup := false
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			for _, n := range counts {
				if n > 1 {
					hasDup = true
				}
			}
			if hasDup == true {
				dupFiles = append(dupFiles, arg)
			}
			f.Close()
		}
	}
	fmt.Println("Files that have dup lines:")
	for _, f := range dupFiles {
		fmt.Println(f)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
