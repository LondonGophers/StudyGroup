// Using the token-based decoder API, write a program that will read an arbitrary
// XML document and construct a tree of generic nodes that represents it. Nodes are
// of two kinds: `CharData` nodes represent text strings, and `Element` nodes
// represent named elements and their attributes. Each element node has a slice of
// child nodes.

// You may find the following declarations helpful.
// ```
//   import "encoding/xml"

//   type Node interface{} // CharData or *Element

//   type CharData string

//   type Element struct {
//     Type      xml.Name
//     Attr      []xml.Attr
//     Children  []Node
//   }
// ```
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{} // CharData or *Element

type stack []Node

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// push adds an element onto the LIFO stack
func (s stack) push(element *Element) {
	s = append(s, element)
}

// pop removes the last element from a LIFO stack
func (s stack) pop() {
	s = s[:len(s)-1]
}

// peek returns the last element from a LIFO stack
func (s stack) peek() *Element {
	n := s[len(s)-1]
	e := n.(Element)
	return &e
}

func construct(dec *xml.Decoder, s stack) {
	tok, err := dec.Token()

	if err == io.EOF {
		return
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "error constructing: %v\n", err)
		os.Exit(1)
	}

	switch tok := tok.(type) {
	case xml.StartElement:
		element := Element{
			Type: tok.Name,
			Attr: tok.Attr,
		}
		e := s.peek()
		e.Children = append(e.Children, element)
		s.push(&element)

	case xml.CharData:
		data := strings.Trim(strings.Trim(string([]byte(tok)), " "), "\n")
		if len(data) > 0 {
			e := s.peek()
			e.Children = append(e.Children, data)
		}

	case xml.EndElement:
		s.pop()
	}

	construct(dec, s)
}

func traverse(node Node) {
	switch n := node.(type) {
	case Element:
		fmt.Println(n.Type.Local)
		for c := range n.Children {
			traverse(c)
		}
	case CharData:
		fmt.Println(n)
	}
}

// fetch https://pastebin.com/raw/ePEp6w2Y | go run andr.io/ch7/ex7_18
func main() {
	stack := make([]Node, 1)
	stack[0] = Element{}
	construct(xml.NewDecoder(os.Stdin), stack)
	traverse(stack[0])
}
