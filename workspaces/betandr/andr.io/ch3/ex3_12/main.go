// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) > 2 {
		res := "NOT "
		if anagrams(os.Args[1], os.Args[2]) {
			res = ""
		}
		fmt.Printf("%s and %s are %sanagrams\n", os.Args[1], os.Args[2], res)
	}
}

// sortableRunes is a type which implements `Less`, `Swap`, and `Len` to use
// func Sort(data Interface) https://golang.org/pkg/sort/#Sort
type sortableRunes []rune

func (s sortableRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortableRunes) Len() int {
	return len(s)
}

// anagrams checks if two strings are anagrams of each other
// n log n complexity due to two `sort.Sort`s which use quick sort.
func anagrams(s1 string, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	sort.Sort(sortableRunes(r1))
	sort.Sort(sortableRunes(r2))

	if string(r1) == string(r2) {
		return true
	}

	return false
}
