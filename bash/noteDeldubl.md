package main

import (
	"fmt"
	"strings"
)

func main() {

	a, b, c, d, e, f := "one", "one", "two", "two", "t", "t"
	var SlicePointString = []*string{&c, &a, &d, &e, &f, &b}

	printPointArr(SlicePointString)

	result := DelDublikat(SlicePointString)

	printPointArr(result)
}

func ReverseSlice(s []*string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func DelDublikat(s []*string) []*string {
	/*
		// Зачем это?
		var count = 0
		for i := 0; i <= len(s)-1; i++ {
			for _, v := range s {
				if strings.EqualFold(*v, *s[i]) {
					count++
				}
				if count == 2 {
					s[i] = v
					count = 0
				}
			}
		}
		fmt.Println("Зачем это")
		printPointArr(s)*/
	fmt.Println("reverse")
	ReverseSlice(s)
	printPointArr(s)
	withoutDublicat := []*string{}
	for ii, v := range s {
		var count1 int
		for i := ii; i <= len(s)-1; i++ {
			if strings.EqualFold(*v, *s[i]) {
				count1++
			}
		}
		if count1 == 1 {

			withoutDublicat = append(withoutDublicat, v)
			count1 = 0
		} else {
			count1 = 0
			s = s[1:]
		}
	}
	fmt.Println("Without dublicat")
	printPointArr(withoutDublicat)
	ReverseSlice(withoutDublicat)
	printPointArr(s)
	return withoutDublicat
}

func printPointArr(s []*string) {
	for _, v := range s {
		fmt.Print(*v, " ")
	}
	fmt.Println()
}
