// In `netcat3`, the interface value `conn` has the concrete type `*net.TCPConn`,
// which represents a TCP connection. A TCP connection consists of two halves that
// may be closed independently using it's `CloseRead` and `CloseWrite` methods.
// Modify the main goroutine of `netcat3` to close only the write half of the
// connection so that the program will continue to print the final echoes from the
// `reverb1` server even after the standard input has been closed. (Doing this for
// the `reverb` server is harder; see Exercise 8.4.)
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	c := conn.(*net.TCPConn)
	c.CloseWrite()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
