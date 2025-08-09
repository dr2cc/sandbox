package main

import (
	"fmt"
	"net/http"
	"time"
)

//
// Еще context :
// https://pahanini.com/posts/go-context/
//
// Здесь интересный пример
// https://habr.com/ru/companies/pt/articles/764850/ - Пакет context в Go: взгляд профессионала
//
// // Здесь те примеры, что ## и ###
// https://habr.com/ru/companies/nixys/articles/461723/

// // ## Функция, которая выводит Hello
// func printHello() {
// 	fmt.Println("Hello from printHello")
// }

// // ### Печатает на стандартный вывод и отправляет int в канал
// func printHello(ch chan int) {
// 	fmt.Println("Hello from printHello")
// 	// Посылает значение в канал
// 	ch <- 2
// }

// #
func hello(w http.ResponseWriter, req *http.Request) {

	// context.Context создается для каждого запроса механизмом net/http
	// и доступен с помощью метода Context()
	ctx := req.Context()
	fmt.Println("drk: hello handler started")
	// defer откладывает выполнение конструкции за ним
	// до окончания выполнения текущей функции
	defer fmt.Println("drk: hello handler ended")

	// select это switch для каналов
	select {

	// func time.After(d Duration) <-chan time.Time
	// возвращает канал типа <-chan Time , который:
	// - Закрывается автоматически после указанного времени (10 * time.Second).
	// - Не требует ручного создания или закрытия.
	//
	// Если действие (контекс) не отменено (к примеру Ctrl+C) то через 10 секунд срабатывает time.After(), выполняется первая ветка
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello after ten sec\n")
	//
	// Метод Done() из контекста (ctx.Done())
	// возвращает канал типа <-chan struct{}, который:
	// - Закрывается при отмене контекста (например, если клиент разорвал соединение).
	// - Управляется самим контекстом, не требует вашего вмешательства.
	//
	// Если контекст отменяется раньше чем сработал time.After() (клиент закрыл соединение),
	// срабатывает ctx.Done()
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("handler hello err:", err)

		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	// Этот эндпойнт из примера:
	// https://gobyexample.com/http-server
	// Но я в него добавил context и каналы!
	ctx := req.Context()
	fmt.Println("drk: heders started")
	defer fmt.Println("drk: heders stopped")

	select {
	case <-time.After(10 * time.Second):
		for name, headers := range req.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	case <-ctx.Done():
		fmt.Println("headers canceled")
	}
}

func main() {
	// # Изучаю context Пример отсюда:
	// https://gobyexample.com/context
	// // Вызов в Git bash (в фоновом режиме)
	//$ go run main.go &
	// // Имитация клиентского запроса
	//$ curl localhost:8090/hello
	// // Остановка фонового процесса
	//$ kill <PID>
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe("localhost:8090", nil)
	// # //

	// //## Встроенная горутина
	// // Определяем функцию внутри и вызываем ее
	// go func() {
	// 	fmt.Println("Hello inline")
	// }()
	// // Вызываем функцию как горутину
	// go printHello()
	// fmt.Println("Hello from main")

	// //### С каналами
	// // Создаем канал. Для этого нам нужно использовать функцию make
	// // Каналы могут быть буферизированными с заданным размером:
	// // ch := make(chan int, 2), но это выходит за рамки данной статьи.
	// // а в статье "Пакет context в Go: взгляд профессионала" она упоминается
	// ch := make(chan int)

	// // Встроенная горутина. Определим функцию, а затем вызовем ее.
	// // Запишем в канал по её завершению
	// go func() {
	// 	fmt.Println("Hello inline")
	// 	// Отправляем значение в канал
	// 	ch <- 1
	// }()

	// // Вызываем функцию как горутину
	// go printHello(ch)
	// fmt.Println("Hello from main")

	// // Получаем первое значение из канала
	// // и сохраним его в переменной, чтобы позже распечатать
	// i := <-ch
	// fmt.Println("Received ", i)

	// // Получаем второе значение из канала
	// // и не сохраняем его, потому что не будем использовать
	// <-ch

}
