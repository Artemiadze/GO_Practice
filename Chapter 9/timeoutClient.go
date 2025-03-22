// go run Server.go
// go run timeoutClient.go
package main

import (
	"fmt"
	"net"
	"time"
)

/*
- SetDeadline(t time.Time) error: устанавливает таймаут на все операции ввода-вывода. Для установки времени применяется структура time.Time

- SetReadDeadline(t time.Time) error: устанавливает таймаут на операции ввода в поток

- SetWriteDeadline(t time.Time) error: устанавливает таймаут на операции вывода из потока
В каком случае они могут пригодиться? В прошлой теме было рассмотрено взаимодействие сервера и клиента. Для чтения данных от клиента сервер использовал буфер фиксированного размера:
1. input := make([]byte, (1024 * 4))
2. n, err := conn.Read(input)
Однако в ряде ситуаций это не лучший способ, особенно когда размер передаваемых данных превышает размер буфера. Мы можем точно не знать, сколько данных возвратит нам сервер. Поэтому определим следующий код клиента:*/

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// получение данных выделено в отдельный цикл for:
	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// получем ответ
		fmt.Print("Перевод:")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5)) // устанавливается таймаут в 5 секунд:
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
		fmt.Println()
	}
}
