package main

// Строка запуска pg в docker
//docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

const (
	host     = "localhost"
	port     = 5436
	user     = "postgres"
	password = "qwerty" //"yourpassword"
	dbname   = "postgres"
	//
	//username = "postgres"
	//host: "localhost"
	//port: "5436"
	//dbname: "postgres"
	//sslmode = "disable"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Формируем строку подключения
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Открываем соединение
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка подключения: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Устанавливаем таймаут для Ping
	db.SetConnMaxLifetime(time.Second * 5)

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка ping: %v", err), http.StatusInternalServerError)
		return
	}

	// Проверяем доступность простым запросом
	var result int
	err = db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка запроса: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "PostgreSQL доступен! Результат теста: %d", result)
}

// curl http://localhost:8080/ping
func main() {
	router := chi.NewRouter()
	router.Get("/ping", healthCheckHandler)
	//
	//http.HandleFunc("/ping", healthCheckHandler)
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
