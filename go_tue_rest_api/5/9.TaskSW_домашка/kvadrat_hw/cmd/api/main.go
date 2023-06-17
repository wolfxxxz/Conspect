package main

import (
	"kvadrat_hw/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const apiPrefix string = "/api/v1"

var port string

func init() {
	//Если есть файл .env то достать из него порт
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server and port: ", port)

	//Инициализация роутера
	router := mux.NewRouter()
	utils.BuildWordResource(router, apiPrefix)
	//utils.BuildManyWordsResource(router, manyWordsResourcePrefix)
	log.Println("Router initalizing successfully")
	//Безконечный цикл
	//log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
