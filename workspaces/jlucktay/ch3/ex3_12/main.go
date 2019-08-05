// Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters
// in a different order.
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		if strings.IndexRune(os.Args[i], '/') != -1 {
			xs := strings.Split(os.Args[i], "/")
			fmt.Printf("  '%s' v '%s': %v\n", xs[0], xs[1], Anagram(xs[0], xs[1]))
		}
	}
}

func Anagram(s1, s2 string) bool {
	characters := make(map[rune]uint)

	for _, r := range []rune(s1) {
		if !unicode.IsSpace(r) {
			characters[unicode.ToLower(r)]++
		}
	}

	for _, r := range []rune(s2) {
		if !unicode.IsSpace(r) {
			if count, exists := characters[unicode.ToLower(r)]; !exists || count == 0 {
				return false
			}

			characters[unicode.ToLower(r)]--
		}
	}

	for _, count := range characters {
		if count != 0 {
			return false
		}
	}

	return true
}
