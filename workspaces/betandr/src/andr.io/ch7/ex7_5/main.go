// The `LimitReader` function in the `io` package accepts an `io.Reader r` and a
// number of bytes `n`, and returns another `Reader` that reads from `r` but
// reports an end-of-file condition after `n` bytes. Implement it.
// ```
//   func LimitReader(r io.Reader, n int64) io.Reader
// ```
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader("Hello, World!")
	lr := LimitReader(r, 5)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading string: %v\n", err)
	}
	fmt.Printf("read: %s\n", string(b))
}

// LimitedReader wraps an `io.Reader` and the number of bytes `N` to read from it.
type LimitedReader struct {
	R io.Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

// LimitReader accepts an `io.Reader` `r` and a number of bytes `n`, and returns
// another `Reader` that reads from `r` but reports an end-of-file condition
// after `n` bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
