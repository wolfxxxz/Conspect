package main

import (
	"log"
	"net/http"
	"slovarik_postman/httpSend"

	"github.com/gorilla/mux"
)

var (
	port string = "8081"
)

func main() {
	log.Println("Trying to start Slovarik")

	router := mux.NewRouter()
	//mux.NewRouter == http://localhost - вероятно используется по умолчанию...
	//Если на вход пришел запрос /themes
	router.HandleFunc("/themes", httpSend.ThemesWordsPrint).Methods("GET") //http://localhost:8081/themes

	//Если на вход пришел запрос  /themes/{id}
	router.HandleFunc("/words/{id}", httpSend.GetWordById).Methods("GET") //http://localhost:8081/words/5

	//Безконечный цикл
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
