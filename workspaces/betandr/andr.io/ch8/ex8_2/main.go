// Implement a concurrent File Transfer Protocol (FTP) server. The server should
// interpret commands from each client such as `cd` to change directory, `ls` to
// list a directory, `get` to send the contents of a file, and `close` to close the
// connection. You can use the standard `ftp` command as the client, or write your
// own.
//
// RFC959 https://tools.ietf.org/html/rfc959
//
// Minimum commands:
// USER, QUIT, PORT, TYPE, MODE, STRU, RETR, STOR, NOOP
//
// Other supported commands:
// PASV, LIST, CWD
//
// go run ./andr.io/ch8/ex8_2 &
// ftp localhost -p 8021
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp4", "localhost:8021")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}

// listen creates a listener and returns the port and IP as a string to that listener
func listen(c net.Conn) (string, error) {
	lnr, err := net.Listen("tcp4", "")
	if err != nil {
		return "", err
	}
	_, portS, err := net.SplitHostPort(lnr.Addr().String())
	hostS, _, err := net.SplitHostPort(c.LocalAddr().String())

	ipAddr, err := net.ResolveIPAddr("ip4", hostS)
	if err != nil {
		return "", err
	}

	port, err := strconv.ParseInt(portS, 10, 64)
	if err != nil {
		return "", err
	}

	ip := ipAddr.IP.To4()

	go handleListener(lnr)

	return fmt.Sprintf("%d,%d,%d,%d,%d,%d", ip[0], ip[1], ip[2], ip[3], port/256, port%256), nil
}

func handleListener(ln net.Listener) error {
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn) // however we're going to run out of connections eventually! :/
	}
}

// handle a request and send an appropriate response via the supplied net.Conn
func handle(c net.Conn, request string) {
	req := strings.Split(request, " ")
	if len(req) < 1 {
		send(c, "502 not implemented.")
		return
	}
	args := req[1:]
	cmd := req[0]

	switch cmd {
	case "USER":
		send(c, "230 User logged in, proceed.")
	case "QUIT":
		send(c, "221 Service closing control connection.")
	case "PASV":
		addr, err := listen(c)
		if err != nil {
			send(c, "451 Requested action aborted: local error in processing.")
			break
		}
		send(c, fmt.Sprintf("227 =%s", addr))
	case "NOOP":
		send(c, "200 Command okay.")
	case "LIST":
		p
		if len(args
		// todo: implement LIST
		send(c, fmt.Sprintf("502 LIST not yet implemented."))
	case "PORT":
		// todo: implement PORT
		send(c, fmt.Sprintf("502 PORT not yet implemented."))
	case "TYPE":
		// todo: implement TYPE
		send(c, fmt.Sprintf("502 TYPE not yet implemented."))
	case "MODE":
		// todo: implement MODE
		send(c, fmt.Sprintf("502 MODE not yet implemented."))
	case "STRU":
		// todo: implement STRU
		send(c, fmt.Sprintf("502 STRU not yet implemented."))
	case "RETR":
		// todo: implement RETR
		send(c, fmt.Sprintf("502 RETR not yet implemented."))
	case "STOR":
		// todo: implement STOR
		send(c, fmt.Sprintf("502 STOR not yet implemented."))
	case "CWD":
		// todo: implement CWD
		send(c, fmt.Sprintf("502 STOR not yet implemented."))
	default:
		send(c, fmt.Sprintf("502 %s not implemented.", cmd))
	}
}

// send a
func send(c net.Conn, cmd string) {
	_, err := fmt.Fprint(c, fmt.Sprintf("%s\r\n", cmd))
	if err != nil {
		log.Fatal("send command:", err)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	log.Printf("new connection: %v", c.RemoteAddr())
	send(c, "220 Service ready for new user.")

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		request := scanner.Text()
		log.Printf("recieved: [%s]", request)
		handle(c, request)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("reading standard input:", err)
	}
}
