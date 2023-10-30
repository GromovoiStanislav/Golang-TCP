package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		_, err := conn.Write([]byte(fmt.Sprintf("%d\n", number)))
		if err != nil {
			fmt.Println("Error sending number to the server:", err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error receiving square from the server:", err)
			return
		}

		square := string(buffer[:n])
		fmt.Printf("Square of %d: %s", number, square)
	}
}
