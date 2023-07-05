package library

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var L []Library

func init() {
	//-----Зачитываем содержимое файла библиотека-----
	L = Takejson("library.json")

}

type Slovarick struct {
	Words []Library
}

type Library struct {
	English string `json:"english"`
	Russian string `json:"russian"`
	Theme   string `json:"theme"`
	ID      int    `json:"id"`
}

// Create Library
func NewLibrary(newEnglish string, newRussian string, newTheme string, newId int) Library {
	i := Library{newEnglish, newRussian, newTheme, newId}
	return i
}

// Сохранить в json file
// Marshal
func Savejson(l []Library, file string) {
	s := Slovarick{
		Words: l,
	}
	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(byteArr))             0664
	err = os.WriteFile(file, byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
}

// Достать с json file
// Unmarshal
func Takejson(file string) []Library {
	//1. Создадим файл дескриптор
	filejson, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer filejson.Close()
	//fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go

	f := make([]byte, 64)
	var data2 string

	for {
		n, err := filejson.Read(f)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data2 = data2 + string(f[:n])
	}

	data := []byte(data2)
	var users Slovarick

	// Теперь задача - перенести все из data в users - это и есть десериализация!
	json.Unmarshal(data, &users)

	Usersu := users.Words

	return Usersu
}

// Вспомогательная функция т.к роутер только печатает и не принимает левых аргументов
func FindWordById(id int) (Library, bool) {
	var word Library
	var found bool
	for _, p := range L {
		if p.ID == id {
			word = p
			found = true
			break
		}
	}
	return word, found
}

// нужно заменить oldBook на newBook v DB
func ReWriteWord(id int, word Library) {
	word.ID = id
	for i, b := range L {
		if b.ID == id {
			L[i] = word
		}
	}
}

// Удалить book from DB
func DelWordFromL(f int) {
	for i, _ := range L {
		if i != len(L)-1 && i >= f-1 {
			L[i] = L[i+1]
		}
		if i == len(L)-1 {
			L = L[:i]
		}
	}
}
