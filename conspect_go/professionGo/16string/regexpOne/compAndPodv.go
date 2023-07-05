package main

import (
	"fmt"
	"regexp"
)

func CompAndPodv() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	subs := pattern.FindStringSubmatch(description)
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}
}
