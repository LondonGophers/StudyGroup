package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := [10]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	reverse(&input)

	if input != [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} {
		t.Errorf("got %#v, want [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}\n", input)
	}

}