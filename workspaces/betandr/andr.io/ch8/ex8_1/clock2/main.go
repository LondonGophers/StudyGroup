// Modify `clock2` to accept a port number, and write a program, `clockwall`, that
// acts as a client of clock servers at once, reading the times from each one and
// displaying the results in a table, akin to the wall of clocks seen in some
// business offices. If you have access to geographically distributed computers,
// run instances remotely; otherwise run local instances on different ports with
// fake time-zones.
// ```
//   $ TZ=US/Eastern     ./clock2 -port 8010 &
//   $ TZ=Asia/Tokyo     ./clock2 -port 8020 &
//   $ TZ=Europe/London  ./clock2 -port 8030 &
//   $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
// ```
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8000, "Port number to run clock")

func sendTime(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Print(err)
	}

	for {
		sendTime(conn)
	}
}
