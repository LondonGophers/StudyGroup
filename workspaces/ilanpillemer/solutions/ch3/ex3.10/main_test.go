package main

import "testing"

func TestComma(t *testing.T) {

	tests := []struct {
		input string
		want  string
	}{
		{"123", "123"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
	}

	for _, tst := range tests {

		got := comma(tst.input)
		if got != tst.want {
			t.Errorf("got %s, want %s", got, tst.want)
		}
	}

}

//func TestBetandr(t *testing.T) {
//
//	tests := []struct {
//		input string
//		want  string
//	}{
//		{"123", "123"},
//		{"12345", "12,345"},
//		{"123456", "123,456"},
//		{"1234567", "1,234,567"},
//		{"12345678", "12,345,678"},
//		{"123456789", "123,456,789"},
//	}
//
//	for _, tst := range tests {
//
//		got := betandr(tst.input)
//		if got != tst.want {
//			t.Errorf("got %s, want %s", got, tst.want)
//		}
//	}
//}

