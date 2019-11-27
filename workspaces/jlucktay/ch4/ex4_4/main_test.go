package rotate_test

import (
	"fmt"

	rotate "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch4/ex4_4"
)

func ExampleLeft_one_length_five() {
	a := [...]int{1, 2, 3, 4, 5}
	rotate.Left(a[:], 1)
	fmt.Println(a)
	// Output: [2 3 4 5 1]
}

func ExampleLeft_one_length_eight() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate.Left(a[:], 1)
	fmt.Println(a)
	// Output: [1 2 3 4 5 6 7 0]
}

func ExampleLeft_three_length_seventeen() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	rotate.Left(a[:], 3)
	fmt.Println(a)
	// Output: [3 4 5 6 7 8 9 10 11 12 13 14 15 16 0 1 2]
}

func ExampleLeft_seven_length_thirty_seven() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
		28, 29, 30, 31, 32, 33, 34, 35, 36}
	rotate.Left(a[:], 7)
	fmt.Println(a)
	// Output: [7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 0 1 2 3 4 5 6]
}

func ExampleRight_one() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate.Right(a, 1)
	fmt.Println(a)
	// Output: [7 0 1 2 3 4 5 6]
}

func ExampleRight_two() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate.Right(a, 2)
	fmt.Println(a)
	// Output: [6 7 0 1 2 3 4 5]
}

func ExampleRight_three() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate.Right(a, 3)
	fmt.Println(a)
	// Output: [5 6 7 0 1 2 3 4]
}

func ExampleRight_four() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate.Right(a, 4)
	fmt.Println(a)
	// Output: [4 5 6 7 0 1 2 3]
}

func ExampleRight_five() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate.Right(a, 5)
	fmt.Println(a)
	// Output: [3 4 5 6 7 0 1 2]
}
