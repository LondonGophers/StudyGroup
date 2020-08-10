package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	CharCount(os.Stdin, os.Stdout)
}

func CharCount(stdin io.Reader, stdout io.Writer) {
	types := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
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
		fmt.Fprintf(stdout, "%-15v%v\n", "type", "count")
		for t, c := range types {
			fmt.Fprintf(stdout, "%-15s%d\n", t, c)
		}
		if invalid > 0 {
			fmt.Fprintf(stdout, "\n%d invalid UTF-8 characters\n", invalid)
		}
	}
}
