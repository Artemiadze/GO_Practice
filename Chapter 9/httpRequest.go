/*
Для управления запросом и его параметрами в Go используется объект http.Request. Он позволяет установить различные настройки, добавить куки, заголовки, определить тело запроса. Для создания объекта http.Request применяется функция http.NewRequest():
	func NewRequest(method, url string, body io.Reader) (*Request, error)
Функция принимает три параметра. Первый параметр - тип запроса в виде строки ("GET", "POST"). Второй параметр - адрес ресурса. Третий параметр - тело запроса.
Для отправки объекта Request можно применять метод Do():
	Do(req *http.Request) (*http.Response, error)
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{} // создаем клиент
	req, err := http.NewRequest(
		"GET", "https://google.com", nil, // создаем запрос (GET-запрос на URL)
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req) // отправляем запрос и получаем ответ
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
