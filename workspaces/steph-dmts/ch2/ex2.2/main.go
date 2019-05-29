package main

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/steph-dmts/gopl/ch2/ex2.2/conv"
	//"fmt"
	"os"
)

func main() {

	if len(os.Args) != 1 {
		for _, args := range os.Args[1:] {
			doConv(args)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			args := scanner.Text()
			doConv(args)
		}
	}
}

func doConv(args string) {
	x, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
	}
	k := conv.Kilo(x)
	p := conv.KToP(k)
	fmt.Printf("%g kilos is %g pounds\n", k, p)
}
