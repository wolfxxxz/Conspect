package main

import (
	"flag"
	"fmt"
)

func main() {
	// Определение флагов
	aPtr := flag.Bool("a", false, "Флаг 'a' для вывода 'привет'")
	bPtr := flag.Bool("b", false, "Флаг 'b' для вывода 'пока'")

	// Парсинг флагов командной строки
	flag.Parse()

	// Вывод значений флагов
	if *aPtr {
		fmt.Println("привет")
	}
	if *bPtr {
		fmt.Println("пока")
	}
}

//go run flagg/main.go -a
/*
#### flag string
func main() {
	var defaultVal = "Привет"
	// Определение флагов с дефолтными значениями
	aPtr := flag.String("a", defaultVal, "Значение флага 'a'")

	// Парсинг флагов командной строки
	flag.Parse()

	// Вывод значений флагов
	fmt.Println("Значение флага 'a':", *aPtr)

}
//makefile
build:
	go build
modinit:
	go mod init flag2
goFlag2:
	./flag2.exe
goFlag2PlusString:
	./flag2.exe -a=hello
	
#### flag StringVar
func main() {
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
*/
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe -path=toml
// toml
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe -path=abracadabra
// abracadabra
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe
// 8080
