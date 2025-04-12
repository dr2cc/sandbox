package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	network = "tcp"
	port    = ":8080"
)

func main() {
	// Преобразование сети и порта в TCP-адрес
	tcpAddr, _ := net.ResolveTCPAddr(network, port)

	// Создание соединения с сервером по TCP-адресу
	conn, _ := net.DialTCP(network, nil, tcpAddr)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	//loop
	for {

		fmt.Print("Сообщение серверу (или 'stop' для завершения): ")

		// Считывание сообщения из стандартного потока ввода
		data, _ := reader.ReadString('\n')

		if strings.TrimSpace(string(data)) == "stop" {
			break
		}

		// Отправка данных в соединение
		conn.Write([]byte(data))

		// Чтение данных из соединения
		message, _ := connReader.ReadString('\n')

		fmt.Print("Ответ от сервера: " + message)

	}
}
