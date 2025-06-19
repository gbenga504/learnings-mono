package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	var buf [1024]byte

	// Read from the connection into the buffer
	_, err := conn.Read(buf[:])

	fmt.Printf("The server received ===> %s\n", string(buf[:]))

	// Write into the connection
	_, err = conn.Write([]byte("I am a server!"))

	if err != nil {
		return
	}
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9008")

	for {
		conn, _ := listener.Accept()
		go process(conn)
	}
}
