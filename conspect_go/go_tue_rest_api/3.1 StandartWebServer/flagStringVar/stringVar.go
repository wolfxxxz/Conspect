package main

import (
	"flag"
	"fmt"
)

func StringVarFlag() {
	var variable string //переменная принимает значение
	var defaultVal = "8080"
	// Определение флагов с дефолтными значениями
	flag.StringVar(&variable, "path", defaultVal, "description")

	// Парсинг флагов командной строки
	flag.Parse()

	if variable == "env" {
		fmt.Println(".env")
	} else if variable == "toml" {
		fmt.Println("toml")
	} else {
		fmt.Println(variable)
	}
}

// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe -path=toml
// toml
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe -path=abracadabra
// abracadabra
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe
// 8080
