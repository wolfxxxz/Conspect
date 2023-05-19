# 3 линтер - проверка + коменты
go install github.com/mgechev/revive@latest
// Ставить в каждый проэкт сразу после go mod init иначе тупит жёстко
*** revive** - показывает ошибки
*** go doc -all** - показывает коменты
***// revive:disable:exported*** - отключить эту функцию от комментов
func PrintHello() {
	fmt.Println("hello")
}
146
# 4 defolt Value
https://pkg.go.dev/std - Стандартная библиотека golang
## const И строгая типизация
const price float32 = 275
const tax float32 = 27.5
***const quantity = 2**

func main() {
	fmt.Println(quantity * (price + tax)) // 605
}
## var и сторогая типизация (mismatched types int and float32)
func main() {
	var price float32 = 275
	var tax float32 = 27.5
	var quantity = 2
	fmt.Println(quantity * (price + tax)) //invalid operation: quantity * (price + tax) (mismatched types int and float32)
}
## iota - даёт последовательно значение каждой новой константе
const (
	Watersports = iota //0
	Soccer //1
	Chess //2
	d //3
	fff //4
	hhh
	jjj
	kkk
	uuu //8
)

func main() {
	fmt.Println(Watersports) //0
	fmt.Println(Soccer)      //1
	fmt.Println(Chess)       //2
	fmt.Println(uuu)         //8
}
## Компилятор именно приобразует тип к типу а не переписывает тип константы
const (
	Watersports = iota
	Soccer
	Chess
	d
	fff
	hhh
	jjj
	kkk
	uuu
)

func main() {
	fmt.Println(Watersports) //0
	fmt.Println(Soccer)      //1
	fmt.Println(Chess)       //2
	fmt.Println(uuu)         //8
	fmt.Printf("%T\n", uuu)  // int
	var Integer int = 55
	fmt.Println(Integer + uuu) // 63
	var Floater float32 = 22.5
	fmt.Println(Floater + uuu) //30.5
	fmt.Printf("%T\n", uuu)    // int
}
## float по умолчанию var UnknownFloat = 55 (float64)
func main() {
	var FloatUnknown = 222.5
	var Float32 float32 = 15
	fmt.Println(FloatUnknown + Float32) // mistake// mismatched types
}
## Дефолтные значения
int - 0
string - ""
rune - 0
bool - false
uint - 0
float - 0
byte - 0
pointer - nil
# Pointer
## Simple
func main() {
	/*
		first := 100
		second := first // просто копия
		first++
		fmt.Println(first)  //101
		fmt.Println(second) //100
	*/

	first := 100
	var second *int = &first
	//second := &first // pointer
	first++
	fmt.Println(first)   //101
	fmt.Println(*second) //101 // (*) даёт указание взять именно значение а не адресс в памяти
	*second++            // перейти по ссилке и увеличить значение
	fmt.Println(first)   //102
    var third *int
	fmt.Println(third) // nil - нулевое значение
    //fmt.Println(*third) // ошибка компиляции *пустое значение
    third = second
	fmt.Println(*third) // можно обойтись и одной звездой // 102
}
## Массивы и поинтеры
func main() {
	names := [3]string{"Alise", "Charlie", "Bob"}
	firstName := &names[1]
	secondName := names[1]
	fmt.Println(*firstName) // Charlie
	fmt.Println(secondName) // Charlie
	sort.Strings(names[:])
	fmt.Println(*firstName) // Bob // Теперь во второй ячейке имя Bob
	fmt.Println(secondName) // Charlie
}
# 5 Простые действия
## Оператор остатка %
func main() {
	fmt.Println(6 % 2) //0
	fmt.Println(5 % 2) //1
	fmt.Println(2 % 5) //2
	fmt.Println(2 % 6) //2
    d := -2 % 6         //
	dd := math.Abs(float64(d)) //абсолютное значение чтобы это не значило
	fmt.Println(dd)
}
## Емкость типов 
func main() {
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal)
	// переполнение в два раза
	fmt.Println(intVal * 2) //-2
	// перешло границы
	fmt.Println(floatVal * 2) //+inf
	// переполнение
	fmt.Println(math.IsInf((floatVal * 2), 0)) //true
}
## bool
func main() {
	first := 100
	const second = 200.00
	equel := first == second
	fmt.Println(equel) // false
	notEquel := first != second
	fmt.Println(notEquel) // true
}
## сравнение указателей
first := 100
	pointfirst := &first
	pointfirst2 := &first
	fmt.Println(pointfirst == pointfirst2) //true
	second := 100
	fmt.Println(&second == pointfirst) //false (адресса в памяти не равны)
	fmt.Println(*pointfirst == second) //true
## logik and or not (&& , || , !)
func main() {
	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)     // true
	fmt.Println(sportsCar)     // false
	fmt.Println(canCategorize) // false
}
## math. Ceil Floor Round RoundToEven 
func main() {
	kayak := 275
	soccerBall := 19.50
	total := kayak + int(soccerBall)
	fmt.Println(total)       //294
	fmt.Println(int8(total)) //38 переполнение
	total2 := soccerBall + float64(kayak)
	fmt.Println(total2) // 294.5
	total3 := kayak + int(math.Round(soccerBall))
	fmt.Println(total3) //295

	// Округление float64
	floatOne := 25.5
	fmt.Println(floatOne)                   // 25.5
	fmt.Println(math.Ceil(floatOne))        //26
	fmt.Println(math.Floor(floatOne))       //25
	fmt.Println(math.Round(floatOne))       //26
	fmt.Println(math.RoundToEven(floatOne)) //26
}
## Парсинг строк strconv
func main() {
	// Ищем bool strconv.ParseBool()
	boolean := "true"
	a, c := strconv.ParseBool(boolean)
	fmt.Println(a, c) // true <nil>

	// ParseFloat
	cc := "222.88"
	d, _ := strconv.ParseFloat(cc, 32) //222.8800048828125
	fmt.Println(d)

	// ParseInt
	ccc := "897"
	dd, _ := strconv.ParseInt(ccc, 10, 0) // 897
	fmt.Println(dd)

	// Atoi
	ddd, _ := strconv.Atoi(ccc) // 897
	fmt.Println(ddd)

	// Itoa
	rrr := strconv.Itoa(ddd) // "897"
	fmt.Println(rrr)
}
## Err
func main() {
	val1 := "0"

	if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
		fmt.Println("Parsed value: ", bool1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}
# 6 potok (if else, for and switch, метка)
## If else названия значений всё равно
type Car struct {
	prise int
	name  string
}

func main() {
	CarFirst := Car{prise: 500, name: "Ferary"}
	CarSecond := Car{prise: 88, name: "Ford"}
	CarThird := Car{prise: 3, name: "Mistake"}
	Garage := []Car{CarFirst, CarSecond, CarThird}

	for i, v := range Garage {
		if v.prise >= 200 {
			Ppp := v.name //
			fmt.Printf("Expensive car %v: \n", Ppp)
		} else if v.prise <= 199 && v.prise >= 10 {
			Ppp := v.name //
			fmt.Printf("Cheap car %v: \n", Ppp)
		} else {
			Ppp := false //
			fmt.Println(i+1, v, Ppp)
		}
	}
}
### Сравнение в два захода (if priseCar, err := strconv.Atoi(v.prise); err == nil )
// 
type Car struct {
	prise string
	name  string
}

func main() {
	CarFirst := Car{prise: "500", name: "Ferary"}
	CarSecond := Car{prise: "88", name: "Ford"}
	CarThird := Car{}
	Garage := []Car{CarFirst, CarSecond, CarThird}

	for i, v := range Garage {
		if priseCar, err := strconv.Atoi(v.prise); err == nil {
			Ppp := priseCar
			fmt.Printf("Car %v and Price %v: \n", i+1, Ppp)
		} else {
			fmt.Printf("Car %v is not exist", i+1)
		}
	}
}
// Car 1 and Price 500: 
// Car 2 and Price 88: 
// Car 3 is not exist
### Range string
type Car struct {
	prise string
	name  string
}

func main() {
	CarFirst := Car{prise: "500", name: "Ferary"}
	CarSecond := Car{prise: "88", name: "Ford"}
	CarThird := Car{}
	Garage := []Car{CarFirst, CarSecond, CarThird}

	var nameCar string
	var nameCarRune rune

	for i, v := range Garage[0].name {
		fmt.Println("Index", i, "Character", string(v), "rune", v)
		nameCar = nameCar + string(v)
		nameCarRune = nameCarRune + v // Руны это вам не стринг, конкат не работает
	}
	fmt.Println(nameCar) //Ferary
	fmt.Println(string(nameCarRune)) //ё
}
## for i := range Car
func main() {
	Car := "Lamborjiny"

	for i := range Car {
		fmt.Print(" ", i)
	}
    // 0 1 2 3 4 5 6 7 8 9
}
## Switch 
### Exampl 1 Simple
func main() {
	Car := "Lamborjiny"

	for i, v := range Car {
		switch v {
		case 'L':
			fmt.Println("L and position ", i) //L and position  0
		case 'j':
			fmt.Println("j and position", i) //j and position 6
		}
	}
}
### Example 2 several value in filter (case 'L', 'y')
func main() {
	Car := "Lamborjiny"

	for i, v := range Car {
		switch v {
		case 'L', 'y':
			fmt.Printf("%v and position %v\n", string(v), i) //L and position  0 / y and position 9
		case 'j', 'a':
			fmt.Printf("%v and position %v\n", string(v), i) //j and position 6 /j and position 6
		}
	}
}
### break and switch
func main() {
	Car := "Lamborjiny"
	for i, v := range Car {
		switch v {
		case 'L', 'a':
			fmt.Printf("%v and position %v\n", string(v), i) //L and position  0 / y and position 9
		case 'j', 'y':
			if v == 'j' { // Работает только если до Print с чем связано непонимаю
				break
			}
			fmt.Printf("%v and position %v\n", string(v), i) //j and position 6
		}
	}
}
### fallthrough Провалы (Странная хрень)
func main() {
	Car := "Lamborjlny"
	for i, v := range Car {
		switch v {
		case 'L':
			fmt.Printf("%v and position %v\n", string(v), i) //L and position  0
			fallthrough
		case 'l':
			fmt.Printf("%v and position %v\n", string(v), i) //L and position  0 /l and position 7
			fallthrough
		case 'y':
			fmt.Printf("%v and position %v\n", string(v), i) //L and position  0 /l and position 7 /y and position 9
		}
	}
}
### default 
func main() {
	Car := "Lamborjlny"
	for i, v := range Car {
		switch v {
		case 'L':
			fmt.Printf("%v and position %v\n", string(v), i) //L and position 0
		case 'y':
			fmt.Printf("%v and position %v\n", string(v), i) //y and position 9
		default:
			fmt.Printf("Default %v and position %v\n", string(v), i) //Default a and position 1, 2, 3, 4, 5, 6, 7, 8
		}
	}
}
### switch condition (условие при обьявлении)
#### Exampl1 
func main() {
	for i := 0; i <= 20; i++ {
		switch i / 2 {
		case 2, 3, 5, 7:
			fmt.Println(i) //4,5,6,7,10,11,14,15
		default:
			fmt.Println("default", i) //0,1,2,3,...20
		}
	}
}
#### Exampl 2 Условие с присвоением переменной (switch val := i / 2; val)
func main() {
	for i := 0; i <= 10; i++ {
		switch val := i / 2; val {
		case 2, 3, 5, 7:
			fmt.Println(val) //2,2,3,3,5
		default:
			fmt.Println("default", i) //0,1,2,3,8,9
		}
	}
}
#### Exampl 3 Условия в case
func main() {
	for i := 0; i <= 10; i++ {
		switch {
		case i == 0:
			fmt.Println("zero value", i) //
		case i < 3:
			fmt.Println("i < 3", i)
		case i >= 3 && i < 7:
			fmt.Println(">= 3 && i < 7", i)
		default:
			fmt.Println("default", i) //
		}
	}
}
## goto target (Цикл по метке)
func main() {
	counter := 0
target:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
	fmt.Println("Over")
}
# 7 Slice, Arr and Map
## Slice and Arr
### Копировать срез и добавить элементы в одну строку
// appendNames := append(names[:], "hat", "gloves")
func main() {
	var names = [3]string{}

	names[0] = "gin"
	names[1] = "vip"
	names[2] = "gon"

	otherArr := &names

	otherArr[0] = "A very long name"
	fmt.Println(names)

	appendNames := append(names[:], "hat", "gloves")
	fmt.Println(appendNames)
}
### нет смысла создавать ссылки на срез - если не нужно его передавать в функцию:)
func main() {
	names := make([]string, 3, 6)

	names[0] = "gin"
	names[1] = "vip"
	names[2] = "gon"

	fmt.Println(names)

	otherArr := []*string{&names[0], &names[1], &names[2]}

	*otherArr[0] = "long Name"

	fmt.Println(names) //[long Name vip gon hat gloves]
}
### Делаем что то со слайсом по ссылке
func main() {
	secondSlice := []int{1}
	AppendMoreInt(&secondSlice)
	fmt.Printf("len %v, and capacity %v \n", len(secondSlice), cap(secondSlice)) //len 51, and capacity 64
}

func AppendMoreInt(slice *[]int) {
	for i := 0; i <= 50; i++ {
		*slice = append(*slice, i)
	}
}
### Дополнительная ёмкость (slice2 := make([]int, 0, 100000)) с замером скорости от gpt
Если cap задана изначально то оста'тся только заполнить ячейки, 
в противном варианте с каждым добавлением элемента приходится:
1 Выделить новое место в памяти
2 копировать слайс полностью
3 добавить элемент 
И так с каждым новым элементом
//  
func main() {
	var wg sync.WaitGroup

	start := time.Now()

	// создаем слайс без начальной емкости
	slice1 := make([]int, 0)

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			slice1 = append(slice1, i)
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Slice without initial capacity: %s\n", elapsed)

	start = time.Now()

	// создаем слайс с начальной емкостью 100000
	slice2 := make([]int, 0, 100000)

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			slice2 = append(slice2, i)
		}(i)
	}

	wg.Wait()

	elapsed = time.Since(start)
	fmt.Printf("Slice with initial capacity: %s\n", elapsed)
}
### Так делать нельзя, капасити инициализируется только через append
    secondSlice := make([]int, 0, 100)
	secondSlice = append(secondSlice, 10) // elem 0 value 1
	secondSlice[1] = 20 // goroutines filed
### Реши что нужно копия массива или ссылки на массив
***Осторожно это может быть больно**
- если взять slice от масива а потом изменить элемент в slice то елемент изменится и в массиве
func main() {
	var Arr = [3]string{"one", "tue", "three"}
	fmt.Println(Arr) //[one tue three]

	oneSlice := Arr[:]
	oneSlice[0] = "four"
	fmt.Println(Arr) //[four tue three]
	// Пока мы не добавляем елементы в слайс
	// он ведёт себя как ссылка на массив
	oneSlice = append(oneSlice, "five") // Добавили новый элемент и слайс стал независим от arr
	oneSlice[1] = "six"
	fmt.Println(Arr) // [four tue three]
}
### Copy Копирует нужные элементы в другое место (copy(oneSlice[:], Arr[:]))
***При этом все зависимости между массивами исчезают**
***Обязательно создать срез через make**
***Указать правильное количество елементов**
func main() {
	var Arr = []string{"tue", "three", "four"}
	fmt.Println(Arr) //[tue three four]

	//oneSlice := []string{} // Будет ошибка
	oneSlice := make([]string, 4)
	oneSlice[0] = "one"
	copy(oneSlice[1:], Arr)

	fmt.Println(oneSlice) //[one tue three four]
}
### Копирование диапазонов Slice
***Нельзя указать размер больше реального размера слайса**
***При копировании происходит только замена существующих значений**
var AllFigures = []string{"1", "2", "3", "4", "5"}
	var CoupleFigures = []string{"six", "seven"}

	//copy(AllFigures[:], CoupleFigures[:])   // [six seven 3 4 5]
	//copy(AllFigures[3:], CoupleFigures[:])  // [1 2 3 six seven]
	//copy(AllFigures[4:], CoupleFigures[1:]) // [1 2 3 4 seven]
	copy(AllFigures[:1], CoupleFigures[:])    // [six 2 3 4 5]
	fmt.Println(AllFigures)
    /*
	newSlice := make([]string, 7)
	newSlice[0] = "8"
	newSlice[1] = "9"
	copy(newSlice[2:7], AllFigures[:]) // [8 9 1 2 3 4 5]

	fmt.Println(newSlice)*/
#### Practice
func main() {
	AllFigures := make([]string, 7)
	AllFigures[2] = "3"
	AllFigures[3] = "4"
	AllFigures[4] = "5"

	newSlice := []string{"one", "tue"}
	copy(AllFigures[:2], newSlice[:]) // [one tue 3 4 5  ]
	fmt.Println(AllFigures)

	var CoupleFigures = []string{"six", "seven"}

	copy(AllFigures[5:], CoupleFigures[:]) // [1 2 3 4 5 six seven]
	fmt.Println(AllFigures)
}
### Удалить едемент со среза
var Arr = []string{"one", "tue", "2", "three", "four"}
	fmt.Println(Arr) //
	deleted := append(Arr[:2], Arr[3:]...)
	fmt.Println(deleted) // [one tue three four]
### Преобразование из Slice в Arr
    var Slice = []string{"one", "tue", "three", "four"}
	fmt.Println(Slice) //	
	ArrPtr := (*[4]string)(Slice)
	Arr := *ArrPtr
	fmt.Print(Arr)
## Map 
### обьявление map 
        products := make(map[string]float64, 10)
		products["25.5"] = 25.5
		products["33.3"] = 33.3
***Или так**
    products := map[string]float64{
		"gorwok": 225.25,
		"cvetok": 33.5,
		"paket":  0,
	}
    fmt.Println(len(products)) //3
### Проверка на наличие в map
func main() {
	products := map[string]float64{
		"gorwok": 225.25,
		"cvetok": 33.5,
		"paket":  0,
	}

	fmt.Println(len(products)) //3

    if value, ok := products["paket"]; ok { // Stored value:  0
		fmt.Println("Stored value: ", value)
	} else {
		fmt.Println("No stored value")
	}

	//value, ok := products["paket2"]  // No stored value
	/*value, ok := products["paket"] // Stored value:  0
	if ok {
		fmt.Println("Stored value: ", value)
	} else {
		fmt.Println("No stored value")
	}*/
}
### Удаление с карты  //delete(products, "paket")
func main() {
	products := map[string]float64{
		"gorwok": 225.25,
		"cvetok": 33.5,
		"paket":  0,
	}

	fmt.Println(len(products)) //3

	delete(products, "paket")

	fmt.Println(len(products)) //2

	if value, ok := products["paket"]; ok { // Stored value:  0
		fmt.Println("Stored value: ", value)
	} else {
		fmt.Println("No stored value")
	}
}
### for key, value := range map
func main() {
	products := map[string]float64{
		"gorwok": 225.25,
		"cvetok": 33.5,
		"paket":  0,
	}

	for key, value := range products {
		fmt.Print(key, " ", value, " || ") //gorwok 225.25 || cvetok 33.5 || paket 0 ||
	}
}
### Sort map - Копировать ключи в слайс а потом по ключу искать значение в мап
func main() {
	products := map[string]float64{
		"gorwok": 225.25,
		"cvetok": 33.5,
		"paket":  0,
	}

	SliceKey := make([]string, 0, len(products))
	for key, _ := range products {
		SliceKey = append(SliceKey, key)
	}

	sort.Strings(SliceKey)
	fmt.Println(SliceKey[0], products[SliceKey[0]]) //cvetok 33.5
}
### Map как Func
var SignatureFunc = map[string]func(a int, b int) int{
	"*": Multiplicate,
	"/": Division,
	"-": Minus,
	"+": Plus,
}

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
### Map как cлайс
func main() {

	// Инициализируем мапу, где ключ - это название категории, а значение - это список строк
	productsByCategory := make(map[string][]string)

	// Добавляем продукты в мапу
	productsByCategory["fruits"] = []string{"apple", "banana", "orange"}
	productsByCategory["vegetables"] = []string{"carrot", "potato", "cucumber"}
	productsByCategory["meat"] = []string{"beef", "pork", "chicken"}

	// Обращаемся к элементам мапы по ключу и работаем со срезом
	fmt.Println(productsByCategory["fruits"][0])       // Выведет "apple"
	fmt.Println(len(productsByCategory["vegetables"])) // Выведет 3

	// Изменяем срез в мапе
	productsByCategory["meat"] = append(productsByCategory["meat"], "lamb")
	fmt.Println(productsByCategory["meat"]) // Выведет [beef pork chicken lamb]
}
### String = []rune(string) => string([]rune) - Круговорот стринг в байт и ...
func main() {
	var wordString = "Бибизянка Dusja"
	var wordRune []rune = []rune(wordString)
	fmt.Println(string(wordRune[2:14])) //бизянка Dusj
}
### Map & Slice func
func main() {
***Мар func**
	MapFunc := make(map[string]func(a, b int) int)
	MapFunc["+"] = func(a, b int) int {
		return a + b
	}
	MapFunc["-"] = func(a, b int) int {
		return a - b
	}

	fmt.Println(MapFunc["+"](3, 3)) //6
	fmt.Println(MapFunc["-"](5, 2)) //3
***Slice func**
	SliceFunc := make([]func(a, b int) int, 0, 2)

	for _, v := range MapFunc {
		SliceFunc = append(SliceFunc, v)
	}
	SliceFunc = append(SliceFunc, func(a, b int) int {
		return a / b
	})

	fmt.Println(SliceFunc[0](2, 2))  // 4, 0, 4, 4...
	fmt.Println(SliceFunc[1](2, 2))  // 4, 0, 4, 4...
	fmt.Println(SliceFunc[2](40, 2)) // 20

}
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
# 10 type STR struct
## type Hello struct
type Structos struct {
	hello Structs
}
type Structs struct {
	article string
}

func main() {
	structur := Structos{
		hello: Structs{
			article: "Hello",
		},
	}

	fmt.Printf("   %v  %T", structur.hello.article, structur.hello) // Hello  main.Structs
}
## new возвращает ссылку
var life = new(Structos)
	var Life2 = &Structos{}
	fmt.Println(life)  // &{{}}
	fmt.Println(Life2) // &{{}}
## Заполнение данных
type Structos struct {
	hello string
	price int
}

func main() {

	structur := Structos{"Yehoo", 20}
	structur2 := Structos{
		hello: "Google",
		price: 20000,
	}
	fmt.Println(structur)  //{Yehoo 20}
	fmt.Println(structur2) //{Google 20000}
}
## Встроенные структуры
func main() {
	type Product struct {
		name, category string
		price          float64
	}

	type StockLevel struct {
		Product           //Встроенный тип называть не обязательно
		Alternate Product // Два типа продукт
		count     int
	}

	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275},
		count:   100,
	}
	fmt.Println("Name:", stockItem.Product.name) //Name: Kayak
	fmt.Println("Count:", stockItem.count)       //Count: 100

	fmt.Println(fmt.Sprint("Name ", stockItem.Product.name)) //Name Kayak
}
## Сравнение структур
Структуры можно сравнивать на == **если внутри них нет среза**
## Преобразование между типами структур Kayak(boat2)
**Сравнение только на ==**
func main() {
	type Kayak struct {
		place int
	}

	type Ship struct {
		place int
	}

	boat := Kayak{3}
	boat2 := Ship{2}
	boat3 := Ship{3}

	fmt.Println(boat == Kayak(boat2)) //false
	fmt.Println(boat == Kayak(boat3)) //true
}
## Анонимные типы структур
### Анонимная структура и функция
// Функция принимает любой type который
/*type TTT struct {
	place int      // place int!!!
}*/

func writePlace(val struct{ place int }) {
	fmt.Println("Place:", val.place)
}

func main() {
	type Kayak struct {
		place int
	}

	boat := Kayak{3}
	writePlace(boat) //Place: 3
}
### Присвоение значений анонимной струкдуре JsonEncoder
import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	type Kayak struct {
		name  string
		place int
	}

	boat := Kayak{"Victory", 3}
	var builder strings.Builder

	json.NewEncoder(&builder).Encode(struct {
		BoatName  string
		BoatPlace int
	}{
		BoatName:  boat.name,
		BoatPlace: boat.place,
	})
	fmt.Println(builder.String()) //{"BoatName":"Victory","BoatPlace":3}
}
### Собрать представителей разных структур в слайс & Map без интерфейса
func main() {
	type Kayak struct {
		name  string
		place int
	}
	one := Kayak{"KayaK", 3}
	type Boat struct {
		name  string
		place int
	}
	tue := Boat{"BoaT", 2}
	type Bowl struct {
		name  string
		place int
	}
	three := Bowl{"BowL", 1}

	//Slice
	Slice := []struct {
		name  string
		place int
	}{one, tue}
	Slice = append(Slice, three)

	for i, v := range Slice {
		fmt.Println(i, v.name, v.place)
	}

	//Maps
	Maps := make(map[string]struct {
		name  string
		place int
	})
	Maps[one.name] = one
	Maps[tue.name] = tue
	Maps[three.name] = three

	fmt.Println(Maps[one.name]) //{KayaK 3}
	fmt.Println(Maps["BowL"])   //{BowL 1}

}
## Pointer & struct
type Kayak struct {
	name  string
	place int
}

func calcPlace(i int, k *Kayak) int {
	return i * k.place
}

func main() {
	one := &Kayak{"KayaK", 3}
	fmt.Println(calcPlace(3, &one)) // 9
}
## Функция конструктора для создания обьекта структуры
### Создание
type Kayak struct {
	name  string
	place int
}

func newKayak(name string, place int) *Kayak {
	return &Kayak{name, place}
}

func main() {
	one := []*Kayak{
		newKayak("First", 3),
		newKayak("Second", 2),
	}

	for i, v := range one {
		fmt.Println(i, v)
	}
}
### Profit
type Kayak struct {
	name  string
	place int
	price float64
}
***Функция конструктора**
func newKayak(name string, place int, price float64) *Kayak {
	return &Kayak{name, place, price + 2}
}

func main() {
	one := []*Kayak{
		newKayak("First", 3, 63), //65
		newKayak("Second", 2, 52), //54
	}

	for i, v := range one {
		fmt.Println(i, v)
	}
}
## Встроенный тип по ссылке
### Создание
type Kayak struct {
	name  string
	place int
	price float64
	*Suplier
}
type Suplier struct {
	name, sity string
}

func newKayak(name string, place int, price float64, s *Suplier) *Kayak {
	return &Kayak{name, place, price, s}
}

func main() {
	acme := &Suplier{"John Weak", "Kiev"}
	one := []*Kayak{
		newKayak("First", 3, 63, acme),
		newKayak("Second", 2, 52, acme),
	}

	for i, v := range one {
		fmt.Println(i, v)
		fmt.Println("Suplier:", v.Suplier.name) //Suplier: John Weak
	}

}
### Проблема при копировании 
// Ссылка копируется как ссылка а при изминении ссылки меняется 
// значение и в оригинале
### Решение func copyBoat(kayak *Kayak) Kayak 
type Kayak struct {
	name  string
	place int
	price float64
	*Suplier
}
type Suplier struct {
	name, sity string
}

func newKayak(name string, place int, price float64, s *Suplier) *Kayak {
	return &Kayak{name, place, price, s}
}

// Возвращает не ссылку
func copyBoat(kayak *Kayak) Kayak {
	k := *kayak
	s := *kayak.Suplier
	k.Suplier = &s
	return k
}

func main() {
	acme := &Suplier{"John Weak", "Kiev"}

	one := []*Kayak{
		newKayak("First", 3, 63, acme),
		newKayak("Second", 2, 52, acme),
	}

	three := newKayak("third", 5, 44, acme)
	four := copyBoat(three)
	four.name = "fourth"
	four.Suplier.name = "Bill Terner"

	one = append(one, three, &four) // &four

	for i, v := range one {
		fmt.Println(i+1, "Name: ", v.name)
		fmt.Println("Suplier:", v.Suplier.name) //Suplier: John Weak
	}
}
## Value == nil
# 11 Methods
## metod
type Product struct {
	name, category string
	price          float64
}

func (p *Product) printDetails() {
	fmt.Println("Name:", p.name, "Category:", p.category, "Price ", p.calcTax(0.2, 100))
}

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.price > threshold {
		return p.price + (p.price * rate)
	}
	return p.price
}

func main() {
	products := []*Product{
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	d := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	products = append(products, &d)

	for _, p := range products {
		p.printDetails()
	}
}
## metod нельзя декларировать повторно
## Вызов метода через тип получателя (анонимно)
type Product struct {
	name, category string
	price          float64
}

func (p Product) printDetails() {
	fmt.Println("Name:", p.name, "Category:", p.category, "Price ", p.calcTax(0.2, 100))
}

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.price > threshold {
		return p.price + (p.price * rate)
	}
	return p.price
}

func main() {
	Product.printDetails(Product{"Kayak", "Watersports", 275})
}
## Псевдоним типа и метод для него
type Product struct {
	name, category string
	price          float64
}

// Псевдоним типа
type ProductList []Product

// Метод для него
func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func main() {
	products := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total: ", total)
	}
}
//Category:  Watersports Total:  323.95
//Category:  Soccer Total:  19.5
## Преобразование типов
type Product struct {
	name, category string
	price          float64
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}
//------------------------
func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}
func main() {
	products := ProductList(getProducts())

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total: ", total)
	}
}
## Размещение типов и методов в отдельных файлах
//service.go

package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}
//-------------------
// product.go

package main

type Product struct {
	name, category string
	price          float64
}
//-----------------
//main.go

package main

import "fmt"

func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.5}
	fmt.Println("Product: ", kayak.name)
	fmt.Println("Insurance: ", insurance.description)
}
//cmd go run .
//Product:  Kayak
//Insurance:  Boat Cover
# 11.2 Interface
## Create interface and methods under it
//main.go
package main

import "fmt"

type Expense interface {
	getName() string
	getCoast(annual bool) float64
}

func main() {
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.5},
	}
	for _, expense := range expenses {
		fmt.Println("Expence: ", expense.getName(), "Cost: ", expense.getCoast(true))
	}
}
//-------------
// product.go

package main

type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCoast(_ bool) float64 {
	return p.price
}
//------------
//service.go

package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCoast(_ bool) float64 {
	return s.monthlyFee * float64(s.durationMonths)
}
## Функция принимает интерфейс calcTotal(expenses []Expense) 
func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCoast(true)
	}
	return
}

//func main()
fmt.Println(calcTotal(expenses)) //1349
## Интерфейс как поле структуры
type Account struct {
	accountNumber int
	expenses      []Expense
}

func main() {

	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.5},
	}
	//expenses Интерфейс как поле структуры
	account := Account{
		accountNumber: 777,
		expenses:      expenses,
	}

	// Доступ как к полю структуры account.expenses
	for _, expense := range account.expenses {
		fmt.Println("Expence: ", expense.getName(), "Cost: ", expense.getCoast(true))
	}
	fmt.Println(calcTotal(expenses))
}
## эффект приемников и метод указателей
product := Product{"Kayak", "Watersports", 275}
    // Происходит копирование
***	var expense Expense = product**
	product.price = 100
	fmt.Println(product) //{Kayak Watersports 100}

	fmt.Println(expense) //{Kayak Watersports 275}

	//Доступ через продукт и к данным и к методам
	fmt.Println(product.name, product.category, product.price, product.getCoast(true), product.getName())
	//Через интерфейс только к методам
	fmt.Println(expense.getCoast(true), expense.getName())
------------------------------------------------
product := Product{"Kayak", "Watersports", 275}
    // Происходит копирование АДРЕССА
***	var expense Expense = &product**
	product.price = 100
	fmt.Println(product) //{Kayak Watersports 100}

	fmt.Println(expense) //{Kayak Watersports 100}
-----------------------------------------------
***Вывод**
//Использовать указатели при создании метода
func (p *Product) getName() string {
	return p.name
}

func (p *Product) getCoast(_ bool) float64 {
	return p.price
}
## Утверждение типа (сужение типа) s := expense.(Service)
### Тестирование типа
#### 1
func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
	}

	for _, expense := range expenses {
		// доступны только методы т.к expense - интерфейс
		fmt.Println(expense.getCoast(true), expense.getName())
		// утверждение типа даёт доступ ко всем возможностям типа
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
	}
}
#### Внимательно panic: interface conversion:
func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
***		&Product{"Kayak", "Watersports", 275},**
	}

	for _, expense := range expenses {
		// доступны только методы т.к expense - интерфейс
		fmt.Println(expense.getCoast(true), expense.getName())
		// утверждение типа даёт доступ ко всем возможностям типа
***		s := expense.(Service)//Product вам не Service**
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
	}
}
#### Test Если всё таки типы замешаны в интерфейсе 
func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCoast(true))
		}

	}
}
### Включение динамических типов Switch удобно сортирует значения по типам
for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
		default:
			fmt.Println("Expense", expense.getName(), expense.getCoast(true))
		}
	}
## Пустой интерфейс
### Что может принять пустой интерфейс - ВСЁ
***Как его разобрать*** - switch case
func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.5, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price, "Cathegory", value.category)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default", value)
		}
	}

}
### Функция принимает пустой интерфейс - 
func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price, "Cathegory", value.category)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default", value)
	}
}
func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.5, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		processItem(item)
	}
}
### Функция принимает []interface{}
func processItem(items []interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price, "Cathegory", value.category)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price, "Cathegory", value.category)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default", value)
		}
	}
}

func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.5, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	processItem(data)

}
# 12 Packet
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
# 13 Тип и композиция интерфейса
GO - не использует наследование - вместо этого использует композицию
## Продвижение
***Продвижение не может быть выполнено если существует метод с таким же именем**
***Если хочется продвижения - переименуйте функцию что бы имена отличались***
***Продвижение будет работать не корректно если у структур одинаковые поля**


func (t Table) PriceForQuantity(quantity int) (total float64) {
	total = t.price * float64(quantity)
	return
}

/*
func (t Chair) PriceForQuantity(quantity int) (total float64) {
	total = t.price * float64(quantity)
	return
}*/

type Room struct {
	Chair
	*Table
}

//----------------------
func NewRoom(c Chair, t *Table) *Room {
	return &Room{c, t}
}
//--------------------------------- Поля с одинаковым названием
type Chair struct {
	leg   int
	price float64
}
type Table struct {
	leg   int
	price float64
}
//-------------------------------

// Or
func NewRoomChairPlTable(newChLeg, newTabLeg int, newChPrice, newTabPrice float64) *Room {
	return &Room{NewChair(newChLeg, newChPrice), NewTable(newTabLeg, newTabPrice)}
}

//---------------------
func (r *Room) QuantityLeg(quantoty int) (total int) {
	total = (r.Chair.leg + r.Table.leg) * quantoty
	return
}

//------------------------------


func NewChair(newLeg int, newPrice float64) Chair {
	return Chair{newLeg, newPrice}
}



func NewTable(newLeg int, newPrice float64) *Table {
	return &Table{newLeg, newPrice}
}

func (t *Table) QuantityLeg(quantity int) (total int) {
	total = t.leg * quantity
	return
}

//--------------------------------------
func main() {
	// room2.PriceForQuantity(2) - Продвижение
	// room2.Table.PriceForQuantity(2) - без продвижения

	room2 := NewRoomChairPlTable(3, 5, 22.1, 15.7)
	fmt.Println(room2.Table.PriceForQuantity(2))
	fmt.Println(room2.PriceForQuantity(2), room2.Chair, room2.Table)

}
## Коллекция типов через интерфейс
type Chair struct {
	leg   int
	price float64
}

func NewChair(newLeg int, newPrice float64) Chair {
	return Chair{newLeg, newPrice}
}

type Table struct {
	name  string
	leg   int
	price float64
	id    int
}

func NewTable(name string, newLeg int, newPrice float64, id int) *Table {
	return &Table{name, newLeg, newPrice, id}
}

// -------------------------------------------------------------
func (t *Table) Price(tax float64) (total float64) {
	total = (t.price * tax) + t.price
	return
}
func (t *Chair) Price(tax float64) (total float64) {
	total = (t.price * tax) + t.price
	return
}

// -------------------------------------------------------------
type onlySliceOrMapOrArr interface {
	Price(tax float64) float64
}

func main() {
	table1 := NewTable("stul", 4, 22.5, 1)
	chair1 := NewChair(3, 12.0)
	d := []onlySliceOrMapOrArr{table1, &chair1}
	for _, v := range d {
		fmt.Println(v.Price(0.3))
	}

}
## Составление интерфейсов
- Смотри интерфейс product.go
type Describable interface {
	GetName() string
	GetCategory() string
	ItemForSale
}
# 14 # Горутины и каналы







# Куча (heap) & стек (stack)
***"куча"** (англ. heap) — это область памяти, которая используется для **динамического** выделения памяти в процессе выполнения программы. Куча используется для хранения данных, которые могут быть выделены и освобождены во время выполнения программы **в произвольном порядке**.

Как правило, в куче хранятся объекты, созданные с помощью оператора new, а также объекты, созданные с помощью функции make. Куча может быть организована как одна большая область памяти, которая делится между несколькими потоками выполнения программы, или как несколько отдельных областей памяти, которые принадлежат отдельным потокам.

Куча может быть управляемой или неуправляемой. В управляемой куче выделение и освобождение памяти происходят автоматически с помощью сборщика мусора, который определяет, когда объект больше не используется и может быть освобожден. В неуправляемой куче выделение и освобождение памяти происходят явно с помощью функций malloc и free (в языке C), либо с помощью операторов new и delete (в языке C++).

***«куча»** (heap) обычно используется для обозначения области оперативной памяти, в которой выделяются **динамические объекты**(например, переменные, массивы, структуры данных), которые создаются и уничтожаются во время выполнения программы.

Оперативная память в компьютере обычно разделена на две области: стек и кучу. Стек обычно выделяется для временных объектов, таких как локальные переменные и вызовы функций, и работает по принципу последним-вошел-первым-вышел (LIFO). Куча, с другой стороны, выделяется для динамических объектов и используется по принципу первым-вошел-первым-вышел (FIFO).

В языке программирования Go, куча используется для выделения и управления динамической памятью, необходимой для хранения объектов, созданных с помощью ключевого слова new или с помощью функции make. Go имеет сборщик мусора, который автоматически освобождает память, занимаемую объектами, которые больше не используются программой. Это позволяет программистам избежать проблем с утечкой памяти и повышает безопасность и надежность программного обеспечения.
# Объектно-ориентированный язык программирования (ООП) - это стиль программирования, который основан на понятии "объекта". Объект представляет собой некоторую сущность, которая имеет свойства и поведение, а также может взаимодействовать с другими объектами. 

Программисты, использующие ООП, строят программы на основе объектов, которые могут быть созданы на основе описания их свойств и поведения в классах. Класс - это шаблон для создания объектов, который определяет их общую структуру и функциональность. 

Основные принципы ООП включают инкапсуляцию, наследование и полиморфизм. 
***Инкапсуляция** позволяет скрыть реализацию объекта и предоставить только интерфейс для его использования. 
***Наследование** позволяет создавать новые классы на основе существующих, наследуя их свойства и функциональность. 
***Полиморфизм** позволяет объектам разных классов использовать общий интерфейс для выполнения одной и той же операции.

Как правило, ООП используется для разработки больших программных систем, где важна структурированность и возможность повторного использования кода.