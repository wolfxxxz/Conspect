package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {

		PrintNumber(i)

	}

}

// revive:disable:exported
func PrintHello() {

	fmt.Println("hello")
}

// revive:enable:exported
// PrintNumber вот такая вот фича с этими коментами
func PrintNumber(i int) {
	fmt.Println(i)
}
