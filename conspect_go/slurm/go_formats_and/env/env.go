package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*.env
app_port = 8080
path = get/go/next
*/

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file:", err)
	}
	port := os.Getenv("app_port") //8080
	path := os.Getenv("path")
	fmt.Println(port)
	fmt.Println(path)
}
