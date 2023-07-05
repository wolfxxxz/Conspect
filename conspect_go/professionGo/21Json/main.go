package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279,
}

type DiscountedProduct struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:",string"`
}

func main() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	encoder.Encode(&dp)
	dp2 := DiscountedProduct{Discount: 10.50}
	encoder.Encode(&dp2)
	fmt.Print(writer.String())
}
