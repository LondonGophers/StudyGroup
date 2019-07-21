// Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?

package main

import "fmt"

func reverse(phrase []byte) []byte {
	runes := []rune(string(phrase))
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return []byte(string(runes))
}

func main() {
	phrase := []byte("thgil worg ot nageb yks eht sa ,yllaniF")
	reversed := reverse(phrase)
	fmt.Println(string(reversed))
}
