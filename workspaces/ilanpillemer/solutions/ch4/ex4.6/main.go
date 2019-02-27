package main

import (
	"unicode"
	"unicode/utf8"
)

func squash(s *[]byte) {
	b := *s
	todo := len(b) + 1 // when we are done there wont be anything todo
	start, end := 0, 0 // unicode run that needs to be replaced with one space
	inSpace := false   // state
	elided := 0        // how much to chop off the end when we are done
	for todo > 0 {
		r, l := utf8.DecodeRune(b[end:])
		if unicode.IsSpace(r) {
			if inSpace {
				end += l // move end pointer
			} else {
				start = end // set begin pointer
				end += l
			}
			inSpace = true
			todo = todo - l
			elided += l // more to chop
			continue
		}
		if inSpace {
			b[start] = ' '
			elided--                   // account for the single space that replaces a run when chopping
			copy(b[start+1:], b[end:]) // replace run with space
			todo = todo - l
			end = start + l // reset pointers
			inSpace = false
			continue
		}
		todo = todo - l
		end = end + l // shift pointer
	}
	b = b[:len(b)-elided] //cut off end
	*s = b
}
