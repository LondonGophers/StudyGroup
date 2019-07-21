// Use `panic` and `recover` to write a function that contains no return
// statement yet returns a non-zero value.
package main

import "fmt"

func dontPanic() (foo int) {
	defer func() {
		if p := recover(); p != nil {
			foo = 42
		}
	}()

	panic("Don't panic")
}

func main() {
	fmt.Printf("The answer is %d.\n", dontPanic())
}
