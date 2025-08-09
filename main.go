package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Драйвер SQLite
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Открываем соединение с SQLite (файл БД будет создан, если его нет)
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка подключения к БД: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close() // Закрываем соединение при выходе из функции

	// Проверяем соединение (Ping)
	err = db.Ping()
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка ping БД: %v", err), http.StatusInternalServerError)
		return
	}

	// Если дошли сюда — соединение работает
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Соединение с SQLite успешно!")
}

func main() {
	// Регистрируем обработчик
	http.HandleFunc("/health", healthCheckHandler)

	// Запускаем сервер
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//Запрос-
	//$ curl localhost:8080/health
}
