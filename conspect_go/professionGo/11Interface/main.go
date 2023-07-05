// main.go
package main

import "fmt"

type Expense interface {
	getName() string
	getCoast(annual bool) float64
}
type Person struct {
	name, city string
}

/*
func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCoast(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses      []Expense
}
*/
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

/*
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

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
}

/*
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = &product
	product.price = 100
	fmt.Println(product) //{Kayak Watersports 100}

	fmt.Println(expense) //{Kayak Watersports 275}

	//Доступ через продукт и к данным и к методам
	fmt.Println(product.name, product.category, product.price, product.getCoast(true), product.getName())
	//Через интерфейс только к методам
	fmt.Println(expense.getCoast(true), expense.getName())
}

/*
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

/*
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

//--------------------------

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
	products := ProductList(getProducts())

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total: ", total)
	}

	/*
		Product.printDetails(Product{"Kayak", "Watersports", 275})
		/*
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
*/
