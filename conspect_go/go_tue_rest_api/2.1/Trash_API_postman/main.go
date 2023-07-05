package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	db   []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 22,
		Price:    500.5,
		Title:    "Pepperoni",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 25,
		Price:    650.23,
		Title:    "BBQ",
	}

	pizza3 := Pizza{
		ID:       3,
		Diameter: 22,
		Price:    450,
		Title:    "Margarita",
	}

	db = append(db, pizza1, pizza2, pizza3)
}

// 2.4
// Второй блок
type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

// Вспомогательная функция или модельный метод
func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

//Конец второго блока

type ErrorMessage struct {
	Message string `json:"message"`
}

// Третий блок
// конфигурация исполнителей
// http.ResponseWriter - куда записать и  *http.Request что записать
func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Хедери
	//Настройка хедера установить format вывода json
	writer.Header().Set("Content - Type", "aplication/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200) //statusCode
	json.NewEncoder(writer).Encode(db)
}

// http.ResponseWriter - куда записать и  *http.Request что записать
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content - Type", "aplication/json")
	// Read Id and convert
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "don`t use id supported int casting"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Trying to send to client pizza with id #:", id)
	pizza, ok := FindPizzaById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		msg := ErrorMessage{Message: "pizza with that id does not exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}
}

// Конец третьего блока

func main() {
	//Первый блок
	log.Println("Trying to start REST API pizza!")
	router := mux.NewRouter()
	//1 Если запрос localhost:8080/pizzas - то вызвать функцию GetAllPizzas(). Желательно подписать тип метода Methods("GET")
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	//2 http.ResponseWriter - куда записать и  *http.Request что записать
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	// конец первого блока
	log.Println("Router configured successfully! Let`s go!")
	// Безконечный цикл который мониторит запросы по адрессу ":"+port (:8080)
	log.Fatal(http.ListenAndServe(":"+port, router))
	//2.4

}
