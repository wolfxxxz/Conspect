package main

import (
	"fmt"
	"io"
	"os"
)

type Word struct {
	russian string
	english string
	tema    string
}

func NewWord(newEnglish string, newRussian string, newTheme string) Word {
	return Word{newRussian, newEnglish, newTheme}
}

func TakeTXT(filetxt string) []*Word {
	//fmt.Println("Start Take txt") //
	file, err := os.Open(filetxt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
	//fmt.Println("range data2") //
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
			if d != "-" {
				sliseString = append(sliseString, d)
				dbyte = []byte{}
			}
		}
		if v == 10 {
			continue
		}
		dbyte = append(dbyte, v)
	}
	//fmt.Println("workTXT", sliseString) //--------------------------------------------------

	SliceLib := []*Word{}

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

		a := NewWord(SliceThreeString[0], SliceThreeString[1], SliceThreeString[2])
		SliceLib = append(SliceLib, &a)
	}

	return SliceLib
	//10 - начало строки
	//13 - enter
	//46 - точка
	//32 - пробел
	//45 - дефис
}
