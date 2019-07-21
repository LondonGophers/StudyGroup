// Write a program that reads a single expression from the standard input, prompts
// the user to provide values for any variables, then evaluates the expression in
// the resulting environment. Handle all errors gracefully.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"andr.io/ch7/ex7_15/eval"
)

func promptForValue(key string) float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter value for %s: ", key)
	r, _ := reader.ReadString('\n')
	r = strings.ReplaceAll(r, "\n", "")
	value, err := strconv.ParseFloat(r, 64)
	if err != nil {
		fmt.Printf("%s is not a float value\n", r)
		return promptForValue(key)
	}
	return value
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: evaluate \"{expression}\"")
		os.Exit(0)
	}
	expr, err := eval.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		os.Exit(1)
	}

	env := make(map[eval.Var]float64)
	for _, v := range expr.Vars() {
		env[v] = promptForValue(v.String())
	}

	result := expr.Eval(env)
	fmt.Printf("%s = %s\n", expr, strconv.FormatFloat(float64(result), 'f', -1, 64))
}
