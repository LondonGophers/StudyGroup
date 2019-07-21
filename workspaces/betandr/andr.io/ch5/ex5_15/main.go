// Write variadic functions `max` and `min`, analogous to `sum`. What should
// these functions do when called with no arguments? Write variants that require
// at least one argument.
package main

import "fmt"

// max returns the largest integer in the vals argument list
func max(vals ...int) (max int) {
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return
}

// max2 returns the largest integer in the vals argument list but requires
// at least one integer to be supplied
func max2(val int, vals ...int) (max int) {
	max = val // start wih the extra value
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return
}

// min returns the smallest integer in the vals argument list
func min(vals ...int) (min int) {
	if len(vals) > 0 {
		min = vals[0] // ensure we have at least one val from the actual args
	}

	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return
}

// min2 returns the smallest integer in the vals argument list but requires
// at least one integer to be supplied
func min2(val int, vals ...int) (min int) {
	min = val // ensure we have at least one val from the actual args

	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return
}

func main() {
	fmt.Printf("max of () is %d\n", max())
	fmt.Printf("max of (3) is %d\n", max(3))
	fmt.Printf("max of (1, 2, 3, 4) is %d\n", max(1, 2, 3, 4))

	fmt.Printf("max of (3) is %d\n", max2(3))
	fmt.Printf("max of (1, 2, 3, 4) is %d\n", max2(1, 2, 3, 4))

	fmt.Printf("min of () is %d\n", min())
	fmt.Printf("min of (3) is %d\n", min(3))
	fmt.Printf("min of (1, 2, 3, 4) is %d\n", min(1, 2, 3, 4))

	fmt.Printf("min of (3) is %d\n", min2(3))
	fmt.Printf("min of (1, 2, 3, 4) is %d\n", min2(1, 2, 3, 4))
}
