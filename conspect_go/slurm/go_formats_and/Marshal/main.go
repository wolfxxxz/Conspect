package main

import (
	"fmt"
)

func main() {
	fmt.Println("formats")
	//formats.ParseJson()
	//formats.CreateJson()
	//formats.LoadAndParseJson()
	//formats.LoadAndParseRawMsgToMap()
	//formats.LoadAndParseRawMsg()
	/*
		if per, ok := formats.TakeJson("exampls/person.json"); ok {
			formats.WriteJson("exampls/newPerson.json", per)
			fmt.Println(per)
		}*/
	//fmt.Println(CreateXml())
	//DecodeXml(CreateXml())
	//NestedXml()
	//WriteXml("exampl.xml")
	//TakeXml("exampl.xml")
	//GetConf("exampl.yml")
	//WriteYaml("exampl2.yml")
	config := &BuildConf{}
	config.ReadBigConf("anchor_example.yml")
	config.WriteBigConf("anchor_example2.yml")

}
