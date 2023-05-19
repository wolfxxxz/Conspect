# 11 Methods
## metod
type Product struct {
	name, category string
	price          float64
}

func (p *Product) printDetails() {
	fmt.Println("Name:", p.name, "Category:", p.category, "Price ", p.calcTax(0.2, 100))
}

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.price > threshold {
		return p.price + (p.price * rate)
	}
	return p.price
}

func main() {
	products := []*Product{
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	d := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	products = append(products, &d)

	for _, p := range products {
		p.printDetails()
	}
}
## metod нельзя декларировать повторно
## Вызов метода через тип получателя (анонимно)
type Product struct {
	name, category string
	price          float64
}

func (p Product) printDetails() {
	fmt.Println("Name:", p.name, "Category:", p.category, "Price ", p.calcTax(0.2, 100))
}

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.price > threshold {
		return p.price + (p.price * rate)
	}
	return p.price
}

func main() {
	Product.printDetails(Product{"Kayak", "Watersports", 275})
}
## Псевдоним типа и метод для него
type Product struct {
	name, category string
	price          float64
}

// Псевдоним типа
type ProductList []Product

// Метод для него
func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func main() {
	products := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total: ", total)
	}
}
//Category:  Watersports Total:  323.95
//Category:  Soccer Total:  19.5
## Преобразование типов
type Product struct {
	name, category string
	price          float64
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}
//------------------------
func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}
func main() {
	products := ProductList(getProducts())

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total: ", total)
	}
}
## Размещение типов и методов в отдельных файлах
//service.go

package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}
//-------------------
// product.go

package main

type Product struct {
	name, category string
	price          float64
}
//-----------------
//main.go

package main

import "fmt"

func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.5}
	fmt.Println("Product: ", kayak.name)
	fmt.Println("Insurance: ", insurance.description)
}
//cmd go run .
//Product:  Kayak
//Insurance:  Boat Cover
# Interface
## Create interface and methods under it
//main.go
package main

import "fmt"

type Expense interface {
	getName() string
	getCoast(annual bool) float64
}

func main() {
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.5},
	}
	for _, expense := range expenses {
		fmt.Println("Expence: ", expense.getName(), "Cost: ", expense.getCoast(true))
	}
}
//-------------
// product.go

package main

type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCoast(_ bool) float64 {
	return p.price
}
//------------
//service.go

package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCoast(_ bool) float64 {
	return s.monthlyFee * float64(s.durationMonths)
}
## Функция принимает интерфейс calcTotal(expenses []Expense) 
func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCoast(true)
	}
	return
}

//func main()
fmt.Println(calcTotal(expenses)) //1349
## Интерфейс как поле структуры
type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {

	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.5},
	}
	//expenses Интерфейс как поле структуры
	account := Account{
		accountNumber: 777,
		expenses:      expenses,
	}

	// Доступ как к полю структуры account.expenses
	for _, expense := range account.expenses {
		fmt.Println("Expence: ", expense.getName(), "Cost: ", expense.getCoast(true))
	}
	fmt.Println(calcTotal(expenses))
}
## эффект приемников и метод указателей
product := Product{"Kayak", "Watersports", 275}
    // Происходит копирование
***	var expense Expense = product**
	product.price = 100
	fmt.Println(product) //{Kayak Watersports 100}

	fmt.Println(expense) //{Kayak Watersports 275}

	//Доступ через продукт и к данным и к методам
	fmt.Println(product.name, product.category, product.price, product.getCoast(true), product.getName())
	//Через интерфейс только к методам
	fmt.Println(expense.getCoast(true), expense.getName())
------------------------------------------------
product := Product{"Kayak", "Watersports", 275}
    // Происходит копирование АДРЕССА
***	var expense Expense = &product**
	product.price = 100
	fmt.Println(product) //{Kayak Watersports 100}

	fmt.Println(expense) //{Kayak Watersports 100}
-----------------------------------------------
***Вывод**
//Использовать указатели при создании метода
func (p *Product) getName() string {
	return p.name
}

func (p *Product) getCoast(_ bool) float64 {
	return p.price
}
## Утверждение типа (сужение типа) s := expense.(Service)
### Тестирование типа
#### 1
func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
	}

	for _, expense := range expenses {
		// доступны только методы т.к expense - интерфейс
		fmt.Println(expense.getCoast(true), expense.getName())
		// утверждение типа даёт доступ ко всем возможностям типа
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
	}
}
#### Внимательно panic: interface conversion:
func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
***		&Product{"Kayak", "Watersports", 275},**
	}

	for _, expense := range expenses {
		// доступны только методы т.к expense - интерфейс
		fmt.Println(expense.getCoast(true), expense.getName())
		// утверждение типа даёт доступ ко всем возможностям типа
***		s := expense.(Service)//Product вам не Service**
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
	}
}
#### Test Если всё таки типы замешаны в интерфейсе 
func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCoast(true))
		}

	}
}
### Включение динамических типов Switch удобно сортирует значения по типам
for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
		default:
			fmt.Println("Expense", expense.getName(), expense.getCoast(true))
		}
	}
## Пустой интерфейс
### Что может принять пустой интерфейс - ВСЁ
***Как его разобрать*** - switch case
func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.5, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price, "Cathegory", value.category)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default", value)
		}
	}

}
### Функция принимает пустой интерфейс - 
func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price, "Cathegory", value.category)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default", value)
	}
}
func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.5, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		processItem(item)
	}
}
### Функция принимает []interface{}
func processItem(items []interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price, "Cathegory", value.category)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default", value)
		}
	}
}

func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.5, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	processItem(data)

}
#
