package main

import (
	"crypto/sha256"
	"fmt"
)

func diffBitCount(sha1, sha2 [sha256.Size]byte) byte {
	var diff byte
	for k, v := range sha1 {
		diff += v &^ sha2[k]
	}
	return diff
}

func printer(s, s1 string, diff byte) {
	fmt.Printf("Number of different bits for checksums of %s and %s: %d\n", s, s1, diff)
}
func main() {
	s, s1, s2 := "alpha", "gamma", "alpha"
	diff := diffBitCount(sha256.Sum256([]byte(s)), sha256.Sum256([]byte(s1)))
	printer(s, s1, diff)
	diff1 := diffBitCount(sha256.Sum256([]byte(s)), sha256.Sum256([]byte(s2)))
	printer(s, s2, diff1)
}
