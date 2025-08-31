package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// // Константа подключения к БД (замените на свои реальные данные)
// const DATABASE_DSN = "postgres://user:password@localhost:5432/dbname?sslmode=disable"

func getDesc(ctx context.Context, db *sql.DB, id string) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT description FROM videos WHERE video_id = ?", id)
	var desc sql.NullString

	err := row.Scan(&desc)
	if err != nil {
		return "", err
	}
	if desc.Valid {
		return desc.String, nil
	}
	return "-----", nil
}

// Окончил на
// Расширение поддерживаемых типов

func main() {
	db, err := sql.Open("sqlite3", "video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//...

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// не забываем освободить ресурс
	defer cancel()

	// // делаем запрос
	// // QueryRowContext выполняет запрос, который, как ожидается, вернет не более одной строки (в нашем случае запрос:
	// // ВЫБРАТЬ Количество(все) как count ИЗ videos
	// // должен вернуть только одну строку- количество
	// row := db.QueryRowContext(ctx,
	// 	"SELECT COUNT(*) as count FROM videos")

	row := db.QueryRowContext(ctx,
		"SELECT title, views, channel_title "+
			"FROM videos ORDER BY views DESC LIMIT 1")
	var (
		title string
		views int
		chati string
	)
	// порядок переменных должен соответствовать порядку колонок в запросе
	err = row.Scan(&title, &views, &chati)
	if err != nil {
		panic(err)
	}
	//fmt.Println(getDesc(ctx, db, "0EbFotkXOiA"))
	fmt.Printf("%s | %d | %s \r\n", title, views, chati)
}
