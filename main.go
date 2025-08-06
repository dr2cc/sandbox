package main

import (
	"fmt"
	"net/http"
	"time"
)

//
// Еще context :
// https://pahanini.com/posts/go-context/
// https://habr.com/ru/companies/pt/articles/764850/
// https://habr.com/ru/companies/nixys/articles/461723/

// #
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

// // ## Функция, которая выводит Hello
// func printHello() {
// 	fmt.Println("Hello from printHello")
// }

// ### Печатает на стандартный вывод и отправляет int в канал
func printHello(ch chan int) {
	fmt.Println("Hello from printHello")
	// Посылает значение в канал
	ch <- 2
}

func main() {
	// // # Изучаю context Пример отсюда:
	// // https://gobyexample.com/context
	// http.HandleFunc("/hello", hello)
	// http.ListenAndServe(":8090", nil)
	// //****/////////////////

	// //## Встроенная горутина
	// // Определяем функцию внутри и вызываем ее
	// go func() {
	// 	fmt.Println("Hello inline")
	// }()
	// // Вызываем функцию как горутину
	// go printHello()
	// fmt.Println("Hello from main")

	// //### С каналами
	// Создаем канал. Для этого нам нужно использовать функцию make
	// Каналы могут быть буферизированными с заданным размером:
	// ch := make(chan int, 2), но это выходит за рамки данной статьи.
	ch := make(chan int)

	// Встроенная горутина. Определим функцию, а затем вызовем ее.
	// Запишем в канал по её завершению
	go func() {
		fmt.Println("Hello inline")
		// Отправляем значение в канал
		ch <- 1
	}()

	// Вызываем функцию как горутину
	go printHello(ch)
	fmt.Println("Hello from main")

	// Получаем первое значение из канала
	// и сохраним его в переменной, чтобы позже распечатать
	i := <-ch
	fmt.Println("Received ", i)

	// Получаем второе значение из канала
	// и не сохраняем его, потому что не будем использовать
	<-ch

}
