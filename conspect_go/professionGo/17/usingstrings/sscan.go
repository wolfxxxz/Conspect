package main

import "fmt"

func SSscan() {
	source := "Lifejacket Watersports 48.95"
	var name, category string
	var price float64
	//Сканирует в переменные со строки
	//разделитель пробел
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		fmt.Printf("Scanned %v values\n", n)
		fmt.Printf("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		fmt.Printf("Error: %v", err.Error())
	}
}

func SSscan2() {
	source := "Lifejacket Watersports 48.95"
	var name, category string
	var price float64
	//Сканирует в переменные со строки
	//разделитель пробел
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)
	if err == nil {
		fmt.Printf("Scanned %v values\n", n)
		fmt.Printf("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		fmt.Printf("Error: %v", err.Error())
	}
}
