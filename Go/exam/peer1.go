package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Peer struct {
	Name          string   `json:"name"`
	IPAddress     string   `json:"ip_address"`
	Port          int      `json:"port"`
	PossiblePeers []string `json:"possible_peers"`
}

func main() {
	peer := Peer{
		Name:      "Peer1",
		IPAddress: "185.139.70.64",
		Port:      8001,
		PossiblePeers: []string{
			"185.104.249.105:8001",
			"185.255.133.113:8001",
		},
	}

	go startServer(peer)

	// Чтение команд
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		switch command {
		case "input":
			sendMessage(peer)
		default:
			fmt.Println("Unknown command")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Запуск сервера
func startServer(peer Peer) {
	address := fmt.Sprintf("%s:%d", peer.IPAddress, peer.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

// Обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var message string
	err := decoder.Decode(&message)
	if err != nil {
		log.Println("Error decoding message:", err)
		return
	}

	// Добавление сообщения в список полученных
	printReceivedMessage(message)
}

// Отправка сообщения другому пиру
func sendMessage(peer Peer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter message text:")
	text, _ := reader.ReadString('\n')

	text = strings.TrimSpace(text)

	message := text

	// Отправка сообщения каждому возможному пиру
	for _, possiblePeer := range peer.PossiblePeers {
		err := sendMessageToPeer(possiblePeer, message)
		if err != nil {
			log.Println("Error sending message to peer:", possiblePeer, err)
		}
	}
}

// Отправка сообщения указанному пиру
func sendMessageToPeer(peer string, message string) error {
	conn, err := net.DialTimeout("tcp", peer, time.Second*2)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Запись сообщения в соединение
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(message)
	if err != nil {
		return err
	}

	return nil
}

// Вывод полученных сообщений
func printReceivedMessage(message string) {
	fmt.Printf("Message: %s\n", message)
}
