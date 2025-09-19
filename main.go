package main

import "fmt"

type employer struct {
	name    string
	address string
	salary  int
}

type customer struct {
	name  string
	email string
}

type creater interface {
	create(user string, info string)
}

func (e employer) create(nick string, contact string) {
	e.name = nick
	e.address = contact
	e.salary = 99999
	fmt.Println("Employer", e)
}

func (c customer) create(nick string, contact string) {
	c.name = nick
	c.email = contact
	fmt.Println("Customer", c)
}

// // Функция для демонстрации работы с интерфейсом
// func processCreator(c creater, name, info string) {
// 	c.create(name, info)
// }

func main() {
	// Создаем экземпляры
	emp := employer{}
	cust := customer{}

	// ✅ Оба типа реализуют интерфейс creater
	var c creater

	// employer реализует creater
	c = emp
	c.create("Alice", "Moscow")

	// customer реализует creater
	c = cust
	c.create("Bob", "bob@email.com")

	// // Использование через функцию
	// processCreator(emp, "Charlie", "SPb")
	// processCreator(cust, "David", "david@email.com")

	// // Проверка типа через утверждение типа
	// if creator, ok := c.(customer); ok {
	// 	fmt.Println("Это customer:", creator)
	// }
}
