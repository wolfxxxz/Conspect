package main

import "fmt"

var i int = 10

func incrementor(in *int) func() int {
	//i := *in
	return func() int {
		i := *in
		i++
		return i
	}
}

func main() {
	i = 0
	inc := incrementor(&i)
	fmt.Println(inc()) // Вывод: 1
	fmt.Println(inc()) // Вывод: 1
	i = 1
	inc2 := incrementor(&i)
	fmt.Println(inc2()) // Вывод: 2
	fmt.Println(inc2()) // Вывод: 2
}
