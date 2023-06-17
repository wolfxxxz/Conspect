package utils

import (
	"slovarik/handlers"

	"github.com/gorilla/mux"
)

func BuildWordResource(router *mux.Router, prefix string) {
	//Достать слово по ид
	router.HandleFunc(prefix+"/{id}", handlers.GetWordById).Methods("GET")       //api/v1/word/id
	router.HandleFunc(prefix, handlers.CreateWord).Methods("POST")               //api/v1/word
	router.HandleFunc(prefix+"/{id}", handlers.UpdateWordById).Methods("PUT")    //api/v1/word/id
	router.HandleFunc(prefix+"/{id}", handlers.DeleteWordById).Methods("DELETE") //api/v1/word/id
}

func BuildManyWordsResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllWords).Methods("GET")                //api/v1/words
	router.HandleFunc(prefix+"/themes", handlers.ThemesWordsPrint).Methods("GET") ////api/v1/words/themes
}
