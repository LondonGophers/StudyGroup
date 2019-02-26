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
