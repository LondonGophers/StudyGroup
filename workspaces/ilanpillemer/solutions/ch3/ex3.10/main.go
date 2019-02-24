package main

import "bytes"

func comma(s string) string {
	var buf bytes.Buffer

	// if its not divisble by three the first part will be less than 3
	start := len(s) % 3
	buf.WriteString(s[:start])

	for i := start; i < len(s); i += 3 {
		if i != 0 { // if it is divisible by three don't start with a comma.
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
