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
