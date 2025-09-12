package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
)

// Строка запуска pg в docker
//
//docker run -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -v sprint3:/var/lib/postgresql/data -d postgres

func main() {
	// // 1. подключение

	// //DriverName
	dn := "postgres"
	//dn := "sqlite3"

	// // DataSourceName
	dsn := "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable"
	//dsn := "video.db"

	db, err := sql.Open(dn, dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println(dn)
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

	// 2. получение данных из таблицы videos (any db - pg or sqlite)
	row := db.QueryRowContext(ctx,
		"SELECT title, views, channel_title "+
			"FROM videos ORDER BY views DESC LIMIT 1")
	var (
		title string
		views int
		chati string
	)

	// 3. Scan() "переводит" полученные данные в GO-типы
	// порядок переменных должен соответствовать порядку колонок в запросе
	err = row.Scan(&title, &views, &chati)
	if err != nil {
		panic(err)
	}
	//fmt.Println(getDesc(ctx, db, "0EbFotkXOiA"))
	fmt.Printf("%s | %d | %s \r\n", title, views, chati)
}

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
