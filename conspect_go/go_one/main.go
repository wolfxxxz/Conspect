package main

import (
	"fmt"
	"strings"
)

func main() {

	a, b, c, d, e, f := "one", "one", "two", "two", "t", "t"
	var SlicePointString = []*string{&c, &a, &d, &e, &f, &b}

	printPointArr(SlicePointString)

	DelDublikat(&SlicePointString)

	printPointArr(SlicePointString)
}

func ReverseSlice(s []*string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func DelDublikat(s *[]*string) {

	printPointArr(*s)

	for ii, v := range *s {
		var count1 int
		for i := ii; i <= len(*s)-1; i++ {
			if strings.EqualFold(*v, *(*s)[i]) { //expected type, found ')'
				count1++
			}
		}
		if count1 == 1 {

			count1 = 0
		} else {
			count1 = 0
			slice := (*s)[1:]
			s = &slice

		}
	}
	fmt.Println("Without dublicat")

	printPointArr(*s)

}

func printPointArr(s []*string) {
	for _, v := range s {
		fmt.Print(*v, " ")
	}
	fmt.Println()
}
