package main

import "fmt"

func ScanArr() {
	var in int
	fmt.Print("Количество елементов среза: ")
	fmt.Scan(&in)

	vals := make([]string, in)
	ivals := make([]interface{}, in)
	for i := 0; i < len(vals); i++ {
		ivals[i] = &vals[i]
	}
	fmt.Print("Enter text to scan: ")
	fmt.Scan(ivals...)
	Printfln("Name: %v", vals)
}
