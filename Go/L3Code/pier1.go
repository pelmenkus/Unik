/*
185.139.70.64 / root / j78Ei7372PRf
185.104.249.105 / root / 3TsSnm37owBs
185.255.133.113 / root / aXAvAa5vd45F
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

type Peer struct {
	IP       string
	Port     int
	IsHead   bool
	NextPeer *Peer
}

type Bulletin struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	AuthorAddress string `json:"authorAddress"`
}

type BulletinBoard struct {
	Bulletins             []Bulletin `json:"bulletins"`
	LastChangePeerAddress string     `json:"lastChangePeer"`
}

func main() {

	// пиры для запуска на разных ip (серверах)
	peer1 := &Peer{IP: "185.139.70.64", Port: 1111, IsHead: true}
	peer2 := &Peer{IP: "185.104.249.105", Port: 2222, IsHead: false}
	peer3 := &Peer{IP: "185.255.133.113", Port: 3333, IsHead: false}

	peer1.NextPeer = peer2
	peer2.NextPeer = peer3
	peer3.NextPeer = peer1
	bulletinBoard := BulletinBoard{}

	port := getPort()

	currentPeer := peer1

	// при запуске пиров на разных ip можно использовать switch по ip
	// для этого нужно раскомментарить код ниже и закомментарить switch по port

	ip := getIPAddress()

	switch ip {
	case peer1.IP:
		currentPeer = peer1
	case peer2.IP:
		currentPeer = peer2
	case peer3.IP:
		currentPeer = peer3
	}

	peer1.Port = port
	peer2.Port = port
	peer3.Port = port

	go startPeer(currentPeer, &bulletinBoard)
	go startCommandInterface(currentPeer, &bulletinBoard)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Exiting...")
}

// получение IP-адреса устройства
func getIPAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// получение порта от пользователя
func getPort() int {
	port := 2222
	fmt.Print("Port: ")
	fmt.Scan(&port)

	address := fmt.Sprintf("localhost:%d", port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	localAddr := l.Addr().(*net.TCPAddr)

	return localAddr.Port
}

// Запуск пира
func startPeer(p *Peer, bulletinBoard *BulletinBoard) {
	address := fmt.Sprintf("%s:%d", p.IP, p.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Peer %s:%d failed to start: %v\n", address, err)
	}
	defer listener.Close()

	log.Printf("Peer started and listening on %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Peer %s encountered an error: %v\n", address, err)
			continue
		}
		go handleConnection(conn, p, bulletinBoard)
	}
}

// Обработка соединения
func handleConnection(conn net.Conn, p *Peer, bulletinBoard *BulletinBoard) {
	defer conn.Close()
	log.Println("New connection established")
	message := make([]byte, 4096)
	n, err := conn.Read(message)
	if err != nil {
		log.Println(err)
		return
	}
	var record BulletinBoard
	err = json.Unmarshal(message[:n], &record)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("The update of the bulletin board has been received")
	bulletinBoard.Bulletins = record.Bulletins
	bulletinBoard.LastChangePeerAddress = record.LastChangePeerAddress

	// отправляем обновление доски объявлений к другому соседу, если сосед
	// не является автором последнего изменения
	address := fmt.Sprintf("%s:%d", p.NextPeer.IP, p.NextPeer.Port)

	if address != bulletinBoard.LastChangePeerAddress {
		sendBulletinBoard(p, bulletinBoard)
	}
}

// Обработчик команд
func startCommandInterface(p *Peer, bulletinBoard *BulletinBoard) {
	for {
		var command string
		fmt.Print("Enter command (post, remove, list): ")
		fmt.Scanln(&command)

		switch command {
		case "post":
			addBulletin(p, bulletinBoard)
		case "remove":
			removeBulletin(p, bulletinBoard)
		case "list":
			listBulletins(bulletinBoard)
		default:
			log.Println("Invalid command")
		}
	}
}

// Отправка доски объявлений соседу
func sendBulletinBoard(p *Peer, bulletinBoard *BulletinBoard) {
	address := fmt.Sprintf("%s:%d", p.NextPeer.IP, p.NextPeer.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	jsonData, err := json.Marshal(bulletinBoard)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.Write(jsonData)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Sent bulletin board to neighbor %s\n", address)
}

// Добавление объявления
func addBulletin(p *Peer, bulletinBoard *BulletinBoard) {
	scanner := bufio.NewScanner(os.Stdin)
	var title, content string

	fmt.Print("Enter title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Enter content: ")
	scanner.Scan()
	content = scanner.Text()

	address := fmt.Sprintf("%s:%d", p.IP, p.Port)
	bulletin := Bulletin{Title: title, Content: content, AuthorAddress: address}

	// Удаление последнего объявления текущего пира
	for i := len(bulletinBoard.Bulletins) - 1; i >= 0; i-- {
		if bulletinBoard.Bulletins[i].AuthorAddress == address {
			bulletinBoard.Bulletins = append(bulletinBoard.Bulletins[:i], bulletinBoard.Bulletins[i+1:]...)
			break
		}
	}

	bulletinBoard.Bulletins = append(bulletinBoard.Bulletins, bulletin)
	bulletinBoard.LastChangePeerAddress = fmt.Sprintf("%s:%d", p.IP, p.Port)

	sendBulletinBoard(p, bulletinBoard)
}

// Удаление объявления по индексу в доске объявления
func removeBulletin(p *Peer, bulletinBoard *BulletinBoard) {
	var index int
	fmt.Print("Enter the index of the bulletin to remove: ")
	fmt.Scanln(&index)

	if index >= 0 && index < len(bulletinBoard.Bulletins) {
		bulletinBoard.Bulletins = append(bulletinBoard.Bulletins[:index], bulletinBoard.Bulletins[index+1:]...)
		bulletinBoard.LastChangePeerAddress = fmt.Sprintf("%s:%d", p.IP, p.Port)

		sendBulletinBoard(p, bulletinBoard)
	} else {
		log.Println("Invalid index")
	}
}

// Вывод списка объявлений в формате: заголовок, содержание, автор
func listBulletins(bulletinBoard *BulletinBoard) {
	for i, bulletin := range bulletinBoard.Bulletins {
		fmt.Printf("%d. %s\n%s\nAuthor: %s\n\n", i, bulletin.Title, bulletin.Content, bulletin.AuthorAddress)
	}
}
