// Rewrite `reverse` to use an array pointer instead of a slice

package main

import "fmt"

func rotate(pos int, a *[10]int) {

	for i, j := 0, pos-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for i, j := pos, len(a)-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(3, &a)
	fmt.Println(a)
}
