package formats

import (
	"encoding/json"
	"fmt"
)

// https://jsonformatter.curiousconcept.com/
type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func ParseJson() {
	birdJson := `{"species":"pigeon","description":"likes to perch onrocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	err := json.Unmarshal([]byte(birdJson), &bird)

	if err != nil {
		panic(err)
	}
	fmt.Println(bird)
}

func CreateJson() {
	bird := Bird{
		Species:     "Eagle",
		Description: "Cool eagle",
		Dimensions: Dimensions{
			Height: 100,
			Width:  50,
		},
	}
	//data, _ := json.Marshal(bird) //форматировать в строчку
	data, _ := json.MarshalIndent(bird, "", "    ") //форматирует читабельно
	fmt.Println(string(data))
}
