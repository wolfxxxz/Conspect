package main

import (
	"fmt"
)

//oneSlice := make([]string, 2)

func main() {
	var wordString = "Бибизянка Dusja"

	var wordRune []rune = []rune(wordString)

	fmt.Println(string(wordRune[2:14])) //бизянка Dusj
}
