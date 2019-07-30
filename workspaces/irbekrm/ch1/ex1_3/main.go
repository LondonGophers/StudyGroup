package main

import (
	"fmt"
	"strings"
	"testing"
)

var testArgs = []string{"alpha", "beta", "gamma"}

func main() {
	ns1 := benchmark(EfficientEcho, testArgs).NsPerOp()
	ns2 := benchmark(InefficientEcho, testArgs).NsPerOp()
	fmt.Printf("Results: \nEfficientEcho average runtime: %v ns\nInefficientEcho average runtime: %v ns\n", ns1, ns2)
}

func benchmark(echo func([]string), args []string) testing.BenchmarkResult {
	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			echo(args)
		}
	}
	r := testing.Benchmark(f)
	return r
}

func EfficientEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func InefficientEcho(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}
