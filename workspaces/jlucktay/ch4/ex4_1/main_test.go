package main_test

import (
	"testing"

	ex "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch4/ex4_1"
	"github.com/matryer/is"
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
	for desc, tC := range testCases {
		tC := tC // pin!
		is := is.New(t)
		t.Run(desc, func(t *testing.T) {
			is.Equal(ex.BitDiffCount(tC.left, tC.right), tC.diff) // actual vs expected
		})
	}
}
