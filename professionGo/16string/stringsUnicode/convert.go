package main

import (
	"fmt"
	"strings"
)

func Convert() {
	str := "HeLLo"
	//возвращает новую строку, содержащую символы указанной строки, преобразованные в нижний регистр.
	fmt.Println("strings.ToLower ", str, "-> ", strings.ToLower(str))
	//преобразованные в нижний регистр.
	fmt.Println("strings.ToLoUpper ", str, "-> ", strings.ToUpper(str))
	str1 := "hello i am john"
	//первый символ каждого слова был в верхнем регистре, а остальные символы — в нижнем
	fmt.Println("strings.Title ", str1, "-> ", strings.Title(str1))
	//все в верхний регистр
	fmt.Println("strings.ToTitle ", str1, "-> ", strings.ToTitle(str1))

}
