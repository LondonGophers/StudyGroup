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

const (
	IsControl int = iota
	IsDigit
	IsGraphic
	IsLetter
	IsLower
	IsMark
	IsNumber
	IsPrint
	IsPunct
	IsSpace
	IsSymbol
	IsTitle
	IsUpper
)

var isKeyNames = [...]string{
	"IsControl",
	"IsDigit",
	"IsGraphic",
	"IsLetter",
	"IsLower",
	"IsMark",
	"IsNumber",
	"IsPrint",
	"IsPunct",
	"IsSpace",
	"IsSymbol",
	"IsTitle",
	"IsUpper",
}

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	isUnicode := make(map[int]uint) // count of rune types per unicode.Is* funcs

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

		if unicode.IsControl(r) {
			isUnicode[IsControl]++
		}
		if unicode.IsDigit(r) {
			isUnicode[IsDigit]++
		}
		if unicode.IsGraphic(r) {
			isUnicode[IsGraphic]++
		}
		if unicode.IsLetter(r) {
			isUnicode[IsLetter]++
		}
		if unicode.IsLower(r) {
			isUnicode[IsLower]++
		}
		if unicode.IsMark(r) {
			isUnicode[IsMark]++
		}
		if unicode.IsNumber(r) {
			isUnicode[IsNumber]++
		}
		if unicode.IsPrint(r) {
			isUnicode[IsPrint]++
		}
		if unicode.IsPunct(r) {
			isUnicode[IsPunct]++
		}
		if unicode.IsSpace(r) {
			isUnicode[IsSpace]++
		}
		if unicode.IsSymbol(r) {
			isUnicode[IsSymbol]++
		}
		if unicode.IsTitle(r) {
			isUnicode[IsTitle]++
		}
		if unicode.IsUpper(r) {
			isUnicode[IsUpper]++
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

	isKeys := []int{}
	for isKey := range isUnicode {
		isKeys = append(isKeys, isKey)
	}
	sort.Ints(isKeys)
	fmt.Printf("\n%-10s\tcount\n", "is iota")
	for _, is := range isKeys {
		fmt.Printf("%-10s\t%5d\n", isKeyNames[is], isUnicode[is])
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
