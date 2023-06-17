package main

import (
	"flag"
	"fmt"
)

func Bolean() {
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
