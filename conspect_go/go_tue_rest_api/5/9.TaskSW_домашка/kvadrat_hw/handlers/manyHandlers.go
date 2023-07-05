package handlers

import (
	"encoding/json"
	"kvadrat_hw/internal/models"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

var kvadrat models.KorKvadr

// Инициализирует json кодировку //Пишем хедеры // установить тип данных
func initHeadersJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// http.ResponseWriter - куда записать и  *http.Request что записать
func GetKvadrat(writer http.ResponseWriter, request *http.Request) {
	//Пишем хедеры // установить тип данных
	initHeadersJson(writer)

	log.Println("Хтось бажае отримати корінь: ")
	var c int = kvadrat.A + kvadrat.B + kvadrat.C
	if c == 0 {
		msg := Message{Message: "Перш ніж випить треба налить"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200) // Статус код для запроса
	var d int = (kvadrat.B * kvadrat.B) - (4 * kvadrat.A * kvadrat.C)
	if d > 0 {
		kvadrat.Nroots = 2
	} else if d == 0 {
		kvadrat.Nroots = 1
	} else if d < 0 {
		kvadrat.Nroots = 0
	}
	//------------------------------------------------------------------------------------------------
	//kvadrat.Nroots = kvadrat.A + kvadrat.B + kvadrat.C
	json.NewEncoder(writer).Encode(kvadrat)
}

func PostKvadrat(writer http.ResponseWriter, request *http.Request) {
	initHeadersJson(writer)
	log.Println("Шось хтось хоче додати ...")

	//NewDecoder - читаем с инпута в word
	err := json.NewDecoder(request.Body).Decode(&kvadrat)
	if err != nil {
		msg := Message{Message: "provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	var c int = kvadrat.A + kvadrat.B + kvadrat.C
	if c == 0 {
		msg := Message{Message: "Твої данні гімно, треба хоть шось більше нуля"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{Message: "201"}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}
