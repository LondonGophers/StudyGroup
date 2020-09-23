// Using a `select` statement, add a timeout to the echo server from Section 8.3
// so that it disconnects any client that shouts nothing within 10 seconds.
//
// go run ./andr.io/ch8/ex8_3/
// go run ./andr.io/ch8/ex8_8/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var active bool

	go func() {
		tick := time.Tick(10 * time.Second)
		for {
			select {
			case <-tick:
				if !active {
					fmt.Fprintln(c, "\t", "Goodbye!")
					c.Close()
					return
				}
				active = false
			}
		}
	}()

	input := bufio.NewScanner(c)

	for input.Scan() {
		active = true
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
