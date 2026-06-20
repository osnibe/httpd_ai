package main

import (
	"fmt"
	"net"
	"os"
)

const port = ":8080"

const response = "HTTP/1.1 200 OK\r\n" +
	"Content-Type: text/plain\r\n" +
	"Content-Length: 13\r\n" +
	"Connection: close\r\n" +
	"\r\n" +
	"Hello, World!"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		return
	}
	fmt.Printf("--- request ---\n%s\n", buf[:n])

	conn.Write([]byte(response))
}

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "listen:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Printf("Listening on %s\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, "accept:", err)
			continue
		}
		fmt.Printf("Connection from %s\n", conn.RemoteAddr())
		handleConnection(conn)
	}
}
