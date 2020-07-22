package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, positions int) []int {
	res := make([]int, len(s))
	for k, v := range s {
		i := (k + positions) % (len(s))
		res[i] = v
	}
	return res
}
