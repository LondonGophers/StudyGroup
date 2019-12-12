package ex1_3_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func echo3(w io.Writer, args []string) {
	fmt.Fprint(w, strings.Join(args[0:], " "))
}

func BenchmarkEcho3_2(b *testing.B)      { benchmark(b, 2, echo3) }
func BenchmarkEcho3_10(b *testing.B)     { benchmark(b, 10, echo3) }
func BenchmarkEcho3_100(b *testing.B)    { benchmark(b, 100, echo3) }
func BenchmarkEcho3_1000(b *testing.B)   { benchmark(b, 1000, echo3) }
func BenchmarkEcho3_10000(b *testing.B)  { benchmark(b, 10000, echo3) }
func BenchmarkEcho3_100000(b *testing.B) { benchmark(b, 100000, echo3) }
