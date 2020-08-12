package main

import (
	"reflect"
	"testing"
)

func TestWordFreq(t *testing.T) {
	tests := []struct {
		name string
		path string
		want map[string]int
	}{
		{
			name: "Success- calculates word frequencies in example file",
			path: "example.txt",
			want: map[string]int{"one": 3, "two": 1, "three": 1, "four": 1, "five": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordFreq(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}
