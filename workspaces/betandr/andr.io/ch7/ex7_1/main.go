// Using the ideas from `ByteCounter`, implement counters for words and for lines.
// You will find `bufio.ScanWords` useful.
package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ByteCounter represents a count of bytes
type ByteCounter int

// WordCounter represents a count of words
type WordCounter int

// LineCounter represents a count of lines
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	words := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	*c += WordCounter(words) // convert int to WordCounter
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	lines := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines++
	}
	*c += LineCounter(lines) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	var d WordCounter
	d.Write([]byte("hello"))
	fmt.Println(d) // "1", = "hello"

	d = 0 // reset the counter
	name = "Dolly"
	fmt.Fprintf(&d, "hello, %s", name)
	fmt.Println(d) // "2", = "hello, Dolly"

	var e LineCounter
	e.Write([]byte("hello"))
	fmt.Println(e) // "1", = "hello"

	e = 0 // reset the counter
	name = "Dolly"
	fmt.Fprintf(&e, "hello\nthere\n%s", name)
	fmt.Println(e) // "3", = "hello" \n "there" \n "Dolly"
}
