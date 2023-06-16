package main

import (
	"fmt"
	"unicode"
)

func Unicode() {
	str := "Ri"
	//Эта функция возвращает true, если указанная руна в нижнем регистре.
	fmt.Println(string(str[0]), unicode.IsLower([]rune(str)[0]), string(str[1]), unicode.IsLower([]rune(str)[1]))
	// Эта функция возвращает строчную руну, связанную с указанной руной.
	fmt.Println(string(str[0]), unicode.ToLower([]rune(str)[0]))
	//Эта функция возвращает значение true, если указанная руна написана в верхнем регистре.
	fmt.Println(string(str[0]), unicode.IsUpper([]rune(str)[0]), string(str[1]), unicode.IsUpper([]rune(str)[1]))

}
