package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {


	x := os.Args[1]
	fmt.Println("x:", x)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(x))
	if err != nil {
		fmt.Println("Error sending X to the server:", err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error receiving data from the server:", err)
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
