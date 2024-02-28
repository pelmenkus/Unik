package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmcdole/gofeed"
)

type News struct {
	Title       string
	Description string
	Link        string
	PubDate     time.Time
}

func main() {
	db, err := sql.Open("mysql", "iu9networkslabs:Je2dTYr6@tcp(students.yss.su)/iu9networkslabs")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS kraev (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) CHARACTER SET utf8mb4 NOT NULL,
			description TEXT CHARACTER SET utf8mb4,
			link VARCHAR(255) CHARACTER SET utf8mb4 NOT NULL,
			pub_date DATETIME NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err, " DB EXEC")
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://vz.ru/rss.xml")
	if err != nil {
		log.Fatal(err, " FEED PARSE")
	}

	for _, item := range feed.Items {
		pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.Published)
		if err != nil {
			log.Printf("Ошибка парсинга даты публикации: %v", err)
			continue
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM kraev WHERE link = ?", item.Link).Scan(&count)
		if err != nil {
			log.Fatal(err, " SELECT")
		}

		if count == 0 {
			// Новость отсутствует в базе данных, добавляем ее
			_, err := db.Exec("INSERT INTO kraev (title, description, link, pub_date) VALUES (?, ?, ?, ?)",
				item.Title, item.Description, item.Link, pubDate)
			if err != nil {
				log.Printf("Ошибка при добавлении новости: %v", err)
			} else {
				log.Printf("Добавлена новость: %s", item.Title)
			}
		} else {
			log.Printf("Новость уже существует: %s", item.Title)
		}
	}
}
