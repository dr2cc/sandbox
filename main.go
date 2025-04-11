package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	network = "tcp"
	port    = ":8080"
)

func main() {
	// Преобразование сети и порта в TCP-адрес
	tcpAddr, _ := net.ResolveTCPAddr(network, port)

	// Открытие сокета-прослушивателя
	listener, _ := net.ListenTCP(network, tcpAddr)
	defer listener.Close()

	log.Printf("Прослушивание порта %s...\n", port)
	for {
		// Принятие TCP-соединения от клиента и создание нового соединения
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Обработка запросов клиента в отдельной горутине
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Подключен клиент %s\n", conn.RemoteAddr().String())

	connReader := bufio.NewReader(conn)
	for {
		// Чтение данных из соединения
		data, err := connReader.ReadString('\n')
		if err != nil {
			break
		}

		// Обработка сообщения от клиента
		message := strings.TrimSpace(string(data))
		log.Printf("Сообщение от %s: %s\n", conn.RemoteAddr().String(), message)
		if message == "STOP" {
			break
		}

		// Отправка данных в соединение
		conn.Write([]byte(data))
	}
	log.Printf("Отключен клиент %s\n", conn.RemoteAddr().String())
}
