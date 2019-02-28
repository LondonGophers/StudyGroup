package main

import (
	"unicode"
	"unicode/utf8"
)

//squashing spaces is just a special case of deuplicating with two
//differences.
//firstly it must only consider spaces
func squash(s *[]byte) {
	b := *s
	end := len(b)
	for i := len(b) - 1; i > 0; i-- {
		lr, ln := utf8.DecodeLastRune(b[:i+1])
		rr, _ := utf8.DecodeLastRune(b[:i+1-ln])

		if !unicode.IsSpace(lr) && !unicode.IsSpace(rr) {
			// they are not equal.. dont do anything..
		}
		if unicode.IsSpace(lr) && unicode.IsSpace(rr) {
			// they are equal copy all from the right left the
			// length of the rightmost rune
			for k := 0; k < ln; k++ {
				for j := i - 1; j < len(b)-1; j++ {
					b[j] = b[j+1]
				}
				// decrement end pointer
				end--
			}
		}
		// if its the last space in the run, change it to an ascii space
		// and shift left if necessary.
		// this will be one less as we want to keep one byte.
		if unicode.IsSpace(lr) && !unicode.IsSpace(rr) && lr != 32 {
			b[i] = ' '
			for k := 0; k < ln-1; k++ {
				for j := i - 1; j < len(b)-1; j++ {
					b[j] = b[j+1]
				}
				end--
			}
		}
	}
	*s = b[:end]
}
