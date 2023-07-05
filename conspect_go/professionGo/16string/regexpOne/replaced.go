package main

import (
	"fmt"
	"regexp"
)

func Replace() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)
}
