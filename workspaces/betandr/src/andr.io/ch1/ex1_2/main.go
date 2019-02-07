// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := ""
	for idx, arg := range os.Args[1:] {
		s += strconv.Itoa(idx) + " " + arg + "\n"
	}
	fmt.Println(s)
}

//!-
