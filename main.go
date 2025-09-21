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
	// Метод create интерфейса creater
	create(user string, info string)
}

func (e *employer) create(nick string, contact string) {
	e.name = nick
	e.address = contact
	e.salary = 99999
	fmt.Println("employer", *e)
	fmt.Println("employer after clean", e.clean())
}

func (e employer) clean() employer {
	e.name = ""
	e.address = ""
	e.salary = 0
	return e
}

func (c *customer) create(nick string, contact string) {
	c.name = nick
	c.email = contact
	fmt.Println("customer", *c)
}

func processCreator(c creater, name, data string) {
	c.create(name, data)
}

func main() {
	// создаю объекты типов employer и customer
	empl := &employer{}
	cust := &customer{}

	// var _ creater = &employer{}      // employer реализует creater
	// var _ creater = (*customer)(nil) // customer реализует creater

	//var c creater

	processCreator(empl, "Bob", "5 Avenue")

	processCreator(cust, "Todd", "tt@gg.hhh")

	// // employer реализует creater
	// c = empl
	// // реализуя такие же методы (поля объекта никак не связаны с интерфейсом)
	// // другие методы у типа могут быть отличные от интерфейса!
	// c.create("Bob", "5 Avenue")

	// c = cust
	// c.create("Todd", "tt@gg.hhh")

}
