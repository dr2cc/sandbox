package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5436
// 	user     = "postgres"
// 	password = "qwerty" //"yourpassword"
// 	dbname   = "postgres"
// )

// DBConfig содержит параметры подключения к БД
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// App представляет основное приложение
type App struct {
	DB *sql.DB
}

// Инициализация подключения к PostgreSQL
func initDB(config DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %v", err)
	}

	// Настройки пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Проверяем подключение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ошибка ping: %v", err)
	}

	return db, nil
}

// Обработчик проверки "здоровья" БД
func (a *App) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем соединение с контекстом запроса
	ctx := r.Context()
	err := a.DB.PingContext(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка соединения с БД: %v", err), http.StatusServiceUnavailable)
		return
	}

	// Выполняем тестовый запрос
	var result int
	err = a.DB.QueryRowContext(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка тестового запроса: %v", err), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "PostgreSQL доступен. Тестовый запрос вернул: %d", result)
}

// Вспомогательная функция для получения переменных окружения
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Конфигурация БД
	config := DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     5436, //5432,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "postgres"),
	}

	// Инициализация БД
	db, err := initDB(config)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	// Создаем приложение
	app := &App{DB: db}

	// Инициализация роутера chi
	r := chi.NewRouter()

	// Промежуточное ПО
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Маршруты
	r.Get("/ping", app.healthCheckHandler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Добро пожаловать!"))
	})

	// Настройка HTTP-сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("Сервер запущен на :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}

// Строка запуска pg в docker
//docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres

// // in git bash
// export DB_HOST=localhost
// export DB_USER=postgres
// export DB_PASSWORD=qwerty
// export DB_NAME=postgres
