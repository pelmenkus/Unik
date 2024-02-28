package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func main() {
	to := getInput("To: ")
	subject := getInput("Subject: ")
	body := getInput("Message body: ")

	username := "dts21@dactyl.su"
	password := "12345678990DactylSUDTS"
	smtpServer := "mail.nic.ru"
	smtpPort := 465

	// Формируем тело сообщения
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)

	// Создаем конфигурацию TLS
	tlsConfig := &tls.Config{
		ServerName:         smtpServer,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpServer, smtpPort), tlsConfig)
	if err != nil {
		fmt.Println("Ошибка при установке TLS-соединения:", err)
		return
	}

	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		fmt.Println("Ошибка при создании SMTP-клиента:", err)
		return
	}

	// Аутентификация
	auth := smtp.PlainAuth("", username, password, smtpServer)
	if err = client.Auth(auth); err != nil {
		fmt.Println("Ошибка при аутентификации:", err)
		return
	}

	// Отправляем сообщение
	if err = client.Mail(username); err != nil {
		fmt.Println("Ошибка при установке отправителя:", err)
		return
	}

	if err = client.Rcpt(to); err != nil {
		fmt.Println("Ошибка при установке получателя:", err)
		return
	}

	w, err := client.Data()
	if err != nil {
		fmt.Println("Ошибка при начале передачи данных:", err)
		return
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Println("Ошибка при записи данных:", err)
		return
	}

	err = w.Close()
	if err != nil {
		fmt.Println("Ошибка при закрытии соединения:", err)
		return
	}

	fmt.Println("Проверочное сообщение успешно отправлено.")
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
