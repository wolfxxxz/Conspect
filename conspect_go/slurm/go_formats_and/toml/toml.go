package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

/*api.toml
bind_addr = ":8081"
logger_level = "debug"
*/

type Addr struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"logger_level"`
}

func ReadToml() {
	adr := Addr{}
	config := &adr
	_, err := toml.DecodeFile("api.toml", config)
	if err != nil {
		log.Println("Can not find configs file. Using default value", err)
	}

	fmt.Println(config.BindAddr) //:8081
	fmt.Println(config.LogLevel) //debug
}
