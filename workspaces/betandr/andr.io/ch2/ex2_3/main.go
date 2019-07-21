// Rewrite `PopCount` to use a loop instead of a single expression. Compare the
// performance of the two versions. (Section 11.4 shows how to compare the
// performance of different implementations systematically.)
package main

import (
	"fmt"

	"andr.io/ch2/ex2_3/popcount"
)

func main() {

	i := uint64(240)
	populationCount := popcount.PopCount(i)
	fmt.Printf("population count of %b is %d\n", i, populationCount)
}
