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
	"sort"
	"strings"
	"text/tabwriter"
	"time"
	"unicode"
)

// clockTimes is the map of times recieved from clocks, in name => time format
var clockTimes map[string]string

func handleConn(c net.Conn, clockName string) {
	bufReader := bufio.NewReader(c)
	for {
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		clockTimes[clockName] = strings.TrimSuffix(string(bytes), "\n")
		updateTime()
	}
}

// expandTitle turns a string like NewYork into New York by finding a capital
// letter and inserting a space.
func camelToTitle(name string) string {
	for i, c := range name {
		if unicode.IsUpper(c) && i > 0 {
			name = name[:i] + " " + name[i:]
		}
	}

	return name
}

// updateTime updates the screen wallclock by clearing the screen then creating
// an ordered list of wallclocks.
func updateTime() {
	fmt.Println("\033[2J") // clear screen
	names := make([]string, 0, len(clockTimes))
	for name := range clockTimes { // get names as sort key
		names = append(names, name)
	}
	sort.Strings(names)

	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, '\t', 0)
	for _, name := range names {
		fmt.Fprintf(tw, format, camelToTitle(name), clockTimes[name])
	}
	tw.Flush()
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

	clockTimes = make(map[string]string)

	for _, clock := range os.Args[1:] { // start all clock listeners
		address := strings.Split(clock, "=")

		if len(address) != 2 {
			fmt.Fprintf(os.Stderr, "cannot handle arg: %s", clock)
			continue
		}
		clockTimes[address[0]] = "...\t"
		go listen(address[0], address[1])
	}

	updateTime() // draw first clocks

	for {
		time.Sleep(time.Minute)
	}
}
