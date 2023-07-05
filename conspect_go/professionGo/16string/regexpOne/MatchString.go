package main

import (
	"fmt"
	"regexp"
)

func MatchString() {
	description := "A boat for one person"
	// Ищем oat в тексте
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match) //true
	} else {
		fmt.Println("Error:", err)
	}
}
