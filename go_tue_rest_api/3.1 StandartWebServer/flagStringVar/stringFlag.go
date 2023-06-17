package main

import (
	"flag"
	"fmt"
)

func StringFlag() {
	var defaultVal = "Привет"
	// Определение флагов с дефолтными значениями
	aPtr := flag.String("a", defaultVal, "Значение флага 'a'")

	// Парсинг флагов командной строки
	flag.Parse()

	// Вывод значений флагов
	fmt.Println("Значение флага 'a':", *aPtr)

}

/*
//makefile
goFlag2PlusString:
	./flag2.exe -a=hello*/
