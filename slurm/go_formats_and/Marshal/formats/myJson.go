package formats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name     string `json:"pogonjalo"`
	LastName string `json:"lastName"`
	Age      int    `json:"years"`
}

func TakeJson(doc string) ([]Person, bool) {
	var person []Person
	jsonData, err := os.ReadFile(doc)
	if err != nil {
		fmt.Println(err)
		return person, false
	}
	json.Unmarshal(jsonData, &person)
	return person, true
}

// Устаревший подход
func TakeJsonOpen(doc string) ([]Person, bool) {
	jsonFile, err := os.Open(doc)
	var person []Person
	if err != nil {
		fmt.Println(err)
		return person, false
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return person, false
	}
	json.Unmarshal(byteValue, &person)
	return person, true
}

// Marshaling

func WriteJson(doc string, db []Person) {
	byteArr, _ := json.MarshalIndent(db, "", "    ")
	err := os.WriteFile(doc, byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
}
