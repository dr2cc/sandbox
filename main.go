package main

import (
	"fmt"
	"net/http"
	"time"
)

// Изучаю context Пример отсюда:
// https://gobyexample.com/context
//
// Еще context :
// https://pahanini.com/posts/go-context/
// https://habr.com/ru/companies/pt/articles/764850/
// https://habr.com/ru/companies/nixys/articles/461723/

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
