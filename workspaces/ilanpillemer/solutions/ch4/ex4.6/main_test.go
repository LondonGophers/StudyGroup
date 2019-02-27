package main

import (
	"bytes"
	"testing"
)

//Write an in-place function that squashes each run of adjacent Unicode
//spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a
//single ASCII space.

func Test(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte{'a', '\t', 0xc2, 0x85, 0xc2, 0xa0, 'b'}, []byte{'a', ' ', 'b'}},
		{[]byte{'a', '\t', 0xc2, 0x85, 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', 'c'}},
		{[]byte{'a', 0xc2, 0x85, 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', 'c'}},
		{[]byte{'a', '\t', 0xc2, 0x85, 'b', 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', ' ', 'b', 'c'}},
		{[]byte{'a', '\t', 0xc2, 0x85, 'b', 0xc2, 0xa0, 'b', 'c', 'c'}, []byte{'a', ' ', 'b', ' ', 'b', 'c', 'c'}},
	}

	for _, z := range tests {
		squash(&z.input)
		if !bytes.Equal(z.input, z.want) {
			t.Errorf("got [%#v], want [%#v]\n", string((z.input)), string(z.want))
		}

	}

}

//func TestStuff(t *testing.T) {
//	fmt.Printf("U+0085 -> %b and is space? :%t \n", rune(0x85), unicode.IsSpace(rune(0x85)))
//	fmt.Printf("U+00A0 -> %b and is space? :%t \n", rune(0xA0), unicode.IsSpace(rune(0xA0)))
//	fmt.Printf("0xc2 0x85 -> %b %b \n", byte(0xc2), byte(0x85))
//	r := []rune{'a', '\t', 0x85, 0xA0, 'b'}
//
//	fmt.Printf("%#v\n", []byte(string(r)))
//}
