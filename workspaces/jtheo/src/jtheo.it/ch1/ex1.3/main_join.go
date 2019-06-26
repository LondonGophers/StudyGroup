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
	_ = strings.Join(os.Args[1:], " ")
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
}

//!-
