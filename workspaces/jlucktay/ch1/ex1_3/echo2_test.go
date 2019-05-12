package ex1_3_test

import (
	"fmt"
	"io"
	"testing"
)

func echo2(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func BenchmarkEcho2_2(b *testing.B)      { benchmark(b, 2, echo2) }
func BenchmarkEcho2_10(b *testing.B)     { benchmark(b, 10, echo2) }
func BenchmarkEcho2_100(b *testing.B)    { benchmark(b, 100, echo2) }
func BenchmarkEcho2_1000(b *testing.B)   { benchmark(b, 1000, echo2) }
func BenchmarkEcho2_10000(b *testing.B)  { benchmark(b, 10000, echo2) }
func BenchmarkEcho2_100000(b *testing.B) { benchmark(b, 100000, echo2) }
