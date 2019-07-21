// Modify the `echo` program to also allow print `os.Args[0]`, the name of the
// command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(os.Args[0] + " " + s)
}

//!-
