package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		left  [32]byte
		right [32]byte
		want  int
	}{
		{[32]byte{0: 1}, [32]byte{0: 0}, 1},
		{[32]byte{0: 1}, [32]byte{0: 1}, 0},
		{[32]byte{0: 8}, [32]byte{0: 1}, 2},
		{[32]byte{0: 9}, [32]byte{0: 1}, 1},
	}

	for _, z := range tests {
		got := DiffCount(z.left, z.right)
		if got != z.want {
			t.Errorf("%.8b <> %.8b , wanted %d got %d", z.left, z.right, z.want, got)
		}

	}
}

//According to the gopl page 83
//approx half of the bits should be different
func TestX(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	diff := DiffCount(c1, c2)
	fmt.Printf("The differing bits of the SHA256 of x and X is %d.\n", diff)
}