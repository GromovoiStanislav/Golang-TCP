package main

import (
	"fmt"
	"net"
	"strconv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	xStr := string(buffer[:n])
	x, err := strconv.Atoi(xStr)
	if err != nil {
		fmt.Println("Error parsing X:", err)
		return
	}

	fmt.Printf("Client requested numbers from 1 to %d\n", x)

	for i := 1; i <= x; i++ {
		_, err := conn.Write([]byte(fmt.Sprintf("%d\n", i)))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}

	fmt.Println("Stream sent to client.")
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server listening on :8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
