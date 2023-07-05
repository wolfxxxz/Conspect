package main

import (
	"fmt"
)

func main() {
	fmt.Println("formats")
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
