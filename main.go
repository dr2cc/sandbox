//17.07.2025
//Только GET по http://localhost:8080/pizzas
//Возвращает список пицц

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pizza struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Пользовательский тип.
type Pizzas []Pizza

// func (ps Pizzas) FindByID(ID int) (Pizza, error) {
// 	for _, pizza := range ps {
// 		if pizza.ID == ID {
// 			return pizza, nil
// 		}
// 	}

// 	return Pizza{}, fmt.Errorf("Couldn't find pizza with ID: %d", ID)
// }

// type Order struct {
// 	PizzaID  int `json:"pizza_id"`
// 	Quantity int `json:"quantity"`
// 	Total    int `json:"total"`
// }

// type Orders []Order

// type ordersHandler struct {
// 	pizzas *Pizzas
// 	orders *Orders
// }

// func (oh ordersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch r.Method {
// 	case http.MethodPost:
// 		var o Order

// 		if len(*oh.pizzas) == 0 {
// 			http.Error(w, "Error: No pizzas found", http.StatusNotFound)
// 			return
// 		}

// 		err := json.NewDecoder(r.Body).Decode(&o)
// 		if err != nil {
// 			http.Error(w, "Can't decode body", http.StatusBadRequest)
// 			return
// 		}

// 		p, err := oh.pizzas.FindByID(o.PizzaID)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
// 			return
// 		}

// 		o.Total = p.Price * o.Quantity
// 		*oh.orders = append(*oh.orders, o)
// 		json.NewEncoder(w).Encode(o)
// 	case http.MethodGet:
// 		json.NewEncoder(w).Encode(oh.orders)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

type pizzasHandler struct {
	pizzas *Pizzas
}

// здесь только Get - List all pizzas on the menu: GET `/pizzas“
func (ph pizzasHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// //Переделал оригинал, с использованием the named path wildcard (именованного подстановочного знака пути)
	// //"/{id}"
	// //Что-бы создать конечную точку, я теперь проверяю здесь соответствие полученного пути нашей конечной точке
	// //Но пока (14.04.2025) не знаю как передать PathValue при тестировании.
	// id := r.PathValue("id")

	// А вот RequestURI получается и от клиента и из теста
	id := r.RequestURI
	if id == "/pizzas" {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			if len(*ph.pizzas) == 0 {
				http.Error(w, "Error: No pizzas found", http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(ph.pizzas)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	//var orders Orders

	pizzas := Pizzas{
		Pizza{
			ID:    1,
			Name:  "Pepperoni",
			Price: 12,
		},
		Pizza{
			ID:    2,
			Name:  "Capricciosa",
			Price: 11,
		},
		Pizza{
			ID:    3,
			Name:  "Margherita",
			Price: 10,
		},
	}

	mux := http.NewServeMux()
	// Так будет работать, но отдавать список пицц при обращеннии по любому адресу
	// Сделал условие в обработчике
	mux.Handle("/{id}", pizzasHandler{&pizzas})
	//mux.Handle("/pizzas", pizzasHandler{&pizzas})
	//mux.Handle("/orders", ordersHandler{&pizzas, &orders})

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
