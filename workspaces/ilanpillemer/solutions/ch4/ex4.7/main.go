package main

import (
	"unicode/utf8"
)

//Modify reverse to reverse the characters of a []byte slice that
//represents a UTF-8-encoded string, in place. Can you do it without
//allocating new memory?
func reverseUnicode(t []byte) {

	s := t
	//convert to runes see bytes.Runes..
	rs := make([]rune, utf8.RuneCount(s))
	i := 0
	for len(s) > 0 {
		r, l := utf8.DecodeRune(s)
		rs[i] = r
		i++
		s = s[l:]
	}

	//reverse the runes
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	// put the runes back in the bytes slice pointed at by the pointer
	// header
	i = 0
	for len(rs) > 0 {
		l := utf8.RuneLen(rs[0])
		utf8.EncodeRune(t[i:i+l], rs[0])
		i = i + l
		rs = rs[1:]
	}
}
