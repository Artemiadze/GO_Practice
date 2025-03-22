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
	"time"
)

func main() {
	// Отправка GET-запроса на URL и вывод тела ответа в стандартный вывод.
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body) // Копирование тела ответа в стандартный вывод.
}

/*
Для осуществления HTTP-запросов также может применяться структура http.Client. Чтобы отправить запрос к веб-ресурсу, можно использовать один из ее методов:
func (c *Client) Do(req *Request) (*Response, error) -> отправляет запрос и возвращает ответ.
func (c *Client) Get(url string) (resp *Response, err error) -> отправляет GET-запрос на URL и возвращает ответ.
func (c *Client) Head(url string) (resp *Response, err error) -> отправляет HEAD-запрос на URL и возвращает ответ.
func (c *Client) Post(url string, contentType string, body io.Reader) (resp *Response, err error) -> отправляет POST-запрос на URL с телом body.
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) -> отправляет POST-запрос на URL с данными формы.

Настройка клиента
Структура http.Client имеет ряд полей, которые позволяют настроить ее поведение:
- Timeout: устанавливает таймаут для запроса
- Jar: устанавливает куки, отправляемые в запросе
- Transport: определяет механиз выполнения запроса
*/
