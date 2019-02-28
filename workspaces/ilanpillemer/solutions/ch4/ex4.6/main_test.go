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
		{[]byte("abc"), []byte("abc")},
		{[]byte(" abc"), []byte(" abc")},
		{[]byte(" abc               "), []byte(" abc ")},
		{[]byte{'a', ' ', 0xc2, 0x85, 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', 'c'}},
		{[]byte{'a', '\t', 0xc2, 0x85, 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', 'c'}},
		{[]byte{'a', 0xc2, 0x85, 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', 'c'}},
		{[]byte{'a', '\t', 0xc2, 0x85, 'b', 0xc2, 0xa0, 'b', 'c'}, []byte{'a', ' ', 'b', ' ', 'b', 'c'}},
		{[]byte{'a', '\t', 0xc2, 0x85, 'b', 0xc2, 0xa0, 'b', 'c', 'c'}, []byte{'a', ' ', 'b', ' ', 'b', 'c', 'c'}},
		{[]byte("a b c d e"), []byte("a b c d e")},
		{[]byte("a  b    c     d      e"), []byte("a b c d e")},
		{[]byte("Finally,  as    the sky      began   to    grow           light"), []byte("Finally, as the sky began to grow light")},
		{[]byte("Finally,  as    the sky      began   to"), []byte("Finally, as the sky began to")},
	}

	for _, z := range tests {
		squash(&z.input)
		if !bytes.Equal(z.input, z.want) {
			t.Errorf("squash: got [%#v], want [%#v]\n", string((z.input)), string(z.want))
		}
	}
}


