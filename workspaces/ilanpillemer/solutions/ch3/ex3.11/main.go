package main

import "bytes"

func comma(s string) string {
	b := []byte(s)
	parts := bytes.Split(b, []byte("."))
	whole := bytes.ToUpper(parts[0])
	//fraction := parts[1]
	wholeExp := ""
	wholeSign := ""

	if bytes.HasPrefix(whole, []byte("-")) {
		wholeSign = "-"
		whole = bytes.TrimPrefix(whole, []byte("-"))
	}

	if bytes.Contains(whole, []byte("E")) {
		index := bytes.Index(whole, []byte("E"))
		whole, wholeExp = whole[:index], string(whole[index:])

	}
	var buf bytes.Buffer

	// if its not divivisble by three the first part will be less than 3
	start := len(whole) % 3
	buf.Write(whole[:start])

	for i := start; i < len(whole); i += 3 {
		if i != 0 { // if it is divisible by three don't start with a comma.
			buf.WriteString(",")
		}
		buf.Write(whole[i : i+3])
	}

	parts[0] = buf.Bytes()
	result := bytes.Join(parts, []byte("."))
	return wholeSign + string(bytes.ToUpper(result)) + wholeExp
}

func main() {}