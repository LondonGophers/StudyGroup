package main

import "bytes"

func comma(s string) string {
	var buf bytes.Buffer
	
	// if its not divivisble by three the first part will be less than 3
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

//func betandr(s string) string {
//	if len(s) < 4 {
//		return s
//	}
//
//	var buf bytes.Buffer
//
//	idx := len(s) % 3
//	buf.WriteString(s[:idx])
//
//	for i := idx; i < len(s); i += 3 {
//		buf.WriteString("," + s[i:i+3]) <--- bug here if its divisble by three and greater than 5 digits
//	}
//
//	return buf.String()
//}

//--- FAIL: TestBetandr (0.00s)
//    main_test.go:47: got ,123,456, want 123,456
//    main_test.go:47: got ,123,456,789, want 123,456,789
//FAIL

func main() {}