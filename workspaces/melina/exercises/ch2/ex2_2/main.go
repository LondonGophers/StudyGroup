//Cf converts its numeric argument to Celcius and Fahrenheit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"study-group/workspaces/melina/exercises/ch2/ex2_2/lengthconv"
	"study-group/workspaces/melina/exercises/ch2/ex2_2/weightconv"
	"gopl.io/ch2/tempconv"
)

func main() {

	var totalArgs []string

	//Read numbers from command line or from standard input if there are no agruments
	if len(os.Args) > 1 {
		totalArgs = os.Args[1:]
	} else {
		fmt.Print("Enter Numbers: ")
		r := bufio.NewReader(os.Stdin)
		line, _ := r.ReadString('\n')
		totalArgs = strings.Fields(line)
	}

	for _, arg := range totalArgs {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		m := lengthconv.Meters(t)
		feet := lengthconv.Feet(t)
		kgs := weightconv.Kilograms(t)
		pounds := weightconv.Pounds(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), m, lengthconv.MToF(m), feet, lengthconv.FToM(feet), kgs, weightconv.KToP(kgs), pounds, weightconv.PToK(pounds))
	}

}
