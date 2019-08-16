// Package adjacent contains the Eliminate() and EliminateSpaces() funcs.
package adjacent

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Eliminate will remove adjacent duplicates in-place in a `[]string` slice.
// The slice argument must be given as a pointer so that changes can be effected upon the slice itself.
func Eliminate(s *[]string) {
	// Start from the second element, as the comparison loop looks at previous vs current index
	index := 1

	for index < len(*s) {
		if (*s)[index-1] == (*s)[index] {
			// Delete duplicate at 'index'
			copy((*s)[index:], (*s)[index+1:])

			// Shrink slice length by one
			*s = (*s)[:len(*s)-1]
		} else {
			index++
		}
	}
}

// EliminateSpaces will squash in-place each run of adjacent Unicode spaces into a single ASCII space.
// Great read at https://blog.golang.org/strings which was very helpful.
func EliminateSpaces(b *[]byte) {
	if !utf8.Valid(*b) {
		panic(fmt.Sprintf("input '% x' does not consist entirely of valid UTF-8-encoded runes", *b))
	}

	index := 0
	prev := utf8.MaxRune + 1 // placeholder

	for index < len(*b) {
		// Take a sub-slice of the given slice
		r, size := utf8.DecodeRune((*b)[index:])
		if r == utf8.RuneError {
			panic("got RuneError")
		}

		if unicode.IsSpace(r) && r == prev {
			// Bring forward the remainder of the slice
			copy((*b)[index:], (*b)[index+size:])
			// Shrink the slice
			(*b) = (*b)[:len(*b)-size]
		} else {
			prev = r
			index += size
		}
	}
}
