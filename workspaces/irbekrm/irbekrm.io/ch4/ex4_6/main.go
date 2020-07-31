package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const Space byte = 32

func main() {
	b := []byte("  ssl jfgkd  ")
	fmt.Println(b)
	bb := SquashAdjacentSpaces(b)
	fmt.Println(bb)
}

// SquashAdjacentSpaces squashes space characters (\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP))
// into a single ASCII space
func SquashAdjacentSpaces(b []byte) []byte {
	idx := 0
	ridx := 0
	seenSpace := false
	// loop over runes in the utf-8 encoded byte slice
	for ridx < len(b) {
		r, l := utf8.DecodeRune(b[ridx:])
		if unicode.IsSpace(r) {
			// The previous char was not a whitespace character(if there was one)
			if !seenSpace {
				b[idx] = Space
				idx++
				seenSpace = true
			}
		} else { // Not a witespacespace character
			for i := 0; i < l; i++ {
				b[idx] = b[ridx+i]
				idx++
			}
			seenSpace = false
		}
		// move to the start of next rune
		ridx += l
	}
	return b[:idx]
}
