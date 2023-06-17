package library

import (
	"fmt"
	"sort"
	"strconv"
)

func SortLibraryTheme(rttr []Library, file string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].Theme > rttr[j].Theme
	})
	SaveTXT(rttr, file)
}

func SortLibraryEnglisch(rttr []Library, file string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].Theme > rttr[j].Theme
	})
	SaveTXT(rttr, file)
}

func SortLibraryRussian(rttr []Library, file string) {
	sort.SliceStable(rttr, func(i, j int) bool {
		return rttr[i].Theme > rttr[j].Theme
	})
	SaveTXT(rttr, file)
}

func SortLibrary(l []Library, file string) {
	fmt.Println("sort Theme 1 || sort Englisch 2 || sort Russian 3")
	c := Scan()
	cc, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println("Incorect, please enter number")
	}
	if cc == 1 {
		SortLibraryTheme(l, file)
	} else if cc == 2 {
		SortLibraryEnglisch(l, file)
	} else if cc == 3 {
		SortLibraryRussian(l, file)
	} else {
		fmt.Println("You are lazy")
	}
}
