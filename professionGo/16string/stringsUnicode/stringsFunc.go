package main

import (
	"fmt"
	"strings"
)

func StringsFieldsFunc() {
	description := "This  is  double  spaced"
	splitter := func(r rune) bool {
		return r == ' '
	}
	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}
}
