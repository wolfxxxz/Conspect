package main

import (
	"fmt"
	"sort"
)

func main() {
	names := [3]string{"Alise", "Charlie", "Bob"}
	firstName := &names[1]
	secondName := names[1]
	fmt.Println(*firstName) // Charlie
	fmt.Println(secondName) // Charlie
	sort.Strings(names[:])
	fmt.Println(*firstName) // Bob // Теперь во второй ячейке имя Bob
	fmt.Println(secondName) // Charlie
}
