// Implement a concurrent File Transfer Protocol (FTP) server. The server should
// interpret commands from each client such as `cd` to change directory, `ls` to
// list a directory, `get` to send the contents of a file, and `close` to close the
// connection. You can use the standard `ftp` command as the client, or write your
// own.
//
// Useful:
// https://tools.ietf.org/html/rfc959
// https://en.wikipedia.org/wiki/List_of_FTP_server_return_codes
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
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var port string

type client struct {
	port       string
	binaryMode bool
}

func main() {
	listener, err := net.Listen("tcp4", "localhost:8021")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s listener started: [%s]", listener.Addr().Network(), listener.Addr())

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

// decodes a port string such as 127,0,0,1,205,138 to an IP/port such as 127.0.0.1:52618
func addressFromPort(port string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(port, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2), nil
}

// handle a request and send an appropriate response via the supplied net.Conn
func (cl *client) handle(c net.Conn, request string) {
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
		filename := "."
		var filenames []string
		if len(args) > 0 {
			filename = args[0]
		}
		file, err := os.Open(filename)
		if err != nil {
			send(c, "550 File not found.")
			return
		}
		stat, err := file.Stat()
		if stat.IsDir() {
			filenames, err = file.Readdirnames(0)
			if err != nil {
				send(c, "550 File unavailable.")
				return
			}
		}

		send(c, "150 File status okay; about to open data connection.")
		if cl.port == "" {
			log.Printf("error: no port for LIST")
			return
		}
		conn, err := net.Dial("tcp", cl.port)
		defer conn.Close()
		if err != nil {
			send(c, fmt.Sprintf("425 Can't open data connection."))
			return
		}

		for _, f := range filenames {
			_, err = fmt.Fprint(conn, f, "\r\n")
		}

		if err != nil {
			send(c, "426 Connection closed: transfer aborted.")
			return
		}

		send(c, "226 LIST successful.")
	case "PORT":
		cl.port, _ = addressFromPort(args[0])
		log.Printf("port: [%s]", cl.port)
		send(c, fmt.Sprintf("200 PORT command successful."))
	case "TYPE":
		if len(args) > 0 {
			switch args[0] {
			case "I":
				cl.binaryMode = true
				send(c, fmt.Sprintf("200 Command OK."))
				break
			case "A":
				cl.binaryMode = false
				send(c, fmt.Sprintf("200 Command OK."))
				break
			default:
				send(c, fmt.Sprintf("501 Syntax error in parameters or arguments."))
			}
		}

	case "MODE":
		// todo: implement MODE
		send(c, "502 MODE not yet implemented.")
	case "STRU":
		// todo: implement STRU
		send(c, "502 STRU not yet implemented.")
	case "RETR":
		if len(args) != 1 {
			send(c, "501 Syntax error in parameters or arguments.")
		}
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			send(c, "550 Requested action not taken. File unavailable.")
			return
		}

		send(c, "150 File status okay; about to open data connection.")
		if cl.port == "" {
			log.Printf("error: no port for LIST")
			return
		}

		conn, err := net.Dial("tcp", cl.port)
		defer conn.Close()
		if err != nil {
			send(c, fmt.Sprintf("425 Can't open data connection."))
			return
		}

		if cl.binaryMode {
			_, err := io.Copy(conn, file)
			if err != nil {
				send(c, "450 Requested file action not taken.")
			}

		} else {
			log.Printf("starting ascii transfer")
		}

		send(c, "226 Requested file action successful.")

	case "STOR":
		// todo: implement STOR
		send(c, "502 STOR not yet implemented.")
	case "CWD":
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}
		os.Chdir(dir)
		send(c, "250 CWD command successful.")
	default:
		send(c, fmt.Sprintf("502 %s not implemented.", cmd))
	}
}

func send(c net.Conn, cmd string) {
	_, err := fmt.Fprint(c, fmt.Sprintf("%s\r\n", cmd))
	if err != nil {
		log.Fatal("send command:", err)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	log.Printf("new connection: [%v]", c.RemoteAddr())
	send(c, "220 Service ready for new user.")

	conn := new(client)
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		request := scanner.Text()
		log.Printf("recieved: [%s]", request)
		conn.handle(c, request)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("reading standard input: ", err)
	}
}
