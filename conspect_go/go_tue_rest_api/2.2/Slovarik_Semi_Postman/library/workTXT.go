package library

import (
	"fmt"
	"io"
	"os"
)

func SaveTXT(s []Library, files string) {
	file, err := os.Create(files)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for _, v := range s {
		file.WriteString(v.English)
		file.WriteString(" - ")
		file.WriteString(v.Russian)
		file.WriteString(" - ")
		file.WriteString(v.Theme)
		file.WriteString("\n")
	}
	//fmt.Println("Done.")
}

// Часть 1 Добавление новых слов в библиотеку Загрузкой с файла txt
func TakeTXT(filetxt string) []Library {
	//fmt.Println("Start")
	file, err := os.Open(filetxt)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	f := make([]byte, 100024) //Длинна строки
	data2 := []byte{}
	for {
		n, err := file.Read(f)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data2 = f[:n]
	}
	//fmt.Println("len data2", len(data2))
	//fmt.Println("range data2")
	sliseString := []string{}
	dbyte := []byte{}
	for i, v := range data2 {
		b := "-"
		//лишние пробелы
		if i < len(data2)-1 {
			if v == 32 && data2[i+1] == 32 {
				continue
			}
		}
		if v == 13 {
			continue
		}

		if v == 10 {
			d := string(dbyte) + b
			sliseString = append(sliseString, d)
			dbyte = []byte{}
		}
		if v == 10 {
			continue
		}
		dbyte = append(dbyte, v)
	}
	//fmt.Println(sliseString)

	SliceLib := []Library{}

	for _, vv := range sliseString {
		SliceThreeString := []string{}
		var Str string
		for _, v := range vv {
			if v == '-' && Str != "" {
				strByte := []byte(Str)
				//Проверка на лишние пробелы
				strByte2 := []byte{}
				for i, v := range strByte {
					if i == 0 && v == 32 {
						continue
					} else if i == len(strByte)-1 && v == 32 {
						continue
					} else {
						strByte2 = append(strByte2, v)
					}
				}
				Str = string(strByte2)
				SliceThreeString = append(SliceThreeString, Str)
				Str = ""
			}
			if v == '-' {
				continue
			}
			Str = Str + string(v)
		}
		if len(SliceThreeString) > 3 {
			SliceThreeString = SliceThreeString[:2]
		}
		for i := 0; len(SliceThreeString) == 2; i++ {
			if len(SliceThreeString) <= 2 {
				SliceThreeString = append(SliceThreeString, "")
			}
		}
		id := len(SliceLib) + 1
		a := NewLibrary(SliceThreeString[0], SliceThreeString[1], SliceThreeString[2], id)
		SliceLib = append(SliceLib, a)
	}
	return SliceLib
	//10 - начало строки
	//13 - enter
	//46 - точка
	//32 - пробел
	//45 - дефис
}

// Часть 2 Добавление новых слов в библиотеку Загрузкой с файла txt
func UpdateLibrary(filetxt string, l []Library) {
	SliceLib := TakeTXT(filetxt)
	c := len(l)
	//--------Соединяем два среза в один--------------
	//rttr = append(rttr, newWords...)
	SliceLib = append(SliceLib, l...)
	//------------Удаляем дубликаты-------------------
	rttt := DelDublikat(SliceLib)
	d := len(rttt)

	r := []Library{}
	SaveTXT(r, filetxt)

	if d != c {
		fmt.Println("                   New Words Add:", d-c)
		Savejson(rttt, "library.json")
		SaveTXT(rttt, "library.txt")
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить в файл 'newWords.txt'")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}
