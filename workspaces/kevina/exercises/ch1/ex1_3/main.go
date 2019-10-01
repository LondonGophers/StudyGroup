package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
}

// echo1 the first version of the echo function
func echo1(args []string) {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%.5fs elapsed in echo1\n", time.Since(start).Seconds())
}

func echo2(args []string) {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%.5fs elapsed in echo2\n", time.Since(start).Seconds())
}

func echo3(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(args[1:], " "))
	fmt.Printf("%.5fs elapsed in echo3\n", time.Since(start).Seconds())
}
