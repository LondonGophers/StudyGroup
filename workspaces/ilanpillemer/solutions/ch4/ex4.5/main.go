//Write an in-place function to eliminate adjacent duplicates in a []string slice.

package main

//When a slice is passed to function under the hood a special value is
//passed that contains a pointer and length. Although changes via the
//pointer can mutate the slice, the length is not mutated.. The only way
//to change the length is to go beyond the slice header and actually get
//a real pointer not a value pretending to be a pointer see
//https://blog.golang.org/slices

func dedup(inputPtr *[]string) {
	input := *inputPtr
	end := len(input)
	for i := len(input) - 1; i > 0; i-- {
		if input[i] != input[i-1] {
			// they are not equal.. dont do anything..
		}
		if input[i] == input[i-1] {
			// they are equal copy all from the right one left
			for j := i - 1; j < len(input)-1; j++ {
				input[j] = input[j+1]
			}
			// decrement end pointer
			end--
		}
	}
	*inputPtr = input[:end]
}