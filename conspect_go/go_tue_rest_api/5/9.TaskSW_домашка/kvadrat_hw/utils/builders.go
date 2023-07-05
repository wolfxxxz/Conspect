package utils

import (
	"kvadrat_hw/handlers"

	"github.com/gorilla/mux"
)

func BuildWordResource(router *mux.Router, prefix string) {
	//Достать слово по ид
	router.HandleFunc(prefix+"/solve", handlers.GetKvadrat).Methods("GET")  //api/v1/solve
	router.HandleFunc(prefix+"/grab", handlers.PostKvadrat).Methods("POST") //api/v1/grab
}
