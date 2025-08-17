package handle

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"server-application3/parse"
	"server-application3/request"
	"server-application3/respond"
	"strings"
	"time"
)

func Handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		conn.SetDeadline(time.Now().Add(5 * time.Minute)) // idle connection

		method, path, headers, body, err := request.Request(reader)
		if err != nil {
			if err == io.EOF {
				return // close connection from client side
			}
			log.Println("Request error:", err)
			respond.InternalServerError(conn)
			return
		}

		form := make(map[string]string)
		if method == "POST" && headers["content-type"] == "application/x-www-form-urlencoded" {
			form, _ = parse.ParseForm(body)
		}

		if method == "POST" && headers["content-type"] == "application/json" {
			var jsonData map[string]interface{}
			err := json.Unmarshal([]byte(body), &jsonData)
			if err != nil {
				log.Println("JSON parse error:", err)
				respond.BadRequest(conn)
				return
			}
			// использовать jsonData
		}

		log.Printf("Method: %s, Path: %s, Body: %q, Form: %v\n", method, path, body, form)
		log.Printf("[%s] %s %s %s\n", time.Now().Format(time.RFC3339), conn.RemoteAddr(), method, path)

		keepAlive := strings.ToLower(headers["connection"]) == "keep-alive"
		respond.Respond(conn, method, path, keepAlive, form)

		if !keepAlive {
			return
		}
	}
}

// Ваш TCP-сервер (Handle(conn net.Conn))

// Это низкоуровневый сервер, который принимает соединение через net.Conn.

// Он вручную парсит HTTP-запрос (request.Request(reader)), проверяет метод, заголовки, тело, формы и JSON.

// Ответы отправляются через respond.Respond(conn, ...).
