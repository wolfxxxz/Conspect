# Packet
## Создание пакетов, описание элементов 
// Комментировать можно и название пакета
// store/product.go
package store

//Продукт описывает элемент для продажи
type Product struct {
	name, category string  //Имя и тип продукта
	price          float64 //Стоимость
}
## Элементы Инкапсулирование 
// main.go
package main

import (
	"fmt"

	"github.com/Wolfxxxz/professionGo/12Packet/store"
)

func main() {
	b := store.Product{
		Name:     "Kabak",
		Category: "Watersports",
	}
	fmt.Println(b)
}

// store/product.go
package store

//Продукт описывает элемент для продажи
type Product struct {
	Name, Category string  //Имя и тип продукта
	price          float64 //Стоимость НЕимпортируемая из за маленькой заглавной
}
## Создание Функции конструктора для доступа ко всем полям
// store/product.go
package store

//Продукт описывает элемент для продажи
type Product struct {
	Name, Category string  //Имя и тип продукта
	price          float64 //Стоимость
}

// store/product.go
// Функция конструктора
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// main.go
package main

import (
	"fmt"

	"github.com/Wolfxxxz/professionGo/12Packet/store"
)

func main() {
	b := store.Product{
		Name:     "Kabak",
		Category: "Watersports",
	}
	fmt.Println(b)

	newProduct2 := store.NewProduct("Boat", "Watersport", 99)
	fmt.Println(newProduct2) //&{Boat Watersport 99}
}
## Связи между наоговыми ставками
// main.go
package main

import (
	"fmt"

	"github.com/Wolfxxxz/professionGo/12Packet/store"
)

func main() {

	newProduct2 := store.NewProduct("Boat", "Watersport", 200)
	fmt.Println(newProduct2) //&{Boat Watersport 99}

}

//---------------
// store/tax.go
package store

const defaultTaxRate float64 = 0.2

//Минимальный порог
const minThreshold = 10

//Налоговая ставка
type taxRate struct {
	rate, threshold float64
}

//Новая налоговая ставка
func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold //const могут совокуплятся сами:)
	}
	return &taxRate{rate, threshold}
}

//Калькулятор налоговой ставки
func (taxRate *taxRate) calcTax(product *Product) float64 {
	if product.price > taxRate.threshold {
		return product.price + (product.price * taxRate.rate)
	}
	return product.price
}
//---------------
// Комментировать можно и название пакета
// store/product.go
package store

var standardTax = newTaxRate(0.25, 20)

//Продукт описывает элемент для продажи
type Product struct {
	Name, Category string  //Имя и тип продукта
	price          float64 //Стоимость
}

// Функция конструктора
// Новый продукт
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Считать новый налог
func (p *Product) Price() float64 {
	return standardTax.calcTax(p)
}

// Новая цена продукта
func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
## Ловушка переопределения (Повторяющиеся названия пакетов)
//packages/fmt/formats.go
package fmt

import "strconv"

func ToCurrency(amount float64) string {
	return "$" + strconv.FormatFloat(amount, 'f', 2, 64)
}
//-----------
// main.go
package main

import (
	CurrencyFmt "12Packet/packages/fmt"
	"12Packet/store"
	"fmt"
)

func main() {

	product := store.NewProduct("Boat", "Watersport", 200)
	fmt.Println(product) //&{Boat Watersport 99}

	fmt.Println("Price:", CurrencyFmt.ToCurrency(product.Price()))

}
//---------- Или так "."
package main

import (
	. "12Packet/packages/fmt"
	"12Packet/store"
	"fmt"
)

func main() {

	product := store.NewProduct("Boat", "Watersport", 200)
	fmt.Println(product) //&{Boat Watersport 99}

	fmt.Println("Price:", ToCurrency(product.Price()))

}
//--------------
## Func init()
Отрабатывает только если был вызван пакет где она находится, 
можно использовать заглушенный импорт если в пакете требуется только эта функция
## Поиск пакетов
https://pkg.go.dev
https://github.com/golang/go/wiki/Projects
## ПАкет color
go get github.com/fatih/color@v1.10.0
## Удалить зависимости   ***go mod tidy***