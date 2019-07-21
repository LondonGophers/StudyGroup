// Write a function `expand(s string, f func(string) string) string` that
// replaces each substring `$foo` within `s` by the text returned by `f("foo")`.
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Message: $foo"
	fmt.Println(expand(s, replace))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("world"), -1)
}

func replace(s string) string {
	return fmt.Sprintf("hello, %s!", s)
}
