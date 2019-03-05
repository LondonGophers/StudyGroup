package echo

import (
	"strconv"
	"testing"
)

var args = generate(100)

func generate(n int) []string {
	s := make([]string, n)
	for i := range s {
		s[i] = strconv.Itoa(i)
	}
	return s
}

func BenchmarkFormat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		format(args)
	}
}

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		join(args)
	}
}

func TestCustom(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f"}
	if got, want := custom(a, " "), join(a); got != want {
		t.Errorf("got: %s, want %s", got, want)
	}
}

func BenchmarkCustom(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = custom(args, " ")
	}
	_ = len(s)
}
