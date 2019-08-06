package main

import "fmt"

func ExampleKB() {
	fmt.Printf("%G %d", kB, uint64(kB))
	//Output: 1000 1000
}

func ExampleMB() {
	fmt.Printf("%G %d", MB, uint64(MB))
	//Output: 1E+06 1000000
}

func ExampleGB() {
	fmt.Printf("%G %d", GB, uint64(GB))
	//Output: 1E+09 1000000000
}

func ExampleTB() {
	fmt.Printf("%G %d", TB, uint64(TB))
	//Output: 1E+12 1000000000000
}

func ExamplePB() {
	fmt.Printf("%G %d", PB, uint64(PB))
	//Output: 1E+15 1000000000000000
}

func ExampleEB() {
	fmt.Printf("%G %d", EB, uint64(EB))
	//Output: 1E+18 1000000000000000000
}

func ExampleZB() {
	fmt.Printf("%G", ZB)
	//Output: 1E+21
}

func ExampleYB() {
	fmt.Printf("%G", YB)
	//Output: 1E+24
}
