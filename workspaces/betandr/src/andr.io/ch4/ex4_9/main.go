// Write a program `wordfreq` to report the frequency of each word in an input
// text file. Call `input.Split(bufio.ScanWords)` before the first call to
// `Scan` to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	freqs := make(map[string]int)

	for scanner.Scan() {
		freqs[scanner.Text()]++
	}

	fmt.Print("\nfreq\tword\n")
	for i, f := range freqs {
		fmt.Printf("%d\t%s\n", f, i)
	}

}
