# 8 func
## Hello
func main() {
	Print()
}
func Print() {
	fmt.Println("Hello func")
}
## Print(a, b int, v float64)  Тип переменной
func main() {
	a, b := 5, 10
	c := 5.5
	Print(a, b, c) //0.5
}

func Print(a, b int, v float64) { //a без типа
	fmt.Println(v + float64(a-b))
}
## пропуск елемента в обьявлении функции
***в шапке пропущен в теле не используется**
func main() {
	a, b := 5, 10
	c := 5.5
	Print(a, b, c) //10.5
}

func Print(a, _ int, v float64) { //a без типа
	fmt.Println(v + float64(a))
}
## func printName(product int, nameS ...string) // Вариативный параметр(по сути массив)
### product int, nameS ...string
func main() {
	printName(5, "fit", "bit") //
}

func printName(product int, nameS ...string) { // var nameS []string
	for _, name := range nameS {
		fmt.Println(name)
	}
} 
### printName(5, namesi...)
func main() {
	var namesi = []string{"first", "second"}
	printName(5, namesi...) //
}

func printName(product int, nameS ...string) {
	for _, name := range nameS {
		fmt.Println(name)
	}
}
## Pointer & func()
func main() {
	var first = 1
	var second = 50
	fmt.Println("first: ", first)  //1
	fmt.Println("second: ", second) // 50
	Revers(&first, &second)
	fmt.Println("first: ", first)  //50
	fmt.Println("second: ", second) //1
}

func Revers(a, b *int) {
	temp := *a //var temp int = 1
	*a = *b    //положить в ячейку памяти a = 50
	*b = temp  //положить в ячейку памяти b = 1
}
## Return без инициализации переменной
func main() {
	var first = 1
	var second = 50
	fmt.Println(Sum(first, second))

}

func Sum(a, b int) float64 {
	return float64(a + b)
}
## Возврат нескольких результатов
### Example 1
func main() {
	var first = 1
	var second = 50
	fmt.Println(Sum(first, second)) //51 false
    a, boolean := Sum(first, second)
	fmt.Println(a, boolean)         ////51 false
}

func Sum(a, b int) (float64, bool) {
	if float64(a+b) > 100 {
		return float64(a + b), true
	} else {
		return float64(a + b), false
	}
}
### Example 2
func main() {
	benefit := map[string]float64{
		"Andrew": 500,
		"John":   25,
	}
	for key, val := range benefit {
		taxAmount, taxDue := calcProfit(val)
		if taxDue {
			fmt.Println("Product: ", key, " Tax: ", taxAmount)
		} else {
			fmt.Println("Product: ", key, "Not tax due")
		}
	}
}

func calcProfit(profit float64) (float64, bool) {
	if profit > 250 {
		return profit * 0.5, true
	}
	return 0, false
}
## возврат именованых возвращаемых данных (или пустой return) 
//FindShortAndLong(a, b string) (short, long string)
### Example 1
func main() {
	var a, b = "go", "walk"
	short, long := FindShortAndLong(b, a)
	fmt.Println("short: ", short)
	fmt.Println("long: ", long)
}

func FindShortAndLong(a, b string) (short, long string) {
	if a > b {
		long = a
		short = b
	} else {
		long = b
		short = a
	}
	return
}
### Example 2
func main() {
	benefit := map[string]float64{
		"Andrew": 500,
		"John":   25.55,
	}
	total1, tax1 := calcTotalPrice(benefit, 10)
	fmt.Println(total1, tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println(total2, tax2)
}

func calcProfit(profit float64) (float64, bool) {
	if profit > 250 {
		return profit * 0.5, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcProfit(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}
# defer
func main() {
	defer fmt.Println("first defer")
	var a, b = "go", "walk"
	short, long := FindShortAndLong(b, a)
	defer fmt.Println("second defer")
	fmt.Println("short: ", short)
	fmt.Println("long: ", long)
	defer fmt.Println("third defer")
}

func FindShortAndLong(a, b string) (short, long string) {
	defer fmt.Println("findShortAndLong defer")
	if a > b {
		long = a
		short = b
	} else {
		long = b
		short = a
	}
	return
}

/*
findShortAndLong defer
short:  go
long:  walk
third defer
second defer
first defer
*/
