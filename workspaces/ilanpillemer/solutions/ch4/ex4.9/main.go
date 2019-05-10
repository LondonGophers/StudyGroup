package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		counts[word]++
	}

	fmt.Printf("Word\tcount\n")
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
