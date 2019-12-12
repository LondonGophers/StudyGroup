// Write a non-recursive version of `comma`, using `bytes.Buffer` instead of string concatenation.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", CommaRecurse(os.Args[i]))
		fmt.Printf("  %s\n", CommaBuffer(os.Args[i]))
	}
}

func CommaRecurse(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return CommaRecurse(s[:n-3]) + "," + s[n-3:]
}

func CommaBuffer(s string) string {
	b := bytes.Buffer{}

	for index := 0; index < len(s); index++ {
		if (len(s)-index)%3 == 0 && index > 0 {
			b.WriteRune(',')
		}

		b.WriteByte(s[index])
	}

	return b.String()
}
