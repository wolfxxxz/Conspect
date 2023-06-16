package main

import (
	"strings"

	"github.com/agnivade/levenshtein"
)

func CompareStringsLevenshtein() bool {
	str1 := "приет"
	str2 := "приве"
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if distance := levenshtein.ComputeDistance(str1, str2); distance <= 2 {
		return true
	} else {
		return false
	}
}
