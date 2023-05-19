// store/tax.go
package store

const defaultTaxRate float64 = 0.2

//Минимальный порог
const minThreshold = 10

var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

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
func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.price > taxRate.threshold {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}
