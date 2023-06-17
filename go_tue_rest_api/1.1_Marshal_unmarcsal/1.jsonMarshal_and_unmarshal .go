package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Professors struct {
	Professors []Professor `json:"Professors"`
}

type Professor struct {
	Name       string     `json:"name"`
	ScienceID  int        `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Bob",
		ScienceID: 81263123,
		IsWorking: true,
		University: University{
			Name: "BMSTU",
			City: "pidar_from_moscow",
		},
	}

	prof2 := Professor{
		Name:      "Abi",
		ScienceID: 101263123,
		IsWorking: true,
		University: University{
			Name: "Loxik",
			City: "Pidar_from_piter",
		},
	}

	prof := []Professor{prof1, prof2}

	db := Professors{
		Professors: prof,
	}

	//------------------------------------------------------
	//   Запихнуть в файл
	//         Marshal

	//1. Превратим профессоровов в последовательность байтов
	// json.MarshalIndent(db, "", " ") "" префикс, "    " 4 пробела
	byteArr, err := json.MarshalIndent(db, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(byteArr))             0664
	err = os.WriteFile("output.json", byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}

	//----------------------------------------------------------------
	//     Вытащить из файла
	//         Unmarshal

	//1. Создадим файл дескриптор
	file, err := os.Open("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go
	// Инициализируем экземпляр Users

	f := make([]byte, 64)
	var data2 string

	for {
		n, err := file.Read(f)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data2 = data2 + string(f[:n])
	}

	data := []byte(data2)
	var users Professors

	// Теперь задача - перенести все из byteValue в users - это и есть десериализация!
	json.Unmarshal(data, &users)

	Usersu := users.Professors

	fmt.Println("SU", Usersu)
}
