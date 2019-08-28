/*
echo prints the name of the command and the command line arguments
passed to it.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

