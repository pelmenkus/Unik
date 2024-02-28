package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Request struct {
	Domain string `json:"domain"`
}

func main() {
	fmt.Print("Enter the domain name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	domain := scanner.Text()

	conn, _, err := websocket.DefaultDialer.Dial("ws://185.139.70.64:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	request := Request{Domain: domain}
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.WriteMessage(websocket.TextMessage, jsonRequest)
	if err != nil {
		log.Fatal(err)
	}

	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(string(response))
	if err != nil {
		return
	}

	time.Sleep(time.Second)
}
