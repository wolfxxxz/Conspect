package main

import (
	"fmt"
	"strings"
)

func Strings() {
	product := "Kayak"
	//Эта функция возвращает true, если строка s содержит substr, и false, если нет.
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	//Эта функция возвращает true, если строка s содержит любой из символов,
	//содержащихся в строке substr.
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	//Эта функция возвращает true, если строка s содержит определенную руну (rune).
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	//Эта функция выполняет сравнение без учета регистра и возвращает
	//true, если строки s1 и s2 совпадают.
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	//Эта функция возвращает значение true, если строка s начинается с префикса (prefix) строки.
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	//Эта функция возвращает значение true, если строка заканчивается суффиксом (suffix) строки.
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
}
