// This is a simple HTTP server written in Go
// Please test it on the client side with the another program I made in Python 'python_client.py'

package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Handle client's requests here
	// For this example, we'll simply echo the received message back to the client.
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		message := buf[:n]
		fmt.Printf("Received message: %s\n", message)

		_, err = conn.Write(message)
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func main() {
	fmt.Println("Go Server started. Listening on port 8080...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}
