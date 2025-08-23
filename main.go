package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// Строка запуска контейнера в докере
//docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres

// // Константа подключения к БД (замените на свои реальные данные)
// const DATABASE_DSN = "postgres://user:password@localhost:5432/dbname?sslmode=disable"

const DATABASE_DSN = "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable"

func checkDBAvailability() bool {
	db, err := sql.Open("postgres", DATABASE_DSN)
	if err != nil {
		log.Printf("DB connection error: %v", err)
		return false
	}
	defer db.Close()

	// Быстрая проверка соединения
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Printf("DB ping failed: %v", err)
		return false
	}

	return true
}

func main() {
	dbAvailable := checkDBAvailability()
	log.Printf("Database available: %v", dbAvailable)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := fmt.Sprintf("Server is running\nDatabase available: %v", dbAvailable)
		w.Write([]byte(status))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if dbAvailable {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Database unavailable"))
		}
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
