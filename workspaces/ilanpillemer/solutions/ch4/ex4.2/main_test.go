package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Printf("%x with a new line following\n", sha256.Sum256([]byte("x\n")))
	fmt.Printf("%x without a newline following\n", sha256.Sum256([]byte("x")))
}