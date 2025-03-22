/*Для отправки запросов в пакете net/http определен ряд функций:
func Get(url string) (resp *Response, err error) -> отправляет GET-запрос на URL и возвращает ответ.
func Head(url string) (resp *Response, err error) -> отправляет HEAD-запрос на URL и возвращает ответ.
func Post(url string, contentType string, body io.Reader) (resp *Response, err error) -> отправляет POST-запрос на URL с телом body.
func PostForm(url string, data url.Values) (resp *Response, err error) -> отправляет POST-запрос на URL с данными формы.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Отправка GET-запроса на URL и вывод тела ответа в стандартный вывод.
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body) // Копирование тела ответа в стандартный вывод.
}
