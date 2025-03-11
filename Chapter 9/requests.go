package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	// Формирование HTTP-запроса
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n"

	// Установка TCP-соединения
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	defer conn.Close() // Закрытие соединения после завершения работы

	// Отправка HTTP-запроса
	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	// Чтение ответа построчно
	reader := bufio.NewReader(conn)
	isBody := false

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			return
		}

		// Вывод строки заголовка
		fmt.Print(line)

		// Если дошли до пустой строки, значит, начинается тело ответа
		if strings.TrimSpace(line) == "" {
			isBody = true
		}

		// Если обнаружен заголовок Location (редирект), сообщаем об этом
		if strings.HasPrefix(line, "Location:") {
			fmt.Println("Перенаправление на:", strings.TrimSpace(line[9:]))
		}

		// Выходим из цикла после заголовков (можно убрать, если нужен весь HTML)
		if isBody {
			break
		}
	}

	fmt.Println("Done")
}
