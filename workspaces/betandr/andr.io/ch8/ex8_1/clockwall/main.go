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
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var newYorkTime string = "waiting..."
var tokyoTime string = "waiting..."
var londonTime string = "waiting..."

func handleConn(c net.Conn, clockName string) {
	bufReader := bufio.NewReader(c)
	for {
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		switch clockName {
		case "NewYork":
			newYorkTime = strings.TrimSuffix(string(bytes), "\n")
		case "Tokyo":
			tokyoTime = strings.TrimSuffix(string(bytes), "\n")
		case "London":
			londonTime = strings.TrimSuffix(string(bytes), "\n")
		}
		updateTime()
	}
}

func updateTime() {
	fmt.Println("\033[2J") // clear screen
	fmt.Printf("New York: %s\nTokyo: %s\nLondon: %s\n\n", newYorkTime, tokyoTime, londonTime)
}

// listen creates a listener
func listen(clockName, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, clockName)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall Place=hostname:8010 ...")
		os.Exit(0)
	}

	for _, clock := range os.Args[1:] { // start all clock listeners
		address := strings.Split(clock, "=")

		if len(address) != 2 {
			fmt.Fprintf(os.Stderr, "cannot handle arg: %s", clock)
			continue
		}
		go listen(address[0], address[1])
	}

	updateTime() // draw first clocks

	for {
		time.Sleep(time.Minute)
	}
}
