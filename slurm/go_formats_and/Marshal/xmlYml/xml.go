package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func CreateXml() string {
	coffee := &Plant{Id: 27, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, "", "   ")
	return string(out)
}

func DecodeXml(input string) {
	var p Plant
	if err := xml.Unmarshal([]byte(input), &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
}

// Хз что она делает но без неё некак
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

type Plant struct {
	//Без поля xml.Name - ничего не будет
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}
type Nesting struct {
	XMLName xml.Name `xml:"nesting"`
	//Пропускаем parent>child> и записать в Plants = plant
	Plants []*Plant `xml:"parent>child>plant"`
}

func WriteXml(file string) {
	tomato := &Plant{Id: 01, Name: "tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ := xml.MarshalIndent(nesting, "", "   ")
	err := os.WriteFile(file, out, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("write out in %v\n", file)
}

func TakeXml(file string) {
	res := &Nesting{}
	xmlData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	xml.Unmarshal(xmlData, res)
	fmt.Println(res)
}
