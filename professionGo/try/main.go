package main

import "fmt"

func main() {

	// Инициализируем мапу, где ключ - это название категории, а значение - это список строк
	productsByCategory := make(map[string][]string)

	// Добавляем продукты в мапу
	productsByCategory["fruits"] = []string{"apple", "banana", "orange"}
	productsByCategory["vegetables"] = []string{"carrot", "potato", "cucumber"}
	productsByCategory["meat"] = []string{"beef", "pork", "chicken"}

	// Обращаемся к элементам мапы по ключу и работаем со срезом
	fmt.Println(productsByCategory["fruits"][0])       // Выведет "apple"
	fmt.Println(len(productsByCategory["vegetables"])) // Выведет 3

	// Изменяем срез в мапе
	productsByCategory["meat"] = append(productsByCategory["meat"], "lamb")
	fmt.Println(productsByCategory["meat"]) // Выведет [beef pork chicken lamb]
}
