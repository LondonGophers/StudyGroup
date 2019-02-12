// Write a variadic version of `strings.Join`.
package main

import (
	"fmt"
	"strings"
)

// join appends the supplied strings in `strs`. It uses `strings.Join`
// which is an efficient way to append strings if you already know them but
// we still have a variadic interface! ;)
func join(strs ...string) (result string) {
	return strings.Join(strs, "")
}

func main() {
	fmt.Println(join("Hello", ",", " ", "World", "!"))
}
