package main

import (
	"unicode"
	"unicode/utf8"
)

//squashing spaces is just a special case of deuplicating with two
//differences. Firstly it must only consider spaces and secondly
//when shifting it must shift the number of bytes of the rune, not just one byte.
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
			shift(b, i, ln)
			end -= ln
		}
		// if its the last space in the run, change it to an ascii space
		// and shift left if necessary.
		// this will be one less as we want to keep one byte.
		if unicode.IsSpace(lr) && !unicode.IsSpace(rr) && lr != 32 {
			b[i] = ' '
			shift(b, i, ln-1)
			end -= ln - 1
		}
	}
	*s = b[:end]
}

func shift(b []byte, p int, l int) {
	for k := 0; k < l; k++ {
		for j := p - 1; j < len(b)-1; j++ {
			b[j] = b[j+1]
		}
	}
}
