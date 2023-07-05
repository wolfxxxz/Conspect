package main

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:8080/

// w - responseWriter - куда писать ответ
// r - request - откуда брать запрос
// обработчик или роутер
func GetGreat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi! I'm new web-server!</h1>") // <h1>___</h1> увеличение маштаба
}

// товарищ который выбирает нужный обработчик в зависимости от запроса
// Хендлер
func RequestHandler() {
	http.HandleFunc("/", GetGreat)               // Если придет запрос по адресу "/" то вызывай GetGreat
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускает веб сервер в режиме слушания
}

func main() {
	RequestHandler()
}

/*
1. Terminal => go run main.go
2. browser => localhost:8080

Hi! I'm new web-server!
*/
