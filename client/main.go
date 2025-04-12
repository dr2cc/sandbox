package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	network = "udp"
	port    = ":8080"
)

func main() {
	// Преобразование сети и порта в UDP-адрес сервера
	udpAddr, _ := net.ResolveUDPAddr(network, port)

	// Создание UDP-соединения
	conn, _ := net.DialUDP(network, nil, udpAddr)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Сообщение серверу (или 'STOP' для завершения): ")

		// Считывание сообщения из стандартного потока ввода
		data, _ := reader.ReadString('\n')
		if data == "STOP\n" {
			break
		}

		// Отправка данных в соединение
		conn.Write([]byte(data))

		// Чтение данных из соединения
		buffer := make([]byte, 1024)
		n, _, _ := conn.ReadFromUDP(buffer)
		fmt.Printf("Ответ от сервера: %s", string(buffer[:n]))
	}
}
