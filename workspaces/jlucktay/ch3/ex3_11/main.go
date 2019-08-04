// Enhance `comma` so it deals correctly with floating-point numbers and an optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", Comma(os.Args[i]))
	}
}

func Comma(s string) string {
	if _, errParse := strconv.ParseFloat(s, 64); errParse != nil {
		panic(errParse)
	}

	buf := bytes.Buffer{}

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	unsignedLen := len(s)
	point := strings.IndexRune(s, '.')

	if point > 0 {
		unsignedLen = point
	}

	for index := 0; index < unsignedLen; index++ {
		if index > 0 && (unsignedLen-index)%3 == 0 {
			buf.WriteRune(',')
		}

		buf.WriteByte(s[index])
	}

	for 0 <= point && point < len(s) {
		buf.WriteByte(s[point])
		point++
	}

	return buf.String()
}
