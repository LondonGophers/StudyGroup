package main_test

import (
	"fmt"

	ex "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch3/ex3_11"
)

func ExampleComma() {
	fmt.Printf("%#v\n", ex.Comma("1234567"))
	// Output: "1,234,567"
}

func ExampleComma_no_commas() {
	fmt.Printf("%#v, %#v, %#v\n", ex.Comma("1"), ex.Comma("12"), ex.Comma("123"))
	// Output: "1", "12", "123"
}

func ExampleComma_one_comma() {
	fmt.Printf("%#v, %#v, %#v\n", ex.Comma("1234"), ex.Comma("12345"), ex.Comma("123456"))
	// Output: "1,234", "12,345", "123,456"
}

func ExampleComma_float() {
	fmt.Printf("%#v\n", ex.Comma("12345.6789"))
	// Output: "12,345.6789"
}

func ExampleComma_positive_float() {
	fmt.Printf("%#v\n", ex.Comma("+12345.6789"))
	// Output: "+12,345.6789"
}

func ExampleComma_negative_float() {
	fmt.Printf("%#v\n", ex.Comma("-12345.6789"))
	// Output: "-12,345.6789"
}
