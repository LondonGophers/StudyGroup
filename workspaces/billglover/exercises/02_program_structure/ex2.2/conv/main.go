package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/go-london-user-group/study-group/workspaces/billglover/exercises/02_program_structure/ex2.2/conv/distance"
	"github.com/go-london-user-group/study-group/workspaces/billglover/exercises/02_program_structure/ex2.2/conv/length"
	"github.com/go-london-user-group/study-group/workspaces/billglover/exercises/02_program_structure/ex2.2/conv/temp"
)

func main() {

	register(new(temp.Converter))
	register(new(length.Converter))
	register(new(distance.Converter))

	var s = new(bufio.Scanner)
	if len(os.Args) == 1 {
		s = bufio.NewScanner(os.Stdin)

	} else {
		s = bufio.NewScanner(strings.NewReader(strings.Join(os.Args[1:], "\n")))
	}

	for s.Scan() {
		v, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert: unable to parse input: %v\n", err)
			os.Exit(1)
		}
		err = convert(v, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert: failed conversion:%v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "convert: unable to read input: %v\n", err)
		os.Exit(1)
	}
}

func convert(v float64, w io.Writer) error {
	for _, c := range conversions {
		txt, err := c.Convert(v)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%-15s: %s\n", c.Name(), txt)
	}

	return nil
}
