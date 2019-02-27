package main

import (
	"reflect"
	"strings"
	"testing"
)

var cases = []struct {
	input  string
	output map[string]int
}{
	{input: "a\nb\nc", output: map[string]int{"a": 1, "b": 1, "c": 1}},
	{input: "e\nd\nf\nd", output: map[string]int{"d": 2, "e": 1, "f": 1}},
}

func TestCountLines(t *testing.T) {
	for _, tc := range cases {
		counts := make(map[string]int)
		countLines(strings.NewReader(tc.input), counts, "test", nil)
		if reflect.DeepEqual(counts, tc.output) == false {
			t.Errorf("unexpected line count: got %d, want %d", counts["a"], 1)
		}
	}
}
