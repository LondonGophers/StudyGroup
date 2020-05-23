package main_test

import (
	"testing"

	"github.com/matryer/is"

	ex "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch4/ex4_1"
)

func TestBitDiffCount(t *testing.T) {
	testCases := map[string]struct {
		left, right [32]byte
		diff        int
	}{
		"Start small": {
			left:  [32]byte{0: 0},
			right: [32]byte{0: 1},
			diff:  1,
		},
		"Double the diff": {
			left:  [32]byte{0: 0},
			right: [32]byte{0: 3},
			diff:  2,
		},
		"And double again": {
			left:  [32]byte{0: 0},
			right: [32]byte{0: 15},
			diff:  4,
		},
		"And again": {
			left:  [32]byte{0: 0},
			right: [32]byte{0: 255},
			diff:  8,
		},
		"Flip the script": {
			left:  [32]byte{0: 255},
			right: [32]byte{0: 0},
			diff:  8,
		},
		"Take it down a notch": {
			left:  [32]byte{0: 127},
			right: [32]byte{0: 0},
			diff:  7,
		},
		"Mix it up": {
			left:  [32]byte{0: 63},
			right: [32]byte{0: 64},
			diff:  7,
		},
		"Keep going": {
			left:  [32]byte{0: 222},
			right: [32]byte{0: 111},
			diff:  4,
		},
	}

	is := is.New(t)

	for desc, tC := range testCases {
		// pin! ref: https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		desc, tC := desc, tC

		t.Run(desc, func(t *testing.T) {
			t.Parallel() // Don't use .Parallel() without pinning.

			is.Equal(ex.BitDiffCount(tC.left, tC.right), tC.diff) // actual vs expected
		})
	}
}
