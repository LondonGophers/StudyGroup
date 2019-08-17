// Package reverse contains a func to reverse a slice.
package reverse

import (
	"fmt"
	"unicode/utf8"
)

// Reverse reverses a slice of ints in place.
func Reverse(a *[]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

// ReverseString is a modified version of Reverse which takes a slice of bytes representing a UTF-8-encoded string and
// reverses it in place.
func ReverseString(b *[]byte) {
	if !utf8.Valid(*b) {
		panic(fmt.Sprintf("input '% x' does not consist entirely of valid UTF-8-encoded runes", *b))
	}

	runesRemaining := utf8.RuneCount(*b)
	reversedSize := 0

	for runesRemaining > 0 {
		// Get the rune currently at the start of the slice
		r, size := utf8.DecodeRune(*b)
		if r == utf8.RuneError {
			panic("got RuneError")
		}

		// Bring forward the remainder of the slice
		copy((*b)[:], (*b)[size:len(*b)-reversedSize])

		// Keep track of the number of bytes that have been reversed
		reversedSize += size

		// Put the rune from the start at the end
		xr := []byte(string(r))
		for xri := 0; xri < len(xr); xri++ {
			(*b)[len(*b)-reversedSize+xri] = xr[xri]
		}

		runesRemaining--
	}
}
