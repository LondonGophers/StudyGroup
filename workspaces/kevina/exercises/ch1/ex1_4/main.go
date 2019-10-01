package main

import (
	"bufio"
	"fmt"
	"os"
)

type instance struct {
	files  map[string]int
	number int
}

func main() {
	counts := make(map[string]instance)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.number > 1 {
			fmt.Printf("%d\t%s\t%v\n", n.number, line, n.files)
		}
	}
}

func countLines(f *os.File, counts map[string]instance) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		files := counts[line].files
		if files == nil {
			files = make(map[string]int)
		}
		files[f.Name()]++
		count := counts[line].number
		count++
		val := instance{
			files:  files,
			number: count,
		}

		counts[line] = val
	}
}
