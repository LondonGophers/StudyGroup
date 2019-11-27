package main_test

import (
	"fmt"

	ex "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch3/ex3_10"
)

func ExampleCommaRecurse() {
	fmt.Printf("%#v\n", ex.CommaRecurse("1234567"))
	// Output: "1,234,567"
}

func ExampleCommaBuffer() {
	fmt.Printf("%#v\n", ex.CommaBuffer("1234567"))
	// Output: "1,234,567"
}

func ExampleCommaBuffer_no_commas() {
	fmt.Printf("%#v, %#v, %#v\n", ex.CommaBuffer("1"), ex.CommaBuffer("12"), ex.CommaBuffer("123"))
	// Output: "1", "12", "123"
}

func ExampleCommaBuffer_one_comma() {
	fmt.Printf("%#v, %#v, %#v\n", ex.CommaBuffer("1234"), ex.CommaBuffer("12345"), ex.CommaBuffer("123456"))
	// Output: "1,234", "12,345", "123,456"
}
