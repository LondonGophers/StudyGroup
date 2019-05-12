package ex1_3_test

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
)

func generateSlice(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, randStringBytesMaskImprSrcSB(rand.Intn(20)))
	}
	return s
}

func benchmark(b *testing.B, size int, benchMe func(io.Writer, []string)) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := generateSlice(size)
		buf := &bytes.Buffer{}
		b.StartTimer()
		benchMe(buf, a)
	}
}
