package main

import (
	"bytes"
	"testing"
)

//Write an in-place function that squashes each run of adjacent Unicode
//spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a
//single ASCII space.

func TestUnicode(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("世界"), []byte("界世")},
		{[]byte("世界界世"), []byte("世界界世")},
		{[]byte("a世a界界b世b"), []byte("b世b界界a世a")},
		{[]byte("בראשית ברא אלקים את השמים ואת הארץ"), []byte("ץראה תאו םימשה תא םיקלא ארב תישארב")},
	}
	for _, z := range tests {
		reverseUnicode(z.input)
		if !bytes.Equal(z.input, z.want) {
			t.Errorf("want %s; got %s\n", z.want, z.input)
		}
	}
}
