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

    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    for _, number := range numbers {
        numberStr := fmt.Sprintf("%d\n", number)
        _, err := conn.Write([]byte(numberStr))
        if err != nil {
            fmt.Println("Error sending number to the server:", err)
            return
        }
    }

    // Отправляем символ конца потока
    _, err = conn.Write([]byte("END\n"))
    if err != nil {
        fmt.Println("Error sending end symbol to the server:", err)
        return
    }

    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error receiving sum from the server:", err)
        return
    }

    sum := string(buffer[:n])
    fmt.Printf("Sum: %s", sum)
}
