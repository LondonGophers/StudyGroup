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

// Node is either a CharData or *Element
type Node interface{}

// CharData represents XML character data (raw text), in which XML
// escape sequences have been replaced by the characters they represent.
type CharData string

// Element represents an XML node with attributes and child nodes.
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func parse(dec *xml.Decoder) *Element {
	var stack []*Element
	root := &Element{Type: xml.Name{Local: "root"}} // add a root node we can return
	stack = append(stack, root)

	for {
		tok, err := dec.Token()

		if err == io.EOF { // end of input
			return root
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "error constructing: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			child := &Element{
				Type: tok.Name,
				Attr: tok.Attr,
			}
			parent := stack[len(stack)-1]                    // pop
			parent.Children = append(parent.Children, child) // push
			stack = append(stack, child)
		case xml.CharData:
			parent := stack[len(stack)-1] // pop
			parent.Children = append(parent.Children, CharData(tok))

		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		}
	}
}

// print starts from the specified node and prints nodes and traverses to child nodes
func print(node Node) {
	switch n := node.(type) {
	case *Element:
		fmt.Printf("element: %s has %d child(ren):\n", n.Type.Local, len(n.Children))
		for _, child := range n.Children {
			print(child)
		}
	case CharData:
		data := strings.Trim(strings.Trim(string([]byte(n)), " "), "\n")
		if len(data) > 1 {
			fmt.Printf("\tdata: %s\n", data)
		}
		return
	}
}

// fetch https://pastebin.com/raw/ePEp6w2Y | go run andr.io/ch7/ex7_18
func main() {
	node := parse(xml.NewDecoder(os.Stdin))
	print(node)
}
