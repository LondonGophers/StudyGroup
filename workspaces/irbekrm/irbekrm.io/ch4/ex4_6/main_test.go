package main

import (
	"reflect"
	"testing"
)

func TestSquashAdjacentSpaces(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "Squashes \t, \n, \v and space",
			input: []byte("one\t\n\v two    \tthree    "),
			want:  []byte("one two three "),
		},
		{
			name:  "Squashes \f, \r and space",
			input: []byte("one\f\f\f  two\r\f    \fthree    \r"),
			want:  []byte("one two three "),
		},
		{
			name:  "Squashes U+0085 (NEL), U+00A0 (NBSP) and space",
			input: []byte("one\u0085\u0085\u0085   two   three\u00A0 \u00A0\u0085 "),
			want:  []byte("one two three "),
		},
		{
			name:  "Works when nothing needs to be squashed",
			input: []byte("one two three "),
			want:  []byte("one two three "),
		},
		{
			name:  "Does not break encoding with multi-byte code points",
			input: []byte("施  \t地\n   世界"),
			want:  []byte("施 地 世界"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquashAdjacentSpaces(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquashAdjacentSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
