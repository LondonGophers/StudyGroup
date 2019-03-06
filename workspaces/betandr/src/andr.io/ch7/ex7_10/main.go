// The `sort.Interface` type can be adapted to other uses. Write a function
// `IsPalindrome(s sort.Interface) bool` that reports whether the sequence `s` is a
// palindrome, in other words, reversing the sequence would not change it. Assume
// that the elements at indices `i` and `j` are equal if
// `!s.Less(i, j) && !s.Less(j, i)`.
package main

import (
	"fmt"
	"sort"
)

// IsPalindrome reports whether the sequence `s` is a  palindrome, in other words,
// reversing the sequence would not change it
func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < s.Len()-1/2; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			return true
		}
	}
	return false
}

func main() {
	sequence := sort.StringSlice([]string{"p", "a", "l", "a", "m", "i", "n", "o"})
	fmt.Printf("%v => %t\n", sequence, IsPalindrome(sequence))

	palindrome := sort.StringSlice([]string{"t", "a", "c", "o", "c", "a", "t"})
	fmt.Printf("%v   => %t\n", palindrome, IsPalindrome(palindrome))

	repeated := sort.StringSlice([]string{"x", "x", "x", "x", "x", "x", "x", "x"})
	fmt.Printf("%v => %t\n", repeated, IsPalindrome(repeated))
}
