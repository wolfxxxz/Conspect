package main

import (
	"flag"
	"log"

	"Slovarik/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string = "configs/api.toml"
)

func init() {

	flag.StringVar(&configPath, "pathtoml", "configs/api.toml", "path to config file in .toml format")

}

func main() {
	//Запускаем функцию init и flag.StringVar()
	flag.Parse()
	log.Println("It works")
	//Server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find configs file. Using default value", err)
	}
	//Теперь нужно попробовать прочитать из .toml/.env , так как там может быть новая информация
	server := api.New(config)

	//api server Start
	log.Fatal(server.Start())

}
