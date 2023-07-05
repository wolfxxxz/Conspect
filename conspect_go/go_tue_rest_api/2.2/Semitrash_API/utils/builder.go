package utils

import (
	"Semitrash_API/handlers"

	"github.com/gorilla/mux"
)

// Router
func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")       // Показать по id
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")               // Создать
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")    // Обновить
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods("DELETE") // Удалить
}

func BuildManyBooksResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET") // Показать все
}
