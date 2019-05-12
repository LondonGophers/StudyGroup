package main

import (
	"bufio"
	"fmt"
	"os"
)

var stop = make(map[string]struct{})

func init() {
	f, err := os.Open("STOPWORDS")
	if err != nil {
		panic(err)
	}
	process := bufio.NewScanner(f)
	process.Split(bufio.ScanWords)
	for process.Scan() {
		w := process.Text()
		stop[w] = struct{}{}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	counts := make(map[string]int)
	for input.Scan() {

		if err := input.Err(); err != nil {
			fmt.Printf("Oh no... %v\n", err)
			os.Exit(1)
		}

		word := input.Text()
		//ignore stop words
		if _, ok := stop[word]; ok {
			continue
		}

		counts[word]++
	}

	fmt.Printf("Word\tcount\n")
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
