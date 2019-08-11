package main_test

import (
	"fmt"

	ex "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch4/ex4_3"
)

func ExampleReverse() {
	a := []int{0, 1, 2, 3, 4, 5}

	ex.Reverse(a)

	fmt.Println(a)

	// Output: [5 4 3 2 1 0]
}

func ExampleReverse_rotate_twice() {
	s := []int{0, 1, 2, 3, 4, 5}

	// Rotate s left by two positions.
	ex.Reverse(s[:2])
	ex.Reverse(s[2:])
	ex.Reverse(s)

	fmt.Println(s)

	// Output: [2 3 4 5 0 1]
}

func ExampleReverseArrPtr() {
	a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}

	ex.ReverseArrPtr(&a)

	fmt.Println(a)

	// Output: [7 6 5 4 3 2 1 0]
}
