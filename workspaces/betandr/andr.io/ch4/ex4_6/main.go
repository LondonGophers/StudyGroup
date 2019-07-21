// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
)

// squashAdjacentSpaces removes extra spaces from a byte slice, for example the
// string "hello,         world" would be squashed to "hello, world"
func squashAdjacentSpaces(phrase []byte) []byte {
	runes := []rune(string(phrase)) // it's easier to handle runes!
	i := 0
	notSpace := true
	for _, c := range runes {
		if unicode.IsSpace(c) {
			if notSpace {
				runes[i] = c
				i++
				notSpace = false
			}
		} else {
			runes[i] = c
			i++
			notSpace = true
		}
	}

	return []byte(string(runes[:i])) // turn back into a byte array, as that's requested
}

func main() {
	phrase := []byte("Finally,  as    the sky      began   to    grow           light")
	unspaced := squashAdjacentSpaces(phrase)
	fmt.Println(string(unspaced))
}
