package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")

	defer conn.Close()

	str := make([]byte, 8)

	n, _ := conn.Read(str)

	if string(str[:n]) == "OK" {
		fmt.Print("Programm is working")
	}
}
