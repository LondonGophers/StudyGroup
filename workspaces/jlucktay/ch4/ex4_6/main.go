// Package adjacent contains the Eliminate() and EliminateSpaces() funcs.
package adjacent

import (
	"bytes"
	"unicode"
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
func EliminateSpaces(b *[]byte) {
	rb := bytes.Runes(*b)
	index := 1
	for index < len(rb) {
		if unicode.IsSpace(rb[index-1]) && unicode.IsSpace(rb[index]) {
			copy(rb[index:], rb[index+1:])
			rb = rb[:len(rb)-1]
		} else {
			index++
		}
	}

	*b = []byte(string(rb))
}
