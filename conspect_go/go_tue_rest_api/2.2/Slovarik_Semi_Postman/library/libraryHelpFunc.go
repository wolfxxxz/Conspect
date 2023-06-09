package library

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Themes struct {
	Themes   string `json:"themes"`
	Quantity int    `json:"quantity"`
}

func NewThemes(newThemes string, newQuantity int) Themes {
	t := Themes{newThemes, newQuantity}
	return t
}

// Показать список Тема : количество
func QuantityThemes(ll []Library) []Themes {
	l := []Library{}
	l = append(l, ll...)
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Theme > l[j].Theme
	})
	SliseThemes := []Themes{}

	var quantity int
	var withoutTheme int
	for i, v := range l {
		if i == len(l)-1 {
			WithTheme := NewThemes("Withaut themes", withoutTheme)
			SliseThemes = append(SliseThemes, WithTheme)
			//fmt.Print("Слов без темы: ", withoutTheme, " \n")
			break
		} else if v.Theme == "" {
			withoutTheme++
		} else if v.Theme == l[i+1].Theme {
			quantity++
		} else if v.Theme != l[i+1].Theme {
			Theme := NewThemes(v.Theme, quantity)
			SliseThemes = append(SliseThemes, Theme)
			//fmt.Print(v.Theme, ": ", quantity, " || ")
			quantity = 0
		}
	}
	return SliseThemes
	//fmt.Println()
}

// Del []Library DUblicat -------------------------------------------------------------------------------
func DelDublikat(s []Library) []Library {
	c := []Library{}
	count := 0
	for ii, v := range s {
		for i := ii; i <= len(s)-1; i++ {
			if strings.EqualFold(v.English, s[i].English) {
				//if v.English == s[i].English {
				count++
			}
		}
		if count == 1 {
			c = append(c, v)
			count = 0
		} else {
			count = 0
		}
	}
	return c
}

// Количество слов
func CalculateWordsLibrary(s []Library) int {
	ii := len(s) - 1
	return ii
}

func Print(l []Library) {
	for _, v := range l {
		fmt.Print(v.English, " - ", v.Russian, " - ", v.Theme)
		fmt.Println()
	}
}

// Scan bufio
func Scan() string {
	in := bufio.NewScanner(os.Stdin)
	var nn string
	if in.Scan() {
		nn = in.Text()
	}
	return nn
}

// Исправить слова
func CorectingWords(i int, l []Library) {
	fmt.Println("Оставить старое значение Enter|| или введите новое значение")
	fmt.Println("write English", l[i].English)
	//var englischW string
	englischW := Scan()
	if englischW != "" {
		l[i].English = englischW
	}
	fmt.Println("write Russian", l[i].Russian)
	russianW := Scan()
	if russianW != "" {
		l[i].Russian = russianW
	}
	//var themeW string
	fmt.Println("write Theme", l[i].Theme)
	themeW := Scan()
	if themeW != "" {
		l[i].Theme = themeW
	}
}

// ----------New words from РУЧКАМИ-------------
func NewWordRukamy(l []Library) {
	c := len(l)
	ll := []Library{}
	//Сам механизм
	for {
		fmt.Println("Для выхода введите 1")
		fmt.Println("English")
		wordsEng := Scan()
		if wordsEng == "1" {
			break
		}

		fmt.Println("Russian")
		wordsRus := Scan()
		if wordsRus == "1" {
			break
		}

		fmt.Println("Theme")
		wordsTheme := Scan()
		if wordsTheme == "1" {
			break
		}
		var id = len(l) + 1
		d := NewLibrary(wordsEng, wordsRus, wordsTheme, id)
		ll = append(l, d)
	}

	ll = append(ll, l...)

	rttt := DelDublikat(ll)
	d := len(rttt)

	if c != d {
		fmt.Println("                   New Words Add:", d-c)
	}
	Savejson(rttt, "library.json")
	SaveTXT(rttt, "library.txt")
	/*} else {
	fmt.Println("ok, go next")*/
}

func AddTheme(ll []Library) {
	//---------------Необходимо вернуть в тойже последовательности
	//Слов без темы
	//Показать список Тема : количество
	l := []Library{}
	l = append(l, ll...)
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Theme > l[j].Theme
	})

	var quantity int
	var withoutTheme int
	for i, v := range l {
		if i == len(l)-1 {
			fmt.Print("Слов без темы: ", withoutTheme, " \n")
			break
		} else if v.Theme == "" {
			withoutTheme++
		} else if v.Theme == l[i+1].Theme {
			quantity++
		} else if v.Theme != l[i+1].Theme {
			fmt.Print(v.Theme, ": ", quantity, " || ")
			quantity = 0
		}
	}
	//fmt.Println()

	var d int
	for _, v := range l {
		if v.Theme == "" {
			d++
		}
	}
	//ищем слова без темы
	var c int
	for i, v := range ll {
		wordTheme := ""
		if v.Theme == "" {
			fmt.Println("Для выхода введите 1, пропустить слово нажмите Enter, редактировать все данные 9")
			fmt.Println(v.English, v.Russian)
			wordTheme = Scan()
		}
		if wordTheme == "9" {
			CorectingWords(i, ll)
			continue
		}
		if wordTheme == "" {
			continue
		} else if wordTheme != "1" {
			ll[i].Theme = wordTheme
			c++
		} else {
			break
		}
	}
	//
	fmt.Println("                   Изменено слов:", c)
	fmt.Println("                   Слов без темы:", d-c)
	fmt.Println("         всего слов в библиотеке:", len(ll)+1)

	Savejson(ll, "library.json")
	//Savejson(ll, "test.json")
	SaveTXT(ll, "library.txt")
}
