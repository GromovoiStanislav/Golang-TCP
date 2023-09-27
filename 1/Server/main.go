package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Буфер для чтения данных от клиента
	buffer := make([]byte, 1024)

	// Читаем данные от клиента
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении данных от клиента:", err)
		return
	}

	clientMessage := string(buffer[:n])
	fmt.Printf("Получено сообщение от клиента: %s\n", clientMessage)

	// Отправляем ответ клиенту
	response := "Привет, клиент TCP!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Ошибка при отправке ответа клиенту:", err)
		return
	}

	fmt.Println("Ответ отправлен клиенту")
}

func main() {
	// Адрес и порт, на котором сервер будет слушать
	serverAddr := "127.0.0.1:8080"

	// Создаем TCP-сервер
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Сервер запущен и слушает %s\n", serverAddr)

	for {
		// Принимаем входящее подключение от клиента
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err)
			continue
		}

		// Запускаем обработчик подключения в отдельной горутине
		go handleConnection(conn)
	}
}

// go run main.go

// go mod init mymodule
// go build -o myapp.exe
// ./myapp
// Ctrl+C

//go build main.go
// go build -o myapp.exe main.go
// ./myapp
// Ctrl+C
