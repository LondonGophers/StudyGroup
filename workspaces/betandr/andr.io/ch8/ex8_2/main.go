// Implement a concurrent File Transfer Protocol (FTP) server. The server should
// interpret commands from each client such as `cd` to change directory, `ls` to
// list a directory, `get` to send the contents of a file, and `close` to close the
// connection. You can use the standard `ftp` command as the client, or write your
// own.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8021")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	bufReader := bufio.NewReader(c)
	for {
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("recieved: %s", string(bytes))
	}
}
