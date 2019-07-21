// Write a function that counts the number of bits that are different in two
// SHA256 hashes. (see Popcount from Section 2.6.2)
package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"andr.io/ch4/ex4_1/popcount"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: shapc {string} {string}")
		os.Exit(0)
	}

	s1 := sha256.Sum256([]byte(os.Args[1]))
	s2 := sha256.Sum256([]byte(os.Args[2]))

	diffCount := 0

	for _, n := range s1 {
		diffCount += popcount.Count(uint64(n))
	}

	for _, n := range s2 {
		diffCount -= popcount.Count(uint64(n))
	}

	fmt.Printf("%d bits different\n", abs(diffCount))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
