package main

import (
    "fmt"
    "net"
    "strconv"
    "strings"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    sum := 0

    buffer := make([]byte, 1024)
    data := ""

    for {
        n, err := conn.Read(buffer)
        if err != nil {
            break
        }

        chunk := string(buffer[:n])
        data += chunk

        if strings.Contains(data, "END") {
            // Обнаружен символ конца потока
            parts := strings.Split(data, "\n")
            for _, part := range parts {
                if part != "" && part != "END" {
                    number, err := strconv.Atoi(part)
                    if err != nil {
                        fmt.Println("Error parsing number:", err)
                        continue
                    }
                    sum += number
                }
            }

            //response := fmt.Sprintf("%d\n", sum)
			response := strconv.Itoa(sum)
            conn.Write([]byte(response))
            fmt.Printf("Received stream of numbers. Sum: %d\n", sum)
            break
        }
    }
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
