// Write a program that prints the SHA256 hash of it standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var shaType = flag.String("a", "sha256", "The algorithm to use: sha256/sha384/sha512")

func main() {
	flag.Parse()
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}

	if *shaType == "sha256" {
		fmt.Printf("%x\n", sha256.Sum256(input))

	} else if *shaType == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384(input))

	} else if *shaType == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512(input))

	} else {
		fmt.Printf("unknown algorithm: %s\n", *shaType)
	}
}
