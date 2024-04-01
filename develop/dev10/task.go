package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		return
	}

	var timeout time.Duration = 10 * time.Second
	var host string
	var port string

	if os.Args[1] == "--timeout" {
		timeout, _ = time.ParseDuration(os.Args[2])
		host = os.Args[3]
		port = os.Args[4]
	} else {
		host = os.Args[1]
		port = os.Args[2]
	}
	fmt.Println(host, port)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	go func() {
		_, _ = io.Copy(conn, os.Stdin)
		conn.Close()
	}()

	_, _ = io.Copy(os.Stdout, conn)
}
