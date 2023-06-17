package testink

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"Slovarik/library"
)

// Сравнение слов
func Work(s []library.Library) (yes int, not int, wordsIncorrect []library.Library, wordsRight []library.Library) {
	fmt.Println("                     START")
	for _, v := range s {
		y, n := Compare(v)
		if y > 0 {
			yes++
			wordsRight = append(wordsRight, v)
		} else if n > 0 {
			not++
			wordsIncorrect = append(wordsIncorrect, v)
		}
	}
	return yes, not, wordsIncorrect, wordsRight
}

// repair mistakes
func WorkMistake(s []library.Library) {
	fmt.Println("                 Work with mistakes")
	//fmt.Println("--- Check and repair ---")
	for {
		if len(s) == 0 {
			break
		}
		v := s[len(s)-1]
		y, _ := Compare(v)

		if y > 0 && len(s) != 1 {
			s = s[:len(s)-1]
			fmt.Println(len(s))
		} else if len(s) == 1 && y > 0 {
			s = []library.Library{}
		}
	}
}

// Тест по количеству
func TestKnowlig(l []library.Library) {
	//fmt.Println("Flag-----------------------------------------------")
	log.Println("                       Start")

	//Scan quantity words for test
	//----------------------------------------------
	fmt.Println("Количество слов для теста")
	var quantity int
	lenLibrary := len(l)
	for {
		cc := Scan()
		i, err := strconv.Atoi(cc)
		if err != nil {
			fmt.Println("Incorect, please enter number")
		} else if i >= lenLibrary {
			fmt.Printf("Incorect, please enter less number. Len Library is: %v\n", lenLibrary)
		} else {
			quantity = i
			break
		}
	}

	//Test_1
	// Cute some
	TestWords := l[:quantity]
	s, e, incorectWords, rightWords := Work(TestWords)
	if len(incorectWords) >= 1 {
		library.Print(incorectWords)
		//Test_2
		WorkMistake(incorectWords)
		fmt.Println(s, e)
	} else {
		fmt.Println("    БЕЗ ОШИБОК !!!")
		fmt.Printf(" Right answers is: %v\n", s)
	}
	//write used words in the end
	rttt := l[quantity:]
	rttt = append(rttt, rightWords...)
	incorectWords = append(incorectWords, rttt...)
	//Сохранить в txt file
	library.SaveTXT(incorectWords, "library.txt")
	//Сохранить в json file
	library.Savejson(incorectWords, "library.json")

	fmt.Println("  All the words in a dictionary: ", len(incorectWords)+1)
	log.Println("Final")
}

// Тест по темам
func ThemesOfWords(l []library.Library) {
	//Упорядочить по теме
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Theme > l[j].Theme
	})
	//library.Savejson(l, "library.json")
	library.SaveTXT(l, "library.txt")

	var quantity int = 1
	for i, v := range l {
		if i == len(l)-1 {
			fmt.Print("Осталось заполнить: ", quantity)
			break
		} else if v.Theme == l[i+1].Theme {
			quantity++
		} else if v.Theme != l[i+1].Theme {
			fmt.Print(v.Theme, ": ", quantity, " ")
			quantity = 1
		}
	}
	fmt.Println()
	fmt.Println("Для теста введите название темы")
	themes := Scan()
	ThemeSlice := []library.Library{}
	if themes == "" {
		fmt.Println("You are lazy")
	} else {
		for _, v := range l {
			if v.Theme == themes {
				ThemeSlice = append(ThemeSlice, v)
			}
		}

		s, e, incorectWords, _ := Work(ThemeSlice)
		if len(incorectWords) >= 1 {

			library.Print(incorectWords)
			//Test_2
			WorkMistake(incorectWords)
			fmt.Println(s, e)
		} else {
			fmt.Println("    БЕЗ ОШИБОК !!!")
			fmt.Printf(" Right answers is: %v\n", s)
		}
	}
}
