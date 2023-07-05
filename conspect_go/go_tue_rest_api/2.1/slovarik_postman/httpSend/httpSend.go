package httpSend

import (
	"encoding/json"
	"log"
	"net/http"
	"slovarik_postman/library"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

var l []library.Library

func init() {
	l = library.Takejson("library.json")
}

func ThemesWordsPrint(writer http.ResponseWriter, request *http.Request) {
	//Пишем хедеры // установить тип данных
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Хтось питає за теми слів")
	writer.WriteHeader(200) // Статус код для запроса
	lThemes := library.QuantityThemes(l)
	json.NewEncoder(writer).Encode(lThemes) // Сериализация и запись в writer (запись в браузер)
	//2
}

// Вспомогательная функция т.к роутер только печатает и не принимает левых аргументов
func FindWordById(id int) (library.Library, bool) {
	var word library.Library
	var found bool
	for _, p := range l {
		if p.ID == id {
			word = p
			found = true
			break
		}
	}
	return word, found
}

// http.ResponseWriter - куда записать и  *http.Request что записать
func GetWordById(writer http.ResponseWriter, request *http.Request) {
	//Пишем хедеры // установить тип данных
	writer.Header().Set("Content-Type", "application/json")
	// Как получить ID Запроса
	vars := mux.Vars(request) // возвращает мап {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Client trying to use invalid id param", err)
		msg := ErrorMessage{Message: "do not use id not supported int casting"}
		writer.WriteHeader(400) //bed request
		json.NewEncoder(writer).Encode(msg)
		return
	}

	log.Println("Хтось питає за слово під номером: ", id)
	word, ok := FindWordById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(word)
	} else {
		msg := ErrorMessage{Message: "word with that id does'nt exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}

	writer.WriteHeader(200) // Статус код для запроса
	//h := FindWordById()
	//json.NewEncoder(writer).Encode()
}
