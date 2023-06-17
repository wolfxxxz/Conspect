package handlers

import (
	"Semitrash_API/models"
	"encoding/json"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get infos about all books in databasa")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
