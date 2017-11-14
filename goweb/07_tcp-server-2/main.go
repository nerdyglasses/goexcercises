// # Building upon the code from the previous exercise:

// In that previous exercise, we WROTE to the connection.

// Now I want you to READ from the connection.

// You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.

// Use bufio.NewScanner() to read from the connection.

// After all of the reading, include these lines of code:

// fmt.Println("Code got here.")
// io.WriteString(c, "I see you connected.")

// Launch your TCP server.

// In your **web browser,** visit localhost:8080.

// Now go back and look at your terminal.

// Can you answer the question as to why "I see you connected." is never written?

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Error while accepting the connection", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected\n")
}
