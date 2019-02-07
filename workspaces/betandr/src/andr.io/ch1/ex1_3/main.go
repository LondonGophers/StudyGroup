// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
  "time"
  "strings"
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
