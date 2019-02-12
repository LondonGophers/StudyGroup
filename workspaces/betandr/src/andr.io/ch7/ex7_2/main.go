// Write a function `CountingWriter` with the signature below that, given an
// `io.Writer`, returns a new `Writer` that wraps the original, and a pointer to an
// `int64` variable that any moment contains the number of bytes written to the new
// `Writer`.
// ```
// func CounterWriter(w io.Writer) (io.Writer, *int64)
// ```
package main

import (
	"fmt"
	"io"
	"os"
)

// WriterCounter wraps an `io.Writer` object and provides an `int64` counter
type WriterCounter struct {
	Counter int64
	Writer  io.Writer
}

func main() {
	w, c := CounterWriter(os.Stdout)
	fmt.Fprint(w, "I know not all\t\t")
	fmt.Printf("[%d bytes]\n", *c)
	fmt.Fprint(w, "that may be coming,\t")
	fmt.Printf("[%d bytes]\n", *c)
	fmt.Fprint(w, "but be it what it will,\t")
	fmt.Printf("[%d bytes]\n", *c)
	fmt.Fprint(w, "I’ll go to it laughing.\t")
	fmt.Printf("[%d bytes]\n", *c)
}

// Write counts the number of bytes written as it writes to a `WriterCounter.Writer`
func (c *WriterCounter) Write(p []byte) (int, error) {
	c.Counter += int64(len(p))
	return c.Writer.Write(p)
}

// CounterWriter wraps an `io.Writer` in a `WriterCounter` which provides counting
func CounterWriter(w io.Writer) (io.Writer, *int64) {
	wc := new(WriterCounter)
	wc.Counter = 0
	wc.Writer = w
	return wc, &wc.Counter
}
