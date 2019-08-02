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

	b := bytes.Buffer{}
	signStart := 0
	unsignedLen := len(s)

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		signStart = 1
		unsignedLen--
		b.WriteByte(s[0])
	}

	if point := strings.IndexRune(s, '.'); point == -1 {
		for index := signStart; index < unsignedLen; index++ {
			if (unsignedLen-index)%3 == 0 && index > signStart {
				b.WriteRune(',')
			}

			b.WriteByte(s[index])
		}
	} else {
		point -= signStart

		for index := signStart; index < point; index++ {
			if (point-index)%3 == 0 && index > signStart {
				b.WriteRune(',')
			}

			b.WriteByte(s[index])
		}

		for index := point; index < len(s); index++ {
			b.WriteByte(s[index])
		}
	}

	return b.String()
}
