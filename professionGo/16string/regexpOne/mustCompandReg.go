package main

import (
	"fmt"
	"regexp"
)

func MustCompileAndRegExp() {
	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
	description := "Kayak. A boat for one person."
	str := pattern.FindString(description)
	fmt.Println("Match:", str)
}
