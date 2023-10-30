package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	done := make(chan struct{})

	go func() {
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				close(done)
				return
			}

			numberStr := string(buffer[:n])
			var number int
			_, err = fmt.Sscanf(numberStr, "%d", &number)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}

			square := number * number
			conn.Write([]byte(fmt.Sprintf("%d\n", square)))
		}
	}()

	<-done
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
