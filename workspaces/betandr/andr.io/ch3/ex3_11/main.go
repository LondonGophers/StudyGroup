// Enhance `comma` so it deals correctly with floating-point numbers and an
// optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	point := strings.LastIndex(s, ".")
	digits := s[point:]
	s = s[:point]

	if strings.HasPrefix(s, "-") {
		buf.WriteString("-")
		s = s[1:]
	}

	if len(s) < 4 {
		return s
	}

	idx := len(s) % 3
	buf.WriteString(s[:idx])

	for i := idx; i < len(s); i += 3 {
		buf.WriteString("," + s[i:i+3])
	}

	return buf.String() + digits
}
