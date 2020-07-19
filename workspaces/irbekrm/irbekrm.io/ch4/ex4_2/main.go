package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var format = flag.String("format", "sha256", "SHA to hash to convert input to. Valid values are sha256, sha512 and sha384")

func main() {
	var sha interface{}
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		b := []byte(input)
		switch *format {
		case "sha256":
			sha = sha256.Sum256(b)
		case "sha512":
			sha = sha512.Sum512(b)
		case "sha384":
			sha = sha512.Sum384(b)
		default:
			log.Fatalf("Unsupported format: %s\n", *format)
		}
		fmt.Printf("The %s hash of %s:\n %x\n", *format, input, sha)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
