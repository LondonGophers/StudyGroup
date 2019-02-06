package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	seen := make(map[string]bool)
	sources := make(map[string][]string)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if !seen[line+filename] {
				seen[line+filename] = true
				sources[line] = append(sources[line], filename)
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s: %d\t%s\n", sources[line], n, line)
		}
	}
}