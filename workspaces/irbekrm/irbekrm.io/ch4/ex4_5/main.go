package main

import "fmt"

func main() {
	s := []string{"alpha", "beta", "beta", "beta", "gamma", "gamma", "gamma"}
	s = eliminateAdjacentDuplicates(s)
	fmt.Println(s) // [alpha beta gamma]
}

func eliminateAdjacentDuplicates(s []string) []string {
	// for each slice item at index i
	for i := 0; i < len(s); i++ {
		ii := i + 1
		// check if item at i+1 is a duplicate
		// if a duplicate found at i+1, delete it, by shifting items on the right by 1 forward and repeat for the new item at i+1
		// TODO: if multiple adjacent duplicates found, eliminate them all together by one shifting the remaining items only once
		for {
			// if we have reached the end of the slice, return
			if ii >= len(s) {
				return s
			}
			// if the item at index i+1 is not a duplicate, break
			if s[i] != s[ii] {
				break
			}
			// if the item at i+1 is a duplicate and is also the last element, return the array without it
			if ii == len(s)-1. {
				return s[:len(s)-1]
			}
			// if duplicate was found and is not the last element
			// delete it by shifting items on the right from the duplicate forward by 1
			s = append(s[:ii], s[ii+1:]...)
		}
	}
	return s
}
