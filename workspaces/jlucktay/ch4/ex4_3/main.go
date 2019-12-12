// Rewrite `reverse` to use an array pointer instead of a slice.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Interactive test of Reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}

		if len(ints) == 8 {
			a := [8]int{}
			copy(a[:], ints)
			ReverseArrPtr(&a)
			ints = a[:]
		} else {
			Reverse(ints)
		}

		fmt.Printf("%v\n", ints)

		if errInput := input.Err(); errInput != nil {
			log.Fatal(errInput)
		}
	}
}

// Reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReverseArrPtr reverses an array of ints in place, via pointer.
func ReverseArrPtr(a *[8]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
