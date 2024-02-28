package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Request struct {
	Domain string `json:"domain"`
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var request Request
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Println(err)
			return
		}

		htmlCode, err := getHTML(request.Domain)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(htmlCode))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func getHTML(domain string) (string, error) {
	url := fmt.Sprintf("http://%s", domain)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	http.HandleFunc("/", handleConnection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
