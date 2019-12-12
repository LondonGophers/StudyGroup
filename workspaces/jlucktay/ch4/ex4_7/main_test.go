package reverse_test

import (
	"fmt"

	ex "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch4/ex4_7"
)

func ExampleReverse() {
	a := []int{0, 1, 2, 3, 4, 5}

	ex.Reverse(&a)

	fmt.Println(a)

	// Output: [5 4 3 2 1 0]
}

func ExampleReverse_odd_length() {
	a := []int{0, 1, 2, 3, 4}

	ex.Reverse(&a)

	fmt.Println(a)

	// Output: [4 3 2 1 0]
}

func ExampleReverse_length_one() {
	a := []int{0}

	ex.Reverse(&a)

	fmt.Println(a)

	// Output: [0]
}

func ExampleReverse_length_two() {
	a := []int{0, 42}

	ex.Reverse(&a)

	fmt.Println(a)

	// Output: [42 0]
}

func ExampleString() {
	b := []byte("Hello, 世界! 👋")

	ex.String(&b)

	fmt.Printf("%s\n", b)

	// Output: 👋 !界世 ,olleH
}
