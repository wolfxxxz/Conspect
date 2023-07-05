package handlers

import (
	"Slovarik_Semi_Postman/library"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// http.ResponseWriter - куда записать и  *http.Request что записать
func GetWordById(writer http.ResponseWriter, request *http.Request) {
	//Пишем хедеры // установить тип данных
	initHeadersJson(writer)
	// Как получить ID Запроса
	vars := mux.Vars(request) // возвращает мап {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Client trying to use invalid id param", err)
		msg := library.ErrorMessage{Message: "do not use id not supported int casting"}
		writer.WriteHeader(400) //bed request
		json.NewEncoder(writer).Encode(msg)
		return
	}

	log.Println("Хтось питає за слово під номером: ", id)
	word, ok := library.FindWordById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(word)
	} else {
		msg := library.ErrorMessage{Message: "word with that id does'nt exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}

	writer.WriteHeader(200) // Статус код для запроса
	//h := FindWordById()
	//json.NewEncoder(writer).Encode()
}

func CreateWord(writer http.ResponseWriter, request *http.Request) {
	initHeadersJson(writer)
	log.Println("Creating new book ....")
	var word library.Library

	//NewDecoder - читаем с инпута в word
	err := json.NewDecoder(request.Body).Decode(&word)
	if err != nil {
		msg := library.Message{Message: "provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBookID int = len(library.L) + 1
	word.ID = newBookID
	library.L = append(library.L, word)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(word)
}

func UpdateWordById(writer http.ResponseWriter, request *http.Request) {
	initHeadersJson(writer)
	log.Println("Updating book...")
	//Достать id запроса
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := library.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	oldword, ok := library.FindWordById(id)
	var newWord library.Library
	if !ok {
		log.Println("Book not found in data base. id:", id)
		writer.WriteHeader(404)
		msg := library.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Достать Слово с Json входящий
	errr := json.NewDecoder(request.Body).Decode(&newWord)
	log.Println(newWord)
	if errr != nil {
		msg := library.Message{Message: "provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//нужно заменить oldWord на newWord in []L.Library
	library.ReWriteWord(oldword.ID, newWord)
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(newWord)
}

func DeleteWordById(writer http.ResponseWriter, request *http.Request) {
	initHeadersJson(writer)
	log.Println("Deleting book...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := library.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	oldword, ok := library.FindWordById(id)
	if !ok {
		log.Println("Book not found in data base. id:", id)
		writer.WriteHeader(404)
		msg := library.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Удалить book from DB
	library.DelWordFromL(oldword.ID)
	msg := library.Message{Message: "successfully deleted requested item"}
	json.NewEncoder(writer).Encode(msg)
}
