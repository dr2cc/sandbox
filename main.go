package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

// App предоставляет основное приложение
type App struct {
	DB *sql.DB
}

// Инициализация подключения к PostgreSQL
func initDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connection error: %v", err)
	}

	// Настройки пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Проверяю подключение с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error to ping: %v", err)
	}

	return db, nil
}

// Обработчик проверки "здоровья" БД
func (a *App) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// test connect to db
	if err := a.DB.PingContext(ctx); err != nil {
		respondWithError(w, http.StatusServiceUnavailable, fmt.Sprintf("Error connecting to DB: %v", err))
		return
	}

	// Executing a test query
	var result int
	if err := a.DB.QueryRowContext(ctx, "SELECT 1").Scan(&result); err != nil {
		respondWithError(w, http.StatusServiceUnavailable, fmt.Sprintf("Test request error: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "available",
		"message": fmt.Sprintf("The test request returned: %d", result),
	})
}

// Helper (вспомогательные) functions for answers
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func main() {
	// Getting DSN from environment variables
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN not specified in env")
	}

	// DB init
	db, err := initDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Creating an app
	app := &App{DB: db}

	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// routes
	router.Get("/ping", app.healthCheckHandler)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	// HTTP server settings
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Server started on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Sever error: %v", err)
	}
}

//Формат: postgres://user:password@host:port/dbname?sslmode=disable
//в моем случае:
//export DATABASE_DSN="postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable"

// Строка запуска pg в docker
//docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres
