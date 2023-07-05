package main

import (
	"Semitrash_API/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrefix + "/book"  //api/v1/book
	manyBooksResourcePrefix string = apiPrefix + "/books" //api/v1/books
)

// 1
func init() {
	//Достаёт порт из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()
	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResourcePrefix(router, manyBooksResourcePrefix)
	log.Println("Router initializing successfully, Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
