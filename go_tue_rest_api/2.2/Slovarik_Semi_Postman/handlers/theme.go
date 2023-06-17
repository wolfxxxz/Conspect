package handlers

import (
	"Slovarik_Semi_Postman/library"
	"encoding/json"
	"log"
	"net/http"
)

func ThemesWordsPrint(writer http.ResponseWriter, request *http.Request) {
	//Пишем хедеры // установить тип данных
	initHeadersJson(writer)
	log.Println("Хтось питає за теми слів")
	writer.WriteHeader(200) // Статус код для запроса
	lThemes := library.QuantityThemes(library.L)
	json.NewEncoder(writer).Encode(lThemes) // Сериализация и запись в writer (запись в браузер)
	//2
}
