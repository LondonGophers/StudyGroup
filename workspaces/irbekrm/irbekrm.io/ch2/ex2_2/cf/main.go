/*
Package cf converts its numeric arguments between
temperatures in Celsius and Fahrenheit,
length in feet and meters and
weight in pounds and kilograms
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/go-london-user-group/study-group/workspaces/irbekrm/ch2/ex2_1"
	lengthconv "github.com/go-london-user-group/study-group/workspaces/irbekrm/ch2/ex2_2/lengthconv"
	weightconv "github.com/go-london-user-group/study-group/workspaces/irbekrm/ch2/ex2_2/weightconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			parseInput(arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			parseInput(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	}
}

func parseInput(s string) {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	asTemp(n)
	asWeight(n)
	asLength(n)
}

func asTemp(number float64) {
	f := tempconv.Fahrenheit(number)
	c := tempconv.Celsius(number)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func asWeight(number float64) {
	p := weightconv.Pounds(number)
	kg := weightconv.Kilograms(number)
	fmt.Printf("%s = %s, %s = %s\n",
		p, weightconv.PToKg(p), kg, weightconv.KgToP(kg))
}

func asLength(number float64) {
	m := lengthconv.Meters(number)
	ft := lengthconv.Feet(number)
	fmt.Printf("%s = %s, %s = %s\n",
		m, lengthconv.MToFt(m), ft, lengthconv.FtToM(ft))
}
