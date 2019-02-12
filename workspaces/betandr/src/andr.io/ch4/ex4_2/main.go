// Write a program that prints the SHA256 hash of it standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {

	shaType := flag.String("a", "sha256", "The algorithm to use: sha256/sha384/sha512")
	s := flag.String("s", "", "The string to hash.")
	flag.Parse()

	if *shaType == "sha256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(*s)))

	} else if *shaType == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(*s)))

	} else if *shaType == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(*s)))

	} else {
		fmt.Println("unknown algorithm %s", *shaType)
	}
}
