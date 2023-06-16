package main

import "fmt"

type Book struct {
	Name  string
	Price float64
}

func (p Book) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	book1 := Book{Name: "Shopin", Price: 55.5}
	var name string
	var price float64
	fmt.Scan(&name, &price)

	book2 := Book{Name: name, Price: price}

	Printfln("Book2 %+v", book2)

	book1Info := book1.String()
	fmt.Println(book1Info) //Product: Shopin, Price: $55.50
	Printfln("Book %+v", book1)
}
