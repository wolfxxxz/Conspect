// main.go
package main

import (
	"12Packet/packages/baks"
	_ "12Packet/packages/data"
	"12Packet/packages/store/cart"
	"12Packet/store"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	product := store.NewProduct("Boat", "Watersport", 200)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	color.Green("Name:" + cart.CustomerName)
	color.Cyan("Total:" + baks.ToCurrency(cart.GetTotal()))
	fmt.Println("Name:", cart.CustomerName)

	fmt.Println("Product:", product.Name, "Price:", baks.ToCurrency(product.Price()))

}
