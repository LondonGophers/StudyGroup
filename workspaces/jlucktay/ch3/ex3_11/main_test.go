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

func ExampleComma_no_comma_float() {
	fmt.Printf("%#v, %#v, %#v\n", ex.Comma("1.23456789"), ex.Comma("12.3456789"), ex.Comma("123.456789"))
	// Output: "1.23456789", "12.3456789", "123.456789"
}

func ExampleComma_one_comma_float() {
	fmt.Printf("%#v, %#v, %#v\n", ex.Comma("1234.56789"), ex.Comma("12345.6789"), ex.Comma("123456.789"))
	// Output: "1,234.56789", "12,345.6789", "123,456.789"
}

func ExampleComma_two_commas_float() {
	fmt.Printf("%#v, %#v, %#v\n", ex.Comma("1234567.89"), ex.Comma("12345678.9"), ex.Comma("123456789.012"))
	// Output: "1,234,567.89", "12,345,678.9", "123,456,789.012"
}

func ExampleComma_positive_float() {
	fmt.Printf("%#v\n", ex.Comma("+12345.678901"))
	// Output: "+12,345.678901"
}

func ExampleComma_negative_float() {
	fmt.Printf("%#v\n", ex.Comma("-12345.678901"))
	// Output: "-12,345.678901"
}

func ExampleComma_long_signed_float() {
	fmt.Printf("%#v\n", ex.Comma("-1234567890.987654321"))
	// Output: "-1,234,567,890.987654321"
}
