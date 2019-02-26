// Write a program that prints the SHA256 hash of it standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaType := flag.String("a", "sha256", "The algorithm to use: sha256/sha384/sha512")
	flag.Parse()

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
	}

	if *shaType == "sha256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))

	} else if *shaType == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(input)))

	} else if *shaType == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(input)))

	} else {
		fmt.Printf("unknown algorithm: %s\n", *shaType)
	}
}
