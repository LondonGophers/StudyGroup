package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	CharCount()
}

func CharCount() {
	types := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			fmt.Print("Good bye")
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "CharCount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			types["letters"]++
		case unicode.IsDigit(r):
			types["digits"]++
		case unicode.IsSpace(r):
			types["whitespace"]++
		default:
			types["other"]++
		}
		fmt.Printf("character\ttype\tcount\n")
		for c, n := range types {
			fmt.Printf("%q\t%d\n", c, n)
		}
		fmt.Print("\nlen\tcount\n")
		if invalid > 0 {
			fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
		}
	}
}
