package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  string
	}{
		{
			name:  "Success- finds a letters, digits, whitespaces",
			input: []byte("a 5"),
			want:  fmt.Sprintf("%-15vcount\n%-15v1\n%-15vcount\n%-15v1\n%-15v1\n%-15vcount\n%-15v1\n%-15v1\n%-15v1\n", "type", "letters", "type", "letters", "whitespace", "type", "letters", "whitespace", "digits"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := new(bytes.Buffer)
			r := bytes.NewBuffer(tt.input)
			CharCount(r, w)
			if got := w.String(); got != tt.want {
				t.Errorf("CharCount() = %v\nwant %v\n", got, tt.want)
			}
		})
	}
}
