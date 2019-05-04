package ex1_3_test

import (
	"fmt"
	"io"
	"testing"
)

func echo1(w io.Writer, args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func BenchmarkEcho1_2(b *testing.B)      { benchmark(b, 2, echo1) }
func BenchmarkEcho1_10(b *testing.B)     { benchmark(b, 10, echo1) }
func BenchmarkEcho1_100(b *testing.B)    { benchmark(b, 100, echo1) }
func BenchmarkEcho1_1000(b *testing.B)   { benchmark(b, 1000, echo1) }
func BenchmarkEcho1_10000(b *testing.B)  { benchmark(b, 10000, echo1) }
func BenchmarkEcho1_100000(b *testing.B) { benchmark(b, 100000, echo1) }
