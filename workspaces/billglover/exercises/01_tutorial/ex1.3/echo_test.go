package echo

import "testing"

var args = generate(100)

func generate(n int) []string {
	s := make([]string, n)
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
