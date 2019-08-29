package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func timeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s\t\t%s\n", name, elapsed)
}

func inefficientPrint() {
	defer timeTaken(time.Now(), "inefficient")
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%v: %s\n", i, os.Args[i])
	}
}

func efficientPrint() {
	defer timeTaken(time.Now(), "efficient")
	fmt.Println(os.Args[:])
}

func main() {

	inefficientPrint()
	efficientPrint()
}
