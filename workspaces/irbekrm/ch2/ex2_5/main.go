package main

import (
	"fmt"
	"testing"

	"github.com/LondonGophers/StudyGroup/workspaces/irbekrm/ch2/ex2_5/popcount"
)

func main() {
	ns1 := benchmark(popcount.Count, 987743843834874).NsPerOp()
	ns2 := benchmark(popcount.CountWithClearing, 987743843834874).NsPerOp()
	fmt.Printf("Average execution length of population count using pre-calculated lookup table: %v ns\n", ns1)
	fmt.Printf("Average execution length of population count obtained by repeatedly clearing last non-zero bit: %v ns\n", ns2)
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
