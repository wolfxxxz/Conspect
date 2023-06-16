package main

import (
	"fmt"
	"strings"
)

func Manipulate() {
	description := "  A  boat  for  "
	// Разбивает на строки при этом чистит пробелы и пустые строки
	field := strings.Fields(description)
	for i, x := range field {
		field[i] = ">" + x + "<"
	}
	fmt.Println("strings.Fields", field) //[>A< >boat< >for<]

	//Разбивает на строки чистит пробелы
	// - оставляет пустые строки
	splits := strings.Split(description, " ")
	for i, x := range splits {
		splits[i] = ">" + x + "<"
	}
	fmt.Println("strings.Split", splits) //[>< >< >A< >< >boat< >< >for< >< ><]

	//Делит слова но ...
	//Оставляет все пробелы
	splitsAfter := strings.SplitAfter(description, " ")
	for i, x := range splitsAfter {
		splitsAfter[i] = ">" + x + "<"
	}
	fmt.Println("strings.SplitAfter", splitsAfter) //[> < > < >A < > < >boat < > < >for < > < ><]

	//Чистит пробелы перед и после
	//Но не внутри
	trimSpase := strings.TrimSpace(description)
	trimSpase = ">" + trimSpase + "<"
	fmt.Println("strings.TrimSpace", trimSpase) //>A  boat  for<
}

func DoubleSpase() {
	description := "This  is  double  spaced"
	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
}
