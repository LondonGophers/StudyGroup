// Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func removeAdjacentDupes(strings []string) []string {
	curr := ""
	idx := 0
	for _, str := range strings {
		if str != curr {
			curr = str
			strings[idx] = str
			idx++
		}
	}
	return strings[:idx]
}

func main() {
	s := []string{"Finally", "as", "as", "as", "the", "the", "sky", "began", "to", "to", "grow", "light"}
	deDuped := removeAdjacentDupes(s)
	fmt.Println(deDuped)
}
