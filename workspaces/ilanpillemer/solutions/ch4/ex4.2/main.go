package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

var hType = flag.String("type", "256", "hash type of 256 512 or 384")

func main() {
	flag.Parse()
	var h hash.Hash
	switch *hType {
	case "256":
		h = sha256.New()
	case "512":
		h = sha512.New()
	case "384":
		h = sha512.New384()
	default:
		flag.Usage()
		os.Exit(1)
	}
	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()
	fmt.Printf("SHA%s: %x\n", *hType, h.Sum(nil))
}