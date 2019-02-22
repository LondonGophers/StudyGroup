package main

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"a", "a", true},
		{"abcdef", "fedcba", true},
		{"abc", "abd", false},
	}

	for _, tst := range tests {
		got := anagram(tst.s1, tst.s2)
		if got != tst.want {
			t.Errorf("%s <> %s .. want %t, got %t\n", tst.s1, tst.s2, tst.want, got)
		}
	}

}