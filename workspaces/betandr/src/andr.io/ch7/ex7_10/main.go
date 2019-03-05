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
		if !equals(i, j, s) {
			return false
		}
	}
	return true
}

// equals returns true if the elements at indices `i` and `j` are equal, false otherwise
func equals(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func main() {
	sequence := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v is palindrome: %t\n", sequence, IsPalindrome(sort.IntSlice(sequence)))

	palindrome := []int{0, 1, 2, 3, 4, 4, 3, 2, 1, 0}
	fmt.Printf("%v is palindrome: %t\n", palindrome, IsPalindrome(sort.IntSlice(palindrome)))

	same := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("%v is palindrome: %t\n", same, IsPalindrome(sort.IntSlice(same)))
}
