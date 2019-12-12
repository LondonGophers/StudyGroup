// Charcount computes counts of Unicode characters, letters, digits, and so on in their Unicode categories, using
// functions like `unicode.IsLetter`.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings

	counts := make(map[rune]int) // counts of Unicode characters
	invalid := 0                 // count of invalid UTF-8 characters

	uiKeys := []string{}
	for uiKey := range unicodeIs {
		uiKeys = append(uiKeys, uiKey)
	}

	sort.Strings(uiKeys)

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		for _, uik := range uiKeys {
			if unicodeIs[uik].check(r) {
				u := unicodeIs[uik]
				u.count++
				unicodeIs[uik] = u
			}
		}
	}

	countKeys := []rune{}
	for countKey := range counts {
		countKeys = append(countKeys, countKey)
	}

	sort.Slice(countKeys, func(i, j int) bool { return countKeys[i] < countKeys[j] })
	fmt.Printf("rune\tcount\n")

	for _, ck := range countKeys {
		fmt.Printf("%q\t%d\n", ck, counts[ck])
	}

	fmt.Print("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\n%-11s %9s\n", "unicode.Is*", "count")

	for _, is := range uiKeys {
		fmt.Printf("%-11s %9d\n", is, unicodeIs[is].count)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
