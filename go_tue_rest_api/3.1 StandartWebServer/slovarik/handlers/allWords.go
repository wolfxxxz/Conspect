package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"slovarik/library"
)

// Инициализирует json кодировку //Пишем хедеры // установить тип данных
func initHeadersJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// api/v1/words
func GetAllWords(writer http.ResponseWriter, request *http.Request) {
	initHeadersJson(writer)
	log.Println("Хтось питає за всі слова")
	writer.WriteHeader(200) // Статус код для запроса
	//lThemes := library.QuantityThemes(l)
	json.NewEncoder(writer).Encode(library.L) // Сериализация и запись в writer (запись в браузер)
	//2
}
