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
//
// go run ./andr.io/ch8/ex8_2 -dir=/path/to/working/directory &
// ftp localhost -p 8021
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var dir = flag.String("dir", ".", "Working incoming directory for files.")

type client struct {
	port       string
	binaryMode bool
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp4", "localhost:8021")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s listener started: [%s]", listener.Addr().Network(), listener.Addr())

	err = os.Chdir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("incoming directory: [%s]", wd)

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

// addressFromPort decodes a port string such as 127,0,0,1,205,138 to an IP/port such as 127.0.0.1:52618
func addressFromPort(port string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(port, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2), nil
}

// validPath ensures that you can't chdir above the starting dir
func validPath(newDir string) bool {
	if strings.HasPrefix(newDir, "/") {
		return false
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(filepath.Join(wd, newDir), *dir) {
		return true
	}
	return false
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
	case "RETR":
		if len(args) != 1 {
			send(c, "501 Syntax error in parameters or arguments.")
		}
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			send(c, "550 File unavailable.")
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
			r := bufio.NewReader(file)
			w := bufio.NewWriter(conn)

			for {
				l, isPrefix, err := r.ReadLine()
				if err != nil {
					if err == io.EOF {
						break
					}
					send(c, "450 Requested file action not taken.")
					return
				}
				w.Write(l)
				if !isPrefix {
					w.Write([]byte("\r\n"))
				}
			}
			w.Flush()
		}

		send(c, "226 Requested file action successful.")

	case "STOR":
		if len(args) != 1 {
			send(c, "501 Syntax error in parameters or arguments.")
			return
		}

		filename := args[0]
		file, err := os.Create(filename)
		if err != nil {
			send(c, "550 No file access.")
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
			send(c, "425 Can't open data connection.")
		}
		_, err = io.Copy(file, conn)
		if err != nil {
			send(c, "450 Requested file action not taken.")
		}

		send(c, "226 Requested file action successful.")
	case "CWD":
		newDir := *dir
		if len(args) > 0 {
			newDir = args[0]
		}
		if validPath(newDir) {
			os.Chdir(newDir)
			send(c, "250 Requested file action okay, completed.")
		} else {
			send(c, "550 Requested action not taken.")
		}
	case "RMD":
		if len(args) != 1 {
			send(c, "550 Requested action not taken.")
			return
		}
		filename := args[0]

		file, err := os.Open(filename)
		if err != nil {
			send(c, "550 File not found.")
			return
		}

		stat, err := file.Stat()
		if !stat.IsDir() {
			if validPath(filename) {
				err := os.Remove(filename)
				if err != nil {
					send(c, "550 Requested action not taken.")
					return
				}

				send(c, "250 Requested file action okay, completed.")
				return
			}

		}

		send(c, "450 Requested file action not taken.")

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
