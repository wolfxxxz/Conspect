# 9 Function type
## Hello
func main() {
	fmt.Println("Hello function Types")
}
## Функция как переменная ***var calcFunc func(float64) float64***
//ОГОНЬ
func main() {
	fmt.Println("Hello function Types")
	products := map[string]float64{
		"Polozy": 555,
		"sholom": 22.5,
	}
	for product, price := range products {
		var calcFunc func(float64) float64 //Обьявление переменной типа функция
		if price > 100 {
			calcFunc = calcWith
		} else {
			calcFunc = calcWithout
		}
		totalPrice := calcFunc(price)
		fmt.Println(product, totalPrice)
	}
}
## нулевое значение calcFunc == nil
 var calcFunc func(float64) float64 
 fmt.Println(calcFunc == nil) // true

 func main() {
	fmt.Println("Hello function Types")
	products := map[string]float64{
		"Polozy": 555,
		"sholom": 22.5,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println(calcFunc == nil) // true
		if price > 100 {
			calcFunc = calcWith
		} else {
			calcFunc = calcWithout
		}
		totalPrice := calcFunc(price)
		fmt.Println(calcFunc == nil) // false
		fmt.Println(product, totalPrice)
	}
}

func calcWith(price float64) float64 {
	return (price * 0.2) + price
}

func calcWithout(price float64) float64 {
	return price
}
## функция как аргумент функции //calculator func(float64) float64
### Example 1
//Если сигнатура функции удовлетворяет запросу сигнатуры 
//(указаны одинаковые входящие и возвращаемые параметры)
// то вместо неё можно поставить любую другую функцию

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func main() {
	fmt.Println("Hello function Types")
	products := map[string]float64{
		"Polozja": 557,
		"sholom":  22.5,
	}
	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWith)
		} else {
			printPrice(product, price, calcWithout)
		}
	}
}

func calcWith(price float64) float64 {
	return (price * 0.2) + price
}

func calcWithout(price float64) float64 {
	return price
}
### Example 2
func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWith
	}
	return calcWithout
}

func main() {

	products := map[string]float64{
		"Polozja": 557,
		"sholom":  22.5,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

func calcWith(price float64) float64 {
	return (price * 0.2) + price
}

func calcWithout(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
### slice сигнатур (var B = []func(a int) int{One, Tue, Thre, Four})
func Open(a int, oneString func(a int) int) {
	fmt.Println(a, " ", oneString(a))
}
func main() {
	var a = []int{1, 2, 3, 4}
	var B = []func(a int) int{One, Tue, Thre, Four}

	for i, v := range a {
		Open(v, B[i])
	}
}

func One(a int) int {
	return a * 2
}
func Tue(a int) int {
	return a * 3
}
func Thre(a int) int {
	return a * 4
}
func Four(a int) int {
	return a * 5
}
### Calkulator na сигнатурах функций
func main() {
	scanner := bufio.NewScanner((os.Stdin))
	for {
		fmt.Print("Введите выражение: ")
		scanner.Scan()
		expression := scanner.Text()

		if expression == "" {
			fmt.Println("Выход")
			return
		}
		Calculator(expression)
	}

}
func Calculator(s string) {
	// чистим пробелы в начале и в конце// на всякий случай
	s = strings.TrimSpace(s)
	// извлекаем числа и знак операции
	var a, b int
	var op string
	if i := strings.IndexAny(s, "+-*/"); i != -1 { // проверяем наличие знака операции
		op = string(s[i])
		a, _ = strconv.Atoi(strings.TrimSpace(s[:i]))   // извлекаем первое число
		b, _ = strconv.Atoi(strings.TrimSpace(s[i+1:])) // извлекаем второе число
	} else {
		fmt.Println("Недопустимый знак операции")
		return
	}
	// выполнение математической операции в зависимости от знака операции
	Count(a, b, op, SignatureFunc[op])
}

var SignatureFunc = map[string]func(a int, b int) int{
	"*": Multiplicate,
	"/": Division,
	"-": Minus,
	"+": Plus,
}

func Count(a, b int, d string, oneString func(a int, b int) int) {
	fmt.Println(a, " ", d, " ", b, " = ", oneString(a, b))
}

func Multiplicate(a int, b int) int {
	return a * b
}
func Division(a int, b int) int {
	return a / b
}
func Minus(a int, b int) int {
	return a - b
}
func Plus(a int, b int) int {
	return a + b
}
## Псевдонимы функциональных типов
// Псевдоним сигнатуры функции
type calcFunc func(float64) float64

func main() {

	products := map[string]float64{
		"Polozja": 557,
		"sholom":  22.5,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

func calcWith(price float64) float64 {
	return (price * 0.2) + price
}

func calcWithout(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator calcFunc) { //Заменяем на псевдоним
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc { //Заменяем на псевдоним
	if price > 100 {
		return calcWith
	}
	return calcWithout
}
## Использование литерального синтаксиса функции (Анонимные)
### Example 1
type calcFunc func(float64) float64

func main() {

	products := map[string]float64{
		"Polozja": 557,
		"sholom":  22.5,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

// Переехали внутрь selectCalculator и стали анонимными
//	func calcWith(price float64) float64 {
//		return (price * 0.2) + price
//	}

//func calcWithout(price float64) float64 {
//	return price
//}

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 { // Литеральный синтаксис
			return price + (price * 0.2)
		}
		return withTax
	}
	var withoutTax calcFunc = func(price float64) float64 { // Литеральный синтаксис
		//withoutTax := func(price float64) float64 { // Литеральный синтаксис
		return price
	}
	return withoutTax
}
### Example 2 В котором все платят 20%
type calcFunc func(float64) float64

func main() {

	products := map[string]float64{
		"Polozja": 275,
		"sholom":  48.95,
	}
	for product, price := range products {
		printPrice(product, price, func(price float64) float64 { // Анонимная функция
			return price + (price * 0.2)
		})
	}
}

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
## Замыкание функций
### Пример замыкания данных
В этом примере мы создаем анонимную функцию printMessage, которая имеет доступ к переменной message, объявленной вне функции.
Это происходит потому, что функция printMessage замыкает (capture) переменную message, и она становится доступной внутри 
этой функции, даже если она определена за ее пределами.

func main() {
    message := "Hello"

    printMessage := func() {
        fmt.Println(message)
    }

    printMessage()
}
### Функция которая возвращает функцию замыкая при этом данные
type print func(string)

func bomb(s string) print {
	return func(v string) {
		if len(v) > len(s) {
			fmt.Println(v, s)
		} else {
			fmt.Println(s, v)
		}
	}
}

func main() {
	Hello := "Hello function"
	Short := bomb("short")
	long := bomb("this is a very long function")
	fmt.Println("-------------")
	Short(Hello)
	long(Hello)
}
### Фабричная функция которая возвращает функцию с коофициентами 
type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// Функция высшего порядка
func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func main() {

	watersportsProducts := map[string]float64{
		"Polozja": 275,
		"sholom":  48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}
}
## Ранняя оценка
### Simple Example
Ранняя оценка (early binding) при замыкании функции в Golang означает, что значения переменных, 
на которые ссылается замыкание, определяются в момент объявления замыкания, а не в момент вызова замыкания.
Это означает, что при каждом вызове замыкания будут использоваться те же значения переменных, что и при его объявлении.
func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	inc := incrementor()
	fmt.Println(inc()) // Вывод: 1
	fmt.Println(inc()) // Вывод: 2
}
### Как избежать ранней оценки //переназначить елемент
func incrementor(i int) func() int {
	return func() int {
		i++
		return i
	}
}
func main() {
	inc := incrementor(0)
	fmt.Println(inc()) // Вывод: 1
	fmt.Println(inc()) // Вывод: 2

	inc = incrementor(0)
	fmt.Println(inc()) // Вывод: 11
	fmt.Println(inc()) // Вывод: 12
}
### Или воспользоватся поинтером (замыкание по указателю)
var i int = 10

func incrementor(in *int) func() int {
	//i := *in
	return func() int {
		i := *in
		i++
		return i
	}
}

func main() {
	i = 0
	inc := incrementor(&i)
	fmt.Println(inc()) // Вывод: 1
	fmt.Println(inc()) // Вывод: 1
	i = 1
	inc2 := incrementor(&i)
	fmt.Println(inc2()) // Вывод: 2
	fmt.Println(inc2()) // Вывод: 2
} 
### Как избежать ранней оценки поинтеры
type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway = true

// Функция высшего порядка
func priceCalcFactory(threshold, rate float64, zeroPrices *bool) calcFunc {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func main() {

	watersportsProducts := map[string]float64{
		"Polozja": 275,
		"sholom":  48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	prizeGiveaway = true

	waterCalc := priceCalcFactory(100, 0.2, &prizeGiveaway)
	prizeGiveaway = false
	soccerCalc := priceCalcFactory(50, 0.1, &prizeGiveaway)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}
}
