package main

import (
	"fmt"
	"net"
)

func main() {
	message := "Hello, I am a server"           // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545") // Создание TCP-сервера

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	//Обработка соединений в бесконечном цикле
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message)) // Отправка данных клиенту
		conn.Close()                // Закрытие соединения с клиентом
	}
}
