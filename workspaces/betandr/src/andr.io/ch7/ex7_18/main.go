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
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// fetch https://pastebin.com/raw/ePEp6w2Y | go run andr.io/ch7/ex7_18
func main() {
	dec := xml.NewDecoder(os.Stdin)

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
			element := Element{
				Type: tok.Name,
				Attr: tok.Attr,
			}
			fmt.Printf("%v ", element)
		case xml.CharData:
			fmt.Printf("%s", tok)
		}
	}
}
