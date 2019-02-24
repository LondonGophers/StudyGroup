package main

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		input []int
		rot   int
		want  []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
	}
	for _, z := range tests {
		got := rotate(z.input, z.rot)
		if !equals(got, z.want) {
			t.Errorf("want %#v, got %#v\n", z.want, got)
		}
	}
}

func TestInPlace(t *testing.T) {

	tests := []struct {
		input []int
		rot   int
		want  []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
	}
	for _, z := range tests {
		rotate2(z.input, z.rot)
		if !equals(z.input, z.want) {
			t.Errorf("want %#v, got %#v\n", z.want, z.input)
		}
	}
}

func equals(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}