package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"StandartWebServer/internal/app/api"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

var variable string //переменная принимает значение
var defaultVal = "8080"
var description = "path to config file in toml or .env format"
var path = "configs/"

func init() {
	flag.StringVar(&variable, "path", defaultVal, description)

}

func main() {
	//flag.Pars читает данные с настройками flag.Stringvar(...)
	flag.Parse()
	log.Println("It works")
	//fmt.Println(configPath)

	config := api.NewConfig()

	//И сохранить в config
	if variable == "env" {
		err := godotenv.Load(path + ".env")
		if err != nil {
			log.Fatal("Could not find .env file:", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")

	} else if variable == "toml" {
		_, err := toml.DecodeFile(path+"api.toml", config)
		if err != nil {
			log.Println("Can not find configs file. Using default value", err)
		}
	} else {
		fmt.Println("you wrote wrong file - config ", variable)
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
