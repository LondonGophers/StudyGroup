package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[6]int) {
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
