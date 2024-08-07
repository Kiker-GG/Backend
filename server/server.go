package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")

	conn, err := listener.Accept()

	if err != nil {
		fmt.Print("Error")
	}

	defer conn.Close()

	conn.Write([]byte("OK"))
}
