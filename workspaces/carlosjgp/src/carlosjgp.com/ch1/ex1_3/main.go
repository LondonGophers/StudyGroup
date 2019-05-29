// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(time.Now().Sub(start))

	var s string
	start = time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += fmt.Sprintf("%s ", os.Args[i])
	}
	fmt.Println(s)
	fmt.Println(time.Now().Sub(start))
}

//!-
