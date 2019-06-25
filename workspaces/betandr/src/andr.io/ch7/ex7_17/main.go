// Extend `xmlselect` so that elements may be selected not just by name, but by
// their attributes too, in the manner of CSS, so that, for instance, an element
// like `<div id="page" class="wide">` could be selected by matching `id` or
// `class` as well as its name.
//
// WORK IN PROGRESS
// fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run andr.io/ch7/ex7_17 div div h2
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run andr.io/ch7/ex7_17 div div h2
func main() {
	dec := xml.NewDecoder(os.Stdin)
	// var stack []string // stack of element names
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if elements, ok := containsAll(stack, os.Args[1:]); ok {
				fmt.Printf("%s: %s\n", strings.Join(elements, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) ([]string, bool) {
	elements := []string{}

	for len(y) <= len(x) {
		if len(y) == 0 {
			return elements, true
		}
		if x[0].Name.Local == y[0] {
			elements = append(elements, x[0].Name.Local)
			y = y[1:]
		} else {
			for _, a := range x[0].Attr {
				if a.Name.Local == y[0] {
					elements = append(elements, a.Name.Local)
					y = y[1:]
				} else if a.Value == y[0] {
					elements = append(elements, a.Name.Local)
					y = y[1:]
				}
			}
		}

		x = x[1:]
	}
	return elements, false
}
