package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello, 池")
	bb := ReverseChars(b)
	fmt.Println(string(bb))
}

func ReverseChars(bytes []byte) []byte {
	lidx := 0              // left index
	ridx := len(bytes) - 1 // right index
	for {
		if lidx >= ridx {
			return bytes
		}
		// Rune on the left and its length
		lr, llen := utf8.DecodeRune(bytes[lidx:])
		// Rune on the right and its length
		rr, rlen := utf8.DecodeLastRune(bytes[:ridx+1])
		// Left rune has more bytes than right rune
		if llen > rlen {
			// Write right rune to the space on left
			for i := 0; i < rlen; i++ {
				bytes[lidx+i] = bytes[ridx-rlen+i+1]
			}
			// Shift everything in between the two runes llen - rlen spaces to the left
			for idx := lidx + rlen; idx <= ridx-llen; idx++ {
				bytes[idx] = bytes[idx+llen-rlen]
			}
			// Write left rune to the space on right
			b := []byte(string(lr))
			for i, bb := range b {
				bytes[ridx-llen+i+1] = bb
			}
			// Right rune has more bytes than left rune
		} else if rlen > llen {
			// Write left rune to the space on the right
			for i := 0; i < llen; i++ {
				bytes[ridx-llen+i+1] = bytes[lidx+i]
			}
			// Shift everything between the two bytes llen - rlen spaces to the right
			for i := ridx - llen; i >= lidx+rlen; i-- {
				bytes[i] = bytes[i-(rlen-llen)]
			}
			// Write right rune to the space on the left
			b := []byte(string(rr))
			for i, bb := range b {
				bytes[lidx+i] = bb
			}
			// Both runes have the same number of bytes
		} else {
			for i := 0; i < llen; i++ {
				bytes[lidx+i], bytes[ridx-rlen+i+1] = bytes[ridx-rlen+i+1], bytes[lidx+i]
			}
		}
		lidx += rlen
		ridx -= llen
	}
}
