package echo

import (
	"fmt"
	"strings"
)

// Concat takes a slice of strings and returns a string containing the space
// separated values. It uses the + operator to join the strings.
func concat(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

// Join takes a slice of strings and returns a string containing the space
// separated values. It uses strings.Join to join the strings.
func join(args []string) string {
	s := strings.Join(args, " ")
	return s
}

// Format takes a slice of strings and returns a string containing the space
// separated values. It uses the fmt.Sprintln function to join the strings.
func format(args []string) string {
	return fmt.Sprintln(args)
}

// Custom takes a slice of strings and returns a string containing the space
// separated values. It uses custom code to join the strings.
func custom(args []string, s string) string {
	var l int

	for i := range args {
		l += len(args[i])
	}
	l += len(s) * (len(args) - 1)

	b := strings.Builder{}
	b.Grow(l)

	for i := range args {
		_, err := b.WriteString(args[i])
		if err != nil {
			panic("unable to write string to buffer")
		}

		if i == len(args)-1 {
			break
		}

		_, err = b.WriteString(s)
		if err != nil {
			panic("unable to write string to buffer")
		}
	}

	return b.String()
}
