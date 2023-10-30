package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Адрес сервера и порт, к которому мы хотим подключиться
	serverAddr := "127.0.0.1:8080"

	// Создаем соединение с сервером
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Сообщение, которое мы хотим отправить серверу
	message := "Привет, сервер TCP!"

	// Отправляем сообщение серверу
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
		os.Exit(1)
	}

	// Буфер для чтения ответа от сервера
	response := make([]byte, 1024)

	// Читаем ответ от сервера
	bytesRead, err := conn.Read(response)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа от сервера:", err)
		os.Exit(1)
	}

	// Преобразуем ответ в строку и выводим его
	fmt.Println("Ответ от сервера:", string(response[:bytesRead]))
}

//go run main.go
