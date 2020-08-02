package main

import (
	"reflect"
	"testing"
)

func TestReverseChars(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "Reverses a byte slice with UTF-8 encoded single-byte Unicode code points",
			input: []byte("Hello, world"),
			want:  []byte("dlrow ,olleH"),
		},
		{
			name:  "Reverses a byte slice with UTF-8 encoded multi-byte Unicode code points",
			input: []byte("wind: 風, air: 氣, time: 時, echo: 響"),
			want:  []byte("響 :ohce ,時 :emit ,氣 :ria ,風 :dniw"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseChars(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
