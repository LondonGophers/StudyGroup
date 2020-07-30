package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("  ssl jfgkd  ")
	fmt.Println(b)
	bb := SquashAdjacentSpaces(b)
	fmt.Println(bb)
}

func SquashAdjacentSpaces(b []byte) []byte {
	idx := 0
	ridx := 0
	seenSpace := false
	for ridx < len(b) {
		r, l := utf8.DecodeRune(b[ridx:])
		if unicode.IsSpace(r) {
			if seenSpace {
				// This is a subsequent space
				ridx += l
			} else {
				// The previous char was not a space (if there was one)
				b[idx] = byte(32)
				idx++
				ridx += l
				seenSpace = true
			}
		} else {
			// Not a space char
			for i := 0; i < l; i++ {
				b[idx] = b[ridx+i]
				idx++
			}
			seenSpace = false
			ridx += l
		}
	}
	return b[:idx]
}
