package main

import (
	"fmt"
	"regexp"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func CompileMetod() {
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
		"=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description, idx))
	}
}
