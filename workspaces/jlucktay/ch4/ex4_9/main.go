// Write a program `wordfreq` to report the frequency of each word in an input text file. Call
// `input.Split(bufio.ScanWords)` before the first call to `Scan` to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of words

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()

		counts[word]++
	}

	fmt.Printf("%-40s count\n", "word")

	for word, count := range counts {
		fmt.Printf("%-40s %d\n", word, count)
	}
}
