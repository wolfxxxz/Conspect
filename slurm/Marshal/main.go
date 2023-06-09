package main

import (
	"Marshal/formats"
	"fmt"
)

func main() {
	fmt.Println("formats")
	//formats.ParseJson()
	//formats.CreateJson()
	//formats.LoadAndParseJson()
	//formats.LoadAndParseRawMsgToMap()
	//formats.LoadAndParseRawMsg()
	if per, ok := formats.TakeJson("exampls/person.json"); ok {
		formats.WriteJson("exampls/newPerson.json", per)
		fmt.Println(per)
	}
}
