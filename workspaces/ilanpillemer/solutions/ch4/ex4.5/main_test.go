package main

import (
	"testing"
)

//Write an in-place function to eliminate adjacent duplicates in a []string slice.
func Test(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{"a", "a", "b", "b", "b", "c", "c", "c", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{}, []string{}},
		{[]string{"Finally", "as", "as", "as", "the", "the", "sky", "began", "to", "to", "grow", "light"}, []string{"Finally", "as",  "the",  "sky", "began",  "to", "grow", "light"}},
	}

	for _, z := range tests {
		dedup(&z.input)
		if !equals(z.input, z.want) {
			t.Errorf("want [%v] got [%v]\n", z.want, z.input)
		}

	}
}

func equals(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}