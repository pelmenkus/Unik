package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"proto"
	"strconv"

	"github.com/skorobogatov/input"
)

// interact - функция, содержащая цикл взаимодействия с сервером.
func interact(conn *net.TCPConn) {
	defer conn.Close()
	encoder := json.NewEncoder(conn)
	for {
		// Чтение команды из стандартного потока ввода
		fmt.Printf("command = ")
		command := input.Gets()

		// Отправка запроса.
		switch command {
		case "quit":
			send_request(encoder, "quit", nil)
			return
		case "add":
			for {
				fmt.Printf("digit = ")
				str := input.Gets()
				digit, err := strconv.Atoi(str)
				if err == nil && digit < 10 && digit >= 0 {
					var number proto.Number
					number.Digit = digit
					send_request(encoder, "add", &number)
					printResp(conn)
				} else {
					if str == "end" {
						break
					} else {
						fmt.Printf(`error: Type digit or "end"\n`)
					}
				}
			}
			continue
		case "count":
			fmt.Printf("digit = ")
			digit, _ := strconv.Atoi(input.Gets())
			var number proto.Number
			number.Digit = digit
			send_request(encoder, "count", &number)
		default:
			fmt.Printf("error: unknown command\n")
			continue
		}
		printResp(conn)
	}
}

// обработка ответа
func printResp(conn *net.TCPConn) {
	var resp proto.Response
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&resp); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Вывод ответа в стандартный поток вывода.
	switch resp.Status {
	case "ok":
		fmt.Printf("ok\n")
	case "failed":
		if resp.Data == nil {
			fmt.Printf("error: data field is absent in response\n")
		} else {
			var errorMsg string
			if err := json.Unmarshal(*resp.Data, &errorMsg); err != nil {
				fmt.Printf("error: malformed data field in response\n")
			} else {
				fmt.Printf("failed: %s\n", errorMsg)
			}
		}
	case "result":
		if resp.Data == nil {
			fmt.Printf("error: data field is absent in response\n")
		} else {
			var number proto.Number
			if err := json.Unmarshal(*resp.Data, &number); err != nil {
				fmt.Printf("error: malformed data field in response\n")
			} else {
				fmt.Printf("result: " + strconv.Itoa(number.Digit) + "\n")
			}
		}
	default:
		fmt.Printf("error: server reports unknown status %q\n", resp.Status)
	}
}

// send_request - вспомогательная функция для передачи запроса с указанной командой
// и данными. Данные могут быть пустыми (data == nil).
func send_request(encoder *json.Encoder, command string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	encoder.Encode(&proto.Request{command, &raw})
}

func main() {
	// Работа с командной строкой, в которой может указываться необязательный ключ -addr.
	var addrStr string
	flag.StringVar(&addrStr, "addr", "127.0.0.1:6000", "specify ip address and port")
	flag.Parse()

	// Разбор адреса, установка соединения с сервером и
	// запуск цикла взаимодействия с сервером.
	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		fmt.Printf("error: %v\n", err)
	} else if conn, err := net.DialTCP("tcp", nil, addr); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		interact(conn)
	}
}
