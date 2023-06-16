# Горутины и каналы
## Последовательное выполнение
package main

import (
	"strconv"
)

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

// Псевдоним типа []*Product
type ProductGroup []*Product

// Псевдоним типа мап
type ProductData = map[string]ProductGroup

// Мап Инициализирована
var Products = make(ProductData)

// возвращает строку с знаком доллара
func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

// Заполняет мап [Watersports]  &{Lifejacket Watersports 49.95}
func init() {
	// список продуктов
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
			//fmt.Print(p.Category, "  ")
			//fmt.Println(p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}

package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		storeTotal += group.TotalPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

// metod
func (group ProductGroup) TotalPrice(category string) (total float64) {
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
package main

import "fmt"

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete")
}
*/
main function started
Watersports product: Kayak
Watersports product: Lifejacket
Watersports subtotal: $328.95
Soccer product: Soccer Ball
Soccer product: Corner Flags
Soccer product: Stadium
Soccer subtotal: $79554.45
Chess product: Thinking Cap
Chess product: Unsteady Chair
Chess product: Bling-Bling King
Chess subtotal: $1291.00
Total: $81174.40
main function complete
*/
## Горутины
