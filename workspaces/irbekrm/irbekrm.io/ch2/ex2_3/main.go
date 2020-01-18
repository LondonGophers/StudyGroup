package main

import (
	"fmt"
	"testing"

	"github.com/LondonGophers/StudyGroup/workspaces/irbekrm/ch2/ex2_3/popcount"
)

func main() {
	ns1 := benchmark(popcount.Count, 98765).NsPerOp()
	ns2 := benchmark(popcount.CountLoop, 98765).NsPerOp()
	fmt.Printf("Average execution length of population count using a single expression:  %v ns\n", ns1)
	fmt.Printf("Average execution length of population count using a loop: %v ns\n", ns2)
}

func benchmark(count func(uint64) int, arg uint64) testing.BenchmarkResult {
	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			count(arg)
		}
	}
	r := testing.Benchmark(f)
	return r
}
