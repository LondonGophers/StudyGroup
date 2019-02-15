// Add support for Kelvin temperatures to `tempFlag`.
//
// Explain why the help message contains °C when the default value of
// `20.0` does not.
// It uses the `String()` method declared in `tempconv`
// func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
package main

import (
	"flag"
	"fmt"

	"andr.io/ch7/ex7_6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
