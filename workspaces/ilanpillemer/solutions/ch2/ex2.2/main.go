// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	tempconv "gitlab.com/london-gophers/study-group/workspaces/ilanpillemer/solutions/ch2/ex2.1"
)

func main() {

	var w *bufio.Scanner
	if len(os.Args) > 1 {
		w = bufio.NewScanner(strings.NewReader(strings.Join(os.Args[1:], "\n")))
	} else {
		w = bufio.NewScanner(os.Stdin)
	}
	for w.Scan() {
		args := strings.Fields(w.Text())
		for _, arg := range args {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FtoC(f), c, tempconv.CtoF(c))
		}
	}
}

//!-
