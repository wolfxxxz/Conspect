package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	Name, Category string
	Price          float64
}
type Customer struct {
	Name, City string
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		switch elemValue.Kind() {
		case reflect.Bool:
			var val bool = elemValue.Bool()
			Printfln("Bool: %v", val)
		case reflect.Int:
			var val int64 = elemValue.Int()
			Printfln("Int: %v", val)
		case reflect.Float32, reflect.Float64:
			var val float64 = elemValue.Float()
			Printfln("Float: %v", val)
		case reflect.String:
			var val string = elemValue.String()
			Printfln("String: %v", val)
		case reflect.Ptr:
			var val reflect.Value = elemValue.Elem()
			if val.Kind() == reflect.Int {
				Printfln("Pointer to Int: %v", val.Int())
			}
		default:
			Printfln("Other: %v", elemValue.String())
		}
	}
}

type Payment struct {
	Currency string
	Amount   float64
}

func main() {
	product := Product{
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	number := 100
	printDetails(true, 10, 23.30, "Alice", &number, product)
}
