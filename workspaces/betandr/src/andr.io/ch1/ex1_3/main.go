// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses `strings.Join`. (Section 1.6
// illustrates part of the `time` package, and Section 11.4 shows how to write
// benchmark tests for systematic performance evaluation.)
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	forStart := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("for loop took: %dms\n", int64(time.Since(forStart)))

	joinStart := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("join took: %dms\n", time.Since(joinStart))
}

//!-
