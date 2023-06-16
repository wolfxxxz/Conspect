package main

import (
	"fmt"
	"strings"
)

func Replacer() {
	text := "It was a boat. A small boat."
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge", "It", "He")
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced) //He was a kayak. A huge kayak.
}
