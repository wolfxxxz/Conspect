package main

import (
	"flag"
	"log"

	"ServerAndDB2/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	//flag.Pars читает данные с настройками flag.Stringvar(...)
	flag.Parse()
	log.Println("It works")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find configs file. Using default value", err)
	}
	server := api.New(config)

	//it`s the same
	log.Fatal(server.Start())

}
