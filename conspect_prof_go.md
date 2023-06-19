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
# 14  Горутины и каналы
## Определения 
Горутины — это легкие потоки, созданные и управляемые средой выполнения Go. Каналы — это конвейеры, передающие значения определенного типа.
## Простой пример
ИСПОЛЬЗОВАНИЕ АДАПТЕРОВ ДЛЯ АСИНХРОННОГО ВЫПОЛНЕНИЯ ФУНКЦИЙ
Не всегда возможно переписать существующие функции или методы для использования каналов, но асинхронно выполнять синхронные функции в оболочке несложно, например:
...
calcTax := func(price float64) float64 {
    return price + (price * 0.2)
}
wrapper := func (price float64, c chan float64)  {
    c <- calcTax(price)
}
resultChannel := make(chan float64)
go wrapper(275, resultChannel)
result := <- resultChannel
fmt.Println("Result:", result)
...
Функция wrapper получает канал, который используется для синхронной отправки значения, полученного при выполнении функции calcTax. Это можно выразить более кратко, определив функцию, не присваивая ее переменной, например:
...
go func (price float64, c chan float64) {
    c <- calcTax(price)
}(275, resultChannel)
...
Синтаксис немного неудобен, потому что аргументы, используемые для вызова функции, выражаются сразу после определения функции. Но результат тот же: синхронная функция может быть выполнена горутиной, а результат будет отправлен через канал.
## chan len and cap
cap(ch)
- размер буфера канала с помощью встроенной функции cap 
**Размер буфера будет сталый**
len(ch)
- количество значений в буфере с помощью функции len
**только в буфере со значениями внутри**
### Exampl simple
func calc(num int, ch chan int) {
	fmt.Println("len empty chan ", len(ch))
	f := num * 2
	ch <- f
	fmt.Println("len ch = ", len(ch))
}
func main() {
	ch := make(chan int, 2)
	fmt.Println("len(ch):", len(ch), " cap(ch):", cap(ch))
	for i := 1; i < 4; i++ {
		go calc(i, ch)
	}
	go func() {

		for v := range ch {
			fmt.Println(v)
		}

	}()

	time.Sleep(time.Second * 3)
}
### Exampl two chan and one select
**select будет выбирать канал пока канал не станет равен nil**
//chInt = nil
func getInt(ch chan int) {
	for i := 1; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 800)
	}
	close(ch)
	fmt.Println("getInt Done")
}

func getStr(ch chan string) {
	ArrString := []string{"hi", "tue", "something", "nothing"}
	for _, strok := range ArrString {
		ch <- strok
		time.Sleep(time.Millisecond * 800)
	}
	close(ch)
	fmt.Println("getString Done")
}

func main() {
	chInt := make(chan int, 2)
	go getInt(chInt)

	chStr := make(chan string, 2)
	go getStr(chStr)

	openCh := 2

	for {
		select {
		case num, ok := <-chInt:
			if ok {
				fmt.Println(num)
			} else {
				fmt.Println("int channel has been closed")
				chInt = nil
				openCh--
			}
		case strok, ok := <-chStr:
			if ok {
				fmt.Println(strok)
			} else {
				fmt.Println("string channel has been closed")
				chStr = nil
				openCh--
			}
		default:
			if openCh == 0 {
				fmt.Println("All channels are closed")
				goto allDone
			}

			time.Sleep(time.Millisecond * 500)
		}
	}
allDone:
	fmt.Println("everithing is the end")
}
### Exampl select sent ch<- num
func getStr(ch chan<- string) {
	ArrString := []string{"hi", "tue", "something", "nothing"}
	for _, strok := range ArrString {
		select {
		case ch <- strok:
			fmt.Println("sent string", strok)
		default:
			fmt.Println("Discard string", strok)
			time.Sleep(time.Millisecond * 800)
		}
	}
	close(ch)
	fmt.Println("getString Done")
}
func getInt(ch chan int) {
	for i := 1; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 800)
	}
	close(ch)
	fmt.Println("getInt Done")
}

func main() {
	chInt := make(chan int, 2)
	go getInt(chInt)

	chStr := make(chan string, 2)
	go getStr(chStr)

	openCh := 2

	for {
		select {
		case num, ok := <-chInt:
			if ok {
				fmt.Println(num)
			} else {
				fmt.Println("int channel has been closed")
				chInt = nil
				openCh--
			}
		case strok, ok := <-chStr:
			if ok {
				fmt.Println(strok)
			} else {
				fmt.Println("string channel has been closed")
				chStr = nil
				openCh--
			}
		default:
			if openCh == 0 {
				fmt.Println("All channels are closed")
				goto allDone
			}

			time.Sleep(time.Millisecond * 500)
		}
	}
allDone:
	fmt.Println("everithing is the end")
}
# 15 err Пропустил
# 16.1 String пакеты strings, unicode, levenshtein
## strings
### strings compare
import (
	"fmt"
	"strings"
)

func packageStrings() {
	product := "Kayak"
	//Эта функция возвращает true, если строка s содержит substr, и false, если нет.
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	//Эта функция возвращает true, если строка s содержит любой из символов,
	//содержащихся в строке substr.
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	//Эта функция возвращает true, если строка s содержит определенную руну (rune).
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	//Эта функция выполняет сравнение без учета регистра и возвращает
	//true, если строки s1 и s2 совпадают.
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	//Эта функция возвращает значение true, если строка s начинается с префикса (prefix) строки.
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	//Эта функция возвращает значение true, если строка заканчивается суффиксом (suffix) строки.
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
}
### strings convert
import (
	"fmt"
	"strings"
)

func Convert() {
	str := "HeLLo"
	//возвращает новую строку, содержащую символы указанной строки, преобразованные в нижний регистр.
	fmt.Println("strings.ToLower ", str, "-> ", strings.ToLower(str))
	//преобразованные в нижний регистр.
	fmt.Println("strings.ToLoUpper ", str, "-> ", strings.ToUpper(str))
	str1 := "hello i am john"
	//первый символ каждого слова был в верхнем регистре, а остальные символы — в нижнем
	fmt.Println("strings.Title ", str1, "-> ", strings.Title(str1))
	//все в верхний регистр
	fmt.Println("strings.ToTitle ", str1, "-> ", strings.ToTitle(str1))
}
### проверка строк
Count(s, sub)
Эта функция возвращает int, которое сообщает, сколько раз указанная подстрока встречается в строке s.
Index(s, sub)
LastIndex(s, sub)
Эти функции возвращают индекс первого или последнего вхождения указанной строки подстроки в строке s или -1, если вхождения нет.
IndexAny(s, chars)
LastIndexAny(s, chars)
Эти функции возвращают первое или последнее вхождение любого символа в указанной строке в пределах строки s или -1, если вхождения нет.
IndexByte(s, b)
LastIndexByte(s, b)
Эти функции возвращают индекс первого или последнего вхождения указанного byte в строке s или -1, если вхождения нет.
IndexFunc(s, func)
LastIndexFunc(s, func)
Эти функции возвращают индекс первого или последнего вхождения символа в строку s, для которого указанная функция возвращает значение true, как описано в разделе «Проверка строк с помощью пользовательских функций».
### strings manipulate
#### Theory
**Fields(s)**
Эта функция разбивает строку на пробельные символы и возвращает срез, содержащий непробельные разделы строки s.
FieldsFunc(s, func)
Эта функция разбивает строку s на символы, для которых пользовательская функция возвращает значение true, и возвращает срез, содержащий оставшиеся части строки.
**Split(s, sub)** - разбивает строку s на срез []string, возвращая string срез. 
SplitN(s, sub, max)
Эта функция похожа на Split, но принимает дополнительный аргумент типа int, указывающий максимальное количество возвращаемых подстрок. Последняя подстрока результирующего среза будет содержать неразделенную часть исходной строки.
SplitAfter(s, sub)
Эта функция похожа на Split, но включает подстроку, используемую в результатах. См. текст после таблицы для демонстрации.
SplitAfterN(s, sub, max)
Эта функция похожа на SplitAfter, но принимает дополнительный аргумент типа int, указывающий максимальное количество возвращаемых подстрок.
#### Example Split(), SplitAfter, Field, TrimSpace
**а если пробелов больше чем один?**
func Manipulate() {
	description := "  A  boat  for  "
	// Разбивает на строки при этом чистит пробелы и пустые строки
	field := strings.Fields(description)
	for i, x := range field {
		field[i] = ">" + x + "<"
	}
	fmt.Println("strings.Fields", field) //[>A< >boat< >for<]

	//Разбивает на строки чистит пробелы
	// - оставляет пустые строки
	splits := strings.Split(description, " ")
	for i, x := range splits {
		splits[i] = ">" + x + "<"
	}
	fmt.Println("strings.Split", splits) //[>< >< >A< >< >boat< >< >for< >< ><]

	//Делит слова но ...
	//Оставляет все пробелы
	splitsAfter := strings.SplitAfter(description, " ")
	for i, x := range splitsAfter {
		splitsAfter[i] = ">" + x + "<"
	}
	fmt.Println("strings.SplitAfter", splitsAfter) //[> < > < >A < > < >boat < > < >for < > < ><]
	//Чистит пробелы перед и после
	//Но не внутри
	trimSpase := strings.TrimSpace(description)
	trimSpase = ">" + trimSpase + "<"
	fmt.Println("strings.TrimSpace", trimSpase) //>A  boat  for<
}
## github.com/agnivade/levenshtein
import (
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
)

func compareStringsLevenshtein(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if distance := levenshtein.ComputeDistance(str1, str2); distance <= 2 {
		return true
	} else {
		return false
	}
}

func main() {
	str1 := "приет"
	str2 := "приве"
	result := compareStringsLevenshtein(str1, str2)
	fmt.Println(result) // Выводит: true
}
## unicode rune (string(str[0])
IsLower(rune)
Эта функция возвращает true, если указанная руна в нижнем регистре.
ToLower(rune)
Эта функция возвращает строчную руну, связанную с указанной руной.
IsUpper(rune)
Эта функция возвращает значение true, если указанная руна написана в верхнем регистре.
ToUpper(rune)
Эта функция возвращает верхнюю руну, связанную с указанной руной.
IsTitle(rune)
Эта функция возвращает true, если указанная руна является заглавной.
ToTitle(rune)
Эта функция возвращает руну в заглавном регистре, связанную с указанной руной.
### example
func Unicode() {
	str := "Ri"
	//Эта функция возвращает true, если указанная руна в нижнем регистре.
	fmt.Println(string(str[0]), unicode.IsLower([]rune(str)[0]), string(str[1]), unicode.IsLower([]rune(str)[1]))
	// Эта функция возвращает строчную руну, связанную с указанной руной.
	fmt.Println(string(str[0]), unicode.ToLower([]rune(str)[0]))
	//Эта функция возвращает значение true, если указанная руна написана в верхнем регистре.
	fmt.Println(string(str[0]), unicode.IsUpper([]rune(str)[0]), string(str[1]), unicode.IsUpper([]rune(str)[1]))
}
## Разделение строк функции
### Theory
TrimSpace(s)
Эта функция возвращает строку s без начальных и конечных пробельных символов.
Trim(s, set)
Эта функция возвращает строку, из которой удаляются все начальные или конечные символы, содержащиеся в наборе (set) строк, из строки s.
TrimLeft(s, set)
Эта функция возвращает строку s без какого-либо начального символа, содержащегося в наборе (set) строк. Эта функция соответствует любому из указанных символов — используйте функцию TrimPrefix для удаления полной подстроки.
TrimRight(s, set)
Эта функция возвращает строку s без каких-либо завершающих символов, содержащихся в наборе (set) строк. Эта функция соответствует любому из указанных символов — используйте функцию TrimSuffix для удаления полной подстроки.
TrimPrefix(s, prefix)
Эта функция возвращает строку s после удаления указанной строки префикса. Эта функция удаляет всю строку префикса (prefix) — используйте функцию TrimLeft для удаления символов из набора.
TrimSuffix(s, suffix)
Эта функция возвращает строку s после удаления указанной строки суффикса (suffix). Эта функция удаляет всю строку суффикса — используйте функцию TrimRight для удаления символов из набора.
TrimFunc(s, func)
Эта функция возвращает строку s, из которой удаляются все начальные или конечные символы, для которых пользовательская функция возвращает значение true.
TrimLeftFunc(s, func)
Эта функция возвращает строку s, из которой удаляются все начальные символы, для которых пользовательская функция возвращает значение true.
TrimRightFunc(s, func)
Эта функция возвращает строку s, из которой удаляются все завершающие символы, для которых пользовательская функция возвращает значение true.
### Example strings.FieldsFunc
func StringsFieldsFunc() {
	description := "This  is  double  spaced"
	splitter := func(r rune) bool {
		return r == ' '
	}
	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}
}
## Обрезка подстрок
func main() {
    description := "A boat for one person"
    prefixTrimmed := strings.TrimPrefix(description, "A boat ")
    wrongPrefix := strings.TrimPrefix(description, "A hat ")
    fmt.Println("Trimmed:", prefixTrimmed)
    fmt.Println("Not trimmed:", wrongPrefix)
}
## Изминение строк (замена строк)
### Theory
Replace(s, old, new, n)
Эта функция изменяет строку s, заменяя вхождения строки old на строку new. Максимальное количество заменяемых вхождений определяется аргументом int n.
ReplaceAll(s, old, new)
Эта функция изменяет строку s, заменяя все вхождения строки old строкой new. В отличие от функции Replace, количество заменяемых вхождений не ограничено.
Map(func, s)
Эта функция генерирует строку, вызывая пользовательскую функцию для каждого символа в строке s и объединяя результаты. Если функция выдает отрицательное значение, текущий символ отбрасывается без замены.
**metods**
Replace(s)
Этот метод возвращает строку, для которой все замены, указанные в конструкторе, были выполнены в строке s.
WriteString(writer, s)
Этот метод используется для выполнения замен, указанных в конструкторе, и записи результатов в io.Writer, описанный в главе 20
### Example Replace
func Replacer() {
	text := "It was a boat. A small boat."
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge", "It", "He")
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced) //He was a kayak. A huge kayak.
}
### replase func(){}()
//убираем все пробелы и точки
func() {
	replacer := strings.NewReplacer(" ", "", ".", "")
	words[i] = replacer.Replace(v)
}()
## Построение и генерация строк
Join(slice, sep)
Эта функция объединяет элементы в указанном срезе строки с указанной строкой-разделителем, помещенной между элементами.
Repeat(s, count)
Эта функция генерирует строку, повторяя строку s указанное количество раз.
## Строительные строки
WriteString(s)
Этот метод добавляет строку s к строящейся строке.
WriteRune(r)
Этот метод добавляет символ r к строящейся строке.
WriteByte(b)
Этот метод добавляет байт b к строящейся строке.
String()
Этот метод возвращает строку, созданную компоновщиком.
Reset()
Этот метод сбрасывает строку, созданную построителем.
Len()
Этот метод возвращает количество байтов, используемых для хранения строки, созданной компоновщиком.
Cap()
Этот метод возвращает количество байтов, выделенных компоновщиком.
Grow(size)
Этот метод увеличивает количество байтов, используемых компоновщиком для хранения строящейся строки.
# 16.2 string regexp
## Регулярные выражения
### Что такое регулярные выражения от GPT
Регулярные выражения (Regular Expressions) - это мощный инструмент для работы с текстом. Они представляют собой шаблоны или строки символов, которые используются для поиска и сопоставления определенных паттернов в тексте.

Регулярные выражения могут быть использованы для различных задач, таких как:

Поиск: Они позволяют найти все вхождения определенного паттерна в тексте. Например, можно найти все email-адреса или ссылки в тексте.

Валидация: Регулярные выражения могут использоваться для проверки, соответствует ли строка определенным правилам. Например, можно проверить, является ли строка допустимым номером телефона или почтовым индексом.

Замена: Регулярные выражения позволяют заменять определенные паттерны в тексте на другие строки. Например, можно заменить все вхождения одного слова на другое.

Разделение: Можно использовать регулярные выражения для разделения строки на подстроки на основе определенного разделителя или паттерна.
### теория
#### Func
Match(pattern, b)
Эта функция возвращает bool значение, указывающее, соответствует ли шаблон байтовому срезу b.
MatchString(patten, s)
Эта функция возвращает bool значение, указывающее, соответствует ли шаблон строке s.
Compile(pattern)
Эта функция возвращает RegExp, который можно использовать для повторного сопоставления с указанным шаблоном, как описано в разделе «Компиляция и повторное использование шаблонов».
MustCompile(pattern)
Эта функция предоставляет те же возможности, что и Compile, но вызывает панику, как описано в главе 15, если указанный шаблон не может быть скомпилирован.
#### metods
MatchString(s)
Этот метод возвращает true, если строка s соответствует скомпилированному шаблону.
FindStringIndex(s)
Этот метод возвращает int срез, содержащий расположение самого левого совпадения, сделанного скомпилированным шаблоном в строке s. Результат nil означает, что совпадений не было.
FindAllStringIndex(s, max)
Этот метод возвращает срез int срезов, содержащих расположение всех совпадений, сделанных скомпилированным шаблоном в строке s. Результат nil означает, что совпадений не было.
FindString(s)
Этот метод возвращает строку, содержащую самое левое совпадение, сделанное скомпилированным шаблоном в строке s. Пустая строка будет возвращена, если совпадений нет.
FindAllString(s, max)
Этот метод возвращает срез строки, содержащий совпадения, сделанные скомпилированным шаблоном в строке s. Аргумент int max указывает максимальное количество совпадений, а -1 указывает отсутствие ограничения. Если совпадений нет, возвращается nil результат.
Split(s, max)
Этот метод разбивает строку s, используя совпадения из скомпилированного шаблона в качестве разделителей, и возвращает срез, содержащий разделенные подстроки.
### Example MatchString
func MatchString() {
	description := "A boat for one person"
	// Ищем oat в тексте
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match) //true
	} else {
		fmt.Println("Error:", err)
	}
}
### Compile - теперь можно юзать методы
func Compile() {
	//Задали паттерн "[A-z]oat"
	pattern, compileErr := regexp.Compile("[A-z]oat")
	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
}
### MustCompile + metod
func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func CompileMetod() {
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
		"=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description, idx))
	}
}
### MustCompile and return string
func MustCompiler() {
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstMatch := pattern.FindString(description)
	allMatches := pattern.FindAllString(description, -1)
	fmt.Println("First match:", firstMatch)
	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}
}
### MustCompAnd and regular expression возвращает string от точки до точки
func MustCompileAndRegExp() {
	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
	description := "Kayak. A boat for one person."
	str := pattern.FindString(description)
	fmt.Println("Match:", str)
}
## Подвыражения regexp.MustCompile
### Theory
FindStringSubmatch(s)
Этот метод возвращает срез, содержащий первое совпадение, сделанное шаблоном, и текст для подвыражений, определяемых шаблоном.
FindAllStringSubmatch(s, max)
Этот метод возвращает срез, содержащий все совпадения и текст подвыражений. Аргумент int используется для указания максимального количества совпадений. Значение -1 указывает все совпадения.
FindStringSubmatchIndex(s)
Этот метод эквивалентен FindStringSubmatch, но возвращает индексы, а не подстроки.
FindAllStringSubmatchIndex(s, max)
Этот метод эквивалентен FindAllStringSubmatch, но возвращает индексы, а не подстроки.
NumSubexp()
Этот метод возвращает количество подвыражений.
SubexpIndex(name)
Этот метод возвращает индекс подвыражения с указанным именем или -1, если такого подвыражения нет.
SubexpNames()
Этот метод возвращает имена подвыражений, выраженные в том порядке, в котором они определены.
### Example
func main() {
    pattern := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
    description := "Kayak. A boat for one person."
    subs := pattern.FindStringSubmatch(description)
    for _, s := range subs {
        fmt.Println("Match:", s)
    }
}
### Именованые подвыражения с подвыподвертом
func CompAndPodv() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	subs := pattern.FindStringSubmatch(description)
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}
}
## Замена подстрок и регулярные выражения
### Theory
ReplaceAllString(s, template)
Этот метод заменяет совпадающую часть строки s указанным шаблоном, который расширяется перед включением в результат для включения подвыражений.
ReplaceAllLiteralString(s, sub)
Этот метод заменяет совпадающую часть строки s указанным содержимым, которое включается в результат без расширения для подвыражений.
ReplaceAllStringFunc(s, func)
Этот метод заменяет совпадающую часть строки s результатом, полученным указанной функцией..
### Example
func Replace() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)
}
### Example and func
func ReplaceFunc() {
	pattern := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	replaced := pattern.ReplaceAllStringFunc(description, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println(replaced)
}
# 17 string fmt. Форматирование и сканирование строк
## Theory Print
Print(...vals)
Эта функция принимает переменное количество аргументов и выводит их значения на стандартный вывод. Пробелы добавляются между значениями, которые не являются строками.
Println(...vals)
Эта функция принимает переменное количество аргументов и выводит их значения на стандартный вывод, разделенные пробелами и сопровождаемые символом новой строки.
Fprint(writer, ...vals)
Эта функция записывает переменное количество аргументов в указанный модуль записи, который я описываю в главе 20. Между значениями, не являющимися строками, добавляются пробелы.
Fprintln(writer, ...vals)
Эта функция записывает переменное количество аргументов в указанный модуль записи, который я описываю в главе 20, за которым следует символ новой строки. Между всеми значениями добавляются пробелы.
### Пробелы 
func main() {
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
	fmt.Printf("Product: %v, Price: $%4.2f\n", Kayak.Name, Kayak.Price)
}
Product: Kayak Price: 275
Product:KayakPrice:275
Product: Kayak, Price: $275.00
## Theory Sprint, Fprint, Errorf
Sprintf(t, ...vals)
Эта функция возвращает строку, созданную путем обработки шаблона t. Остальные аргументы используются в качестве значений для глаголов шаблона.
Printf(t, ...vals)
Эта функция создает строку, обрабатывая шаблон t. Остальные аргументы используются в качестве значений для глаголов шаблона. Строка записывается на стандартный вывод.
Fprintf(writer, t, ...vals)
Эта функция создает строку, обрабатывая шаблон t. Остальные аргументы используются в качестве значений для глаголов шаблона. Строка записывается в модуль Writer, который описан в главе 20.
Errorf(t, ...values)
Эта функция создает ошибку, обрабатывая шаблон t. Остальные аргументы используются в качестве значений для глаголов шаблона. Результатом является значение error, метод Error которого возвращает отформатированную строку.
### errorf
func main() {
	err := fmt.Errorf("Error for index %v", 10)
	fmt.Println(err.Error())
}
## Theory %v, %#v, %T Глаголы форматирования
**%v** - Эта команда отображает формат значения по умолчанию. Изменение глагола со знаком плюс (%+v) включает имена полей при записи значений структуры.
**%#v** Эта команда отображает значение в формате, который можно использовать для повторного создания значения в файле кода Go.
**%T** Эта команда отображает тип значения Go.
### Exemple
func main() {
	var str string = "StringA"
	fmt.Printf("Value %v, Go syntax %#v, Type %T\n", str, str, str)
	//Value StringA, Go syntax "StringA", Type string
}
### Example %+v created my own verb
type Book struct {
	name  string
	pages int
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	book1 := Book{name: "Shopin", pages: 555}

	Printfln("Val and ...%+v", book1) //Val and ...{name:Shopin pages:555}
}
## fmt.Sprintf and format verb
type Book struct {
	Name  string
	Price float64
}

func (p Book) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}

func main() {
	book1 := Book{Name: "Shopin", Price: 55.5}

	book1Info := book1.String()
	fmt.Println(book1Info) //Product: Shopin, Price: $55.50
}
## %b , %d, %o, %x verb
%b
Эта команда отображает целочисленное значение в виде двоичной строки.
%d
Эта команда отображает целочисленное значение в виде десятичной строки. Это формат по умолчанию для целочисленных значений, применяемый при использовании глагола %v.
%o, %O
Эти команды отображают целочисленное значение в виде восьмеричной строки. Глагол %O добавляет префикс 0o.
%x, %X
Эти команды отображают целочисленное значение в виде шестнадцатеричной строки. Буквы от A до F отображаются в нижнем регистре с помощью глагола %x и в верхнем регистре с помощью глагола %X.
### Example
func main() {
    number := 250
    Printfln("Binary: %b", number)
    Printfln("Decimal: %d", number)
    Printfln("Octal: %o, %O", number, number)
    Printfln("Hexadecimal: %x, %X", number, number)
}
## %b, %e, %E, %f, %F, %g, %G ,%x, %X verb
%b
Эта команда отображает значение с плавающей запятой с показателем степени и без десятичной точки.
%e, %E
Эти команды отображают значение с плавающей запятой с показателем степени и десятичным разрядом. %e использует индикатор степени в нижнем регистре, а %E использует индикатор в верхнем регистре.
%f, %F
Эти команды отображают значение с плавающей запятой с десятичным разрядом, но без экспоненты. Команды %f и %F производят одинаковый результат.
%g
Этот глагол адаптируется к отображаемому значению. Формат %e используется для значений с большими показателями степени, в противном случае используется формат %f. Это формат по умолчанию, применяемый при использовании глагола %v.
%G
Этот глагол адаптируется к отображаемому значению. Формат %E используется для значений с большими показателями степени, в противном случае используется формат %f.
%x, %X
Эти команды отображают значение с плавающей запятой в шестнадцатеричном представлении со строчными (%x) или прописными (%X) буквами.
### Example
func main() {
    number := 279.00
    Printfln("Decimalless with exponent: %b", number)
    Printfln("Decimal with exponent: %e", number)
    Printfln("Decimal without exponent: %f", number)
    Printfln("Hexadecimal: %x, %X", number, number)
}
### Example Форматом значений с плавающей запятой можно управлять
func main() {
    number := 279.00
    Printfln("Decimal without exponent: >>%8.2f<<", number)
}
## + 0 - Модификаторы
```
+
Этот модификатор (знак плюс) всегда печатает знак, положительный или отрицательный, для числовых значений.
0
Этот модификатор использует нули, а не пробелы, в качестве заполнения, когда ширина превышает количество символов, необходимое для отображения значения.
-
Этот модификатор (символ вычитания) добавляет отступ справа от числа, а не слева.
```
### Example
func main() {
    number := 279.00
    Printfln("Sign: >>%+.2f<<", number)
    Printfln("Zeros for Padding: >>%010.2f<<", number)
    Printfln("Right Padding: >>%-8.2f<<", number)
}
## rune verb
%s
Этот глагол отображает строку. Это формат по умолчанию, применяемый при использовании глагола %v.
%c
Этот глагол отображает характер. Необходимо соблюдать осторожность, чтобы избежать разделения строк на отдельные байты, как это объясняется в тексте после таблицы.
%U
Эта команда отображает символ в формате Unicode, так что вывод начинается с U+, за которым следует шестнадцатеричный код символа.
### Example
func main() {
    name := "Kayak"
    Printfln("String: %s", name)
    Printfln("Character: %c", []rune(name)[0])
    Printfln("Unicode: %U", []rune(name)[0])
}
## bool verb
%t
Эта команда форматирует логические значения и отображает значение true или false.
func main() {
    name := "Kayak"
    Printfln("Bool: %t", len(name) > 1)
    Printfln("Bool: %t", len(name) > 100)
}
## *Pointer verb
%p
Эта команда отображает шестнадцатеричное представление места хранения указателя
## fmt.Scan 
### Theory
**Функция Scan считывает строку из стандартного ввода и сканирует ее на наличие значений, разделенных пробелами**
Scan(...vals)
Эта функция считывает текст из стандарта и сохраняет значения, разделенные пробелами, в указанные аргументы. Новые строки обрабатываются как пробелы, и функция читает до тех пор, пока не получит значения для всех своих аргументов. Результатом является количество прочитанных значений и error, описывающая любые проблемы.
Scanln(...vals)
Эта функция работает так же, как Scan, но останавливает чтение, когда встречает символ новой строки.
Scanf(template, ...vals)
Эта функция работает так же, как Scan, но использует строку шаблона для выбора значений из получаемых входных данных.
Fscan(reader, ...vals)
Эта функция считывает значения, разделенные пробелами, из указанного средства чтения, описанного в главе 20. Новые строки обрабатываются как пробелы, и функция возвращает количество прочитанных значений и ошибку, описывающую любые проблемы.
Fscanln(reader, ...vals)
Эта функция работает так же, как Fscan, но останавливает чтение, когда встречает символ новой строки.
Fscanf(reader, template, ...vals)
Эта функция работает так же, как Fscan, но использует шаблон для выбора значений из получаемых входных данных.
Sscan(str, ...vals)
Эта функция просматривает указанную строку в поисках значений, разделенных пробелами, которые присваиваются остальным аргументам. Результатом является количество просканированных значений и ошибка, описывающая любые проблемы.
Sscanf(str, template, ...vals)
Эта функция работает так же, как Sscan, но использует шаблон для выбора значений из строки.
Sscanln(str, template, ...vals)
Эта функция работает так же, как Sscanf, но останавливает сканирование строки, как только встречается символ новой строки.
### Example scan
func main() {
    var name string
    var category string
    var price float64
    fmt.Print("Enter text to scan: ")
    n, err := fmt.Scan(&name, &category, &price)
    if (err == nil) {
        Printfln("Scanned %v values", n)
        Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
    } else {
        Printfln("Error: %v", err.Error())
    }
}
### Example Сканирование в срез
var in int
	fmt.Print("Количество елементов среза: ")
	fmt.Scan(&in)

	vals := make([]string, in)
	ivals := make([]interface{}, in)
	for i := 0; i < len(vals); i++ {
		ivals[i] = &vals[i]
	}
	fmt.Print("Enter text to scan: ")
	fmt.Scan(ivals...)
	Printfln("Name: %v", vals)
### fmt.Sscan Сканирует в переменные со строки
func SSscan() {
	source := "Lifejacket Watersports 48.95"
	var name, category string
	var price float64
	//Сканирует в переменные со строки
	//разделитель пробел
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		fmt.Printf("Scanned %v values\n", n)
		fmt.Printf("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		fmt.Printf("Error: %v", err.Error())
	}
}
# 18 Math, sort, rand.
## Printfln
func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
## Theory float64
Abs(val)
Эта функция возвращает абсолютное значение значения float64, то есть расстояние от нуля без учета направления.
Ceil(val)
Эта функция возвращает наименьшее целое число, равное или превышающее указанное значение float64. Результатом также является значение float64, хотя оно представляет собой целое число.
Copysign(x, y)
Эта функция возвращает значение float64, которое является абсолютным значением x со знаком y.
Floor(val)
Эта функция возвращает наибольшее целое число, которое меньше или равно указанному значению float64. Результатом также является значение float64, хотя оно представляет собой целое число.
Max(x, y)
Эта функция возвращает самое большое из указанных значений float64.
Min(x, y)
Эта функция возвращает наименьшее из указанных значений float64.
Mod(x, y)
Эта функция возвращает остаток x/y.
Pow(x, y)
Эта функция возвращает значение x, возведенное в степень y.
Round(val)
Эта функция округляет указанное значение до ближайшего целого числа, округляя половинные значения в большую сторону. Результатом является значение float64, хотя оно представляет собой целое число.
RoundToEven(val)
Эта функция округляет указанное значение до ближайшего целого числа, округляя половинные значения до ближайшего четного числа. Результатом является значение float64, хотя оно представляет собой целое число.
### Example float64
func main() {
	val1 := 279.00
	val2 := 48.95
	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))
}
## Theory int float and uint
MaxInt8
MinInt8
Эти константы представляют наибольшее и наименьшее значения, которые могут быть сохранены с использованием int8.
MaxInt16
MinInt16
Эти константы представляют наибольшее и наименьшее значения, которые могут быть сохранены с использованием типа int16.
MaxInt32
MinInt32
Эти константы представляют наибольшее и наименьшее значения, которые могут быть сохранены с помощью int32.
MaxInt64
MinInt64
Эти константы представляют наибольшее и наименьшее значения, которые могут быть сохранены с помощью int64.
MaxUint8
Эта константа представляет наибольшее значение, которое может быть представлено с помощью uint8. Наименьшее значение равно нулю.
MaxUint16
Эта константа представляет наибольшее значение, которое может быть представлено с помощью uint16. Наименьшее значение равно нулю.
MaxUint32
Эта константа представляет наибольшее значение, которое может быть представлено с помощью uint32. Наименьшее значение равно нулю.
MaxUint64
Эта константа представляет наибольшее значение, которое может быть представлено с помощью uint64. Наименьшее значение равно
MaxFloat32
MaxFloat64
Эти константы представляют самые большие значения, которые могут быть представлены с использованием значений float32 и float64.
SmallestNonzeroFloat32
SmallestNonzeroFloat64
Эти константы представляют наименьшие ненулевые значения, которые могут быть представлены с использованием значений float32 и float64.
## Theory math/rand
**//rand.Seed(time.Now().UnixNano()) - уже не нужен**
Seed(s)
Эта функция устанавливает начальное значение, используя указанное значение int64.
Float32()
Эта функция генерирует случайное значение float32 в диапазоне от 0 до 1.
Float64()
Эта функция генерирует случайное значение float64 в диапазоне от 0 до 1.
Int()
Эта функция генерирует случайное int значение.
Intn(max)
Эта функция генерирует случайное int число меньше указанного значения, как описано после таблицы.
UInt32()
Эта функция генерирует случайное значение uint32.
UInt64()
Эта функция генерирует случайное значение uint64.
Shuffle(count, func)
Эта функция используется для рандомизации порядка элементов, как описано после таблицы.
### Example rand.Int()
import "math/rand"

func main() {
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rand.Int())
	}
}

// Value 0 : 255403926430401497
// Value 1 : 3354664602058181095
// Value 2 : 1488680173336863149
// Value 3 : 6298702337299772946
// Value 4 : 4159130017212465651
### Example rand в диапазоне rand.Intn(10)
func main() {
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rand.Intn(10))
	}
}

// Value 0 : 5
// Value 1 : 5
// Value 2 : 9
// Value 3 : 9
// Value 4 : 5
### Example 3 Как сдвинуть получаемые значения
func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, IntRange(10, 20))
	}
}
### Example 4 перемешать []string{} rand.Shuffle
var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {
	//rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}
}
## Sort []string []int []float
### Theory
Float64s(slice)
Эта функция сортирует срез значений float64. Элементы сортируются на месте.
Float64sAreSorted(slice)
Эта функция возвращает значение true, если элементы в указанном срезе float64 упорядочены.
Ints(slice)
Эта функция сортирует срез значений int. Элементы сортируются на месте.
IntsAreSorted(slice)
Эта функция возвращает значение true, если элементы в указанном int срезе упорядочены.
Strings(slice)
Эта функция сортирует срез string значений. Элементы сортируются на месте.
StringsAreSorted(slice)
Эта функция возвращает значение true, если элементы в указанном срезе string упорядочены.
### Example 1 sort без перемещения  sort.Ints(ints)
func main() {
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints) //Ints Sorted: [-1 2 4 9 10]
	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats) //Floats Sorted: [19.5 48.95 279]
	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings) //Strings Already Sorted: [Kayak Lifejacket Stadium]
	}
}
### example 2 newSlise := create, copy, sort
func main() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
}
## Поиск отсортированных данных
### Theory
SearchInts(slice, val)
Эта функция ищет в отсортированном срезе указанное значение int. Результатом является индекс указанного значения или, если значение не найдено, индекс, по которому значение может быть вставлено при сохранении порядка сортировки.
SearchFloat64s(slice, val)
Эта функция ищет в отсортированном срезе указанное значение float64. Результатом является индекс указанного значения или, если значение не найдено, индекс, по которому значение может быть вставлено при сохранении порядка сортировки.
SearchStrings(slice, val)
Эта функция ищет в отсортированном срезе указанное string значение. Результатом является индекс указанного значения или, если значение не найдено, индекс, по которому значение может быть вставлено при сохранении порядка сортировки.
Search(count, testFunc)
Эта функция вызывает тестовую функцию для указанного количества элементов. Результатом является индекс, для которого функция возвращает значение true. Если совпадений нет, результатом является индекс, в который можно вставить указанное значение для сохранения порядка сортировки.
### Example
func main() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts) //Ints Sorted: [-1 2 4 9 10]
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v", indexOf4) //Index of 4: 2
	//Цыфры 3 в срезе нет но её можно поставить на второе место
	Printfln("Index of 3: %v", indexOf3) //Index of 3: 2
	//дополнительная проверка на наличие в срезе
	Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4) //(present: true)
	Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3) //(present: false)
}
## Сортировка типов данных
### Theory
Len()
Этот метод возвращает количество элементов, которые будут отсортированы.
Less(i, j)
Этот метод возвращает значение true, если элемент с индексом i должен появиться в отсортированной последовательности перед элементом j. Если Less(i,j) и Less(j, i) оба false, то элементы считаются равными.
Swap(i, j)
Этот метод меняет местами элементы по указанным индексам.
**--------**
Sort(data)
Эта функция использует методы, описанные в таблице 18-8, для сортировки указанных данных.
Stable(data)
Эта функция использует методы, описанные в таблице 18-8, для сортировки указанных данных без изменения порядка элементов с одинаковым значением.
IsSorted(data)
Эта функция возвращает значение true, если данные отсортированы.
Reverse(data)
Эта функция меняет порядок данных.
### Example 1
//productsort.go
import "sort"

type Product struct {
	Name  string
	Price float64
}
type ProductSlice []Product

func ProductSlices(p []Product) {
	sort.Sort(ProductSlice(p))
}
func ProductSlicesAreSorted(p []Product) {
	sort.IsSorted(ProductSlice(p))
}
func (products ProductSlice) Len() int {
	return len(products)
}
func (products ProductSlice) Less(i, j int) bool {
	return products[i].Price < products[j].Price
}
func (products ProductSlice) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}
// main.go
func main() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlices(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
### Example 2 sortName struct
type Product struct {
	Name  string
	Price float64
}
type ProductSlice []Product

func ProductSlices(p []Product) {
	sort.Sort(ProductSlice(p))
}
func ProductSlicesAreSorted(p []Product) {
	sort.IsSorted(ProductSlice(p))
}
func (products ProductSlice) Len() int {
	return len(products)
}
func (products ProductSlice) Less(i, j int) bool {
	return products[i].Price < products[j].Price
}
func (products ProductSlice) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

type ProductSliceName struct{ ProductSlice }

func ProductSlicesByName(p []Product) {
	sort.Sort(ProductSliceName{p})
}
func (p ProductSliceName) Less(i, j int) bool {
	return p.ProductSlice[i].Name < p.ProductSlice[j].Name
}

func main() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	//ProductSlices(products)
	ProductSlicesByName(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
### Сравнение
//productsort.go
import "sort"

type Product struct {
	Name  string
	Price float64
}
type ProductSlice []Product

func ProductSlices(p []Product) {
	sort.Sort(ProductSlice(p))
}

func ProductSlicesAreSorted(p []Product) {
	sort.IsSorted(ProductSlice(p))
}

func (products ProductSlice) Len() int {
	return len(products)
}

func (products ProductSlice) Less(i, j int) bool {
	return products[i].Price < products[j].Price
}

func (products ProductSlice) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

type ProductSliceName struct{ ProductSlice }

func ProductSlicesByName(p []Product) {
	sort.Sort(ProductSliceName{p})
}

func (p ProductSliceName) Less(i, j int) bool {
	return p.ProductSlice[i].Name < p.ProductSlice[j].Name
}

type ProductComparison func(p1, p2 Product) bool

type ProductSliceFlex struct {
	ProductSlice
	ProductComparison
}

func (flex ProductSliceFlex) Less(i, j int) bool {
	return flex.ProductComparison(flex.ProductSlice[i], flex.ProductSlice[j])
}
func SortWith(prods []Product, f ProductComparison) {
	sort.Sort(ProductSliceFlex{prods, f})
}
**----**
func main() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	//ProductSlices(products)
	//ProductSlicesByName(products)
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
# 19 Даты, время и продолжительность time
## Theory and practice
### func Time
Now()
Эта функция создает Time, представляющий текущий момент времени.
Date(y, m, d, h, min, sec, nsec, loc)
Эта функция создает объект Time, представляющий указанный момент времени, который выражается аргументами года, месяца, дня, часа, минуты, секунды, наносекунды и Location. (Тип Location описан в разделе «Синтаксический анализ значений времени из строк».)
Unix(sec, nsec)
Эта функция создает значение Time из числа секунд и наносекунд с 1 января 1970 года по Гринвичу, широко известного как время Unix.
### metods time
Date()
Этот метод возвращает компоненты года, месяца и дня. Год и день выражаются как значения int, а месяц — как значение Month.
Clock()
Этот метод возвращает компоненты часа, минут и секунд Time.
Year()
Этот метод возвращает компонент года, выраженный как int.
YearDay()
Этот метод возвращает день года, выраженный как int от 1 до 366 (для учета високосных лет).
Month()
Этот метод возвращает компонент месяца, выраженный с использованием типа Month.
Day()
Этот метод возвращает день месяца, выраженный как int.
Weekday()
Этот метод возвращает день недели, выраженный как Weekday.
Hour()
Этот метод возвращает час дня, выраженный как int от 0 до 23.
Minute()
Этот метод возвращает количество минут, прошедших до часа дня, выраженное как int от 0 до 59.
Second()
Этот метод возвращает количество секунд, прошедших до минуты часа, выраженное как int от 0 до 59.
Nanosecond()
Этот метод возвращает количество наносекунд, прошедших до секунды минуты, выраженное как int от 0 до 999 999 999.
### type Time
Month
Этот тип представляет месяц, а пакет time определяет постоянные значения для названий месяцев на английском языке: January, February и т. д. Тип Month определяет метод String, который использует эти имена при форматировании строк.
Weekday
Этот тип представляет день недели, а пакет time определяет постоянные значения для названий дней недели на английском языке: Sunday, Monday и т. д. Тип Weekday определяет метод String, который использует эти имена при форматировании строк.
### Example 1 PrintTime time.Now and
func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v: Month: %v Year: %v",
		label, t.Day(), t.Month(), t.Year())
}
func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	//or
	fmt.Println(current.Day(), current.Month(), current.Year())
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}
### Exemple Format(layout)
//Format(layout) - Этот метод возвращает отформатированную строку, созданную с использованием указанного макета.

func PrintTime(label string, t *time.Time) {
	layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(layout))
}
func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
	fmt.Println(current.Day(), current.Month(), current.Year())
}
### Константы компоновки
#### Theory
ANSIC
Mon Jan _2 15:04:05 2006
UnixDate
Mon Jan _2 15:04:05 MST 2006
RubyDate
Mon Jan 02 15:04:05 -0700 2006
RFC822
02 Jan 06 15:04 MST
RFC822Z
02 Jan 06 15:04 -0700
RFC850
Monday, 02-Jan-06 15:04:05 MST
RFC1123
Mon, 02 Jan 2006 15:04:05 MST
RFC1123Z
Mon, 02 Jan 2006 15:04:05 -0700
RFC3339
2006-01-02T15:04:05Z07:00
RFC3339Nano
2006-01-02T15:04:05.999999999Z07:00
Kitchen
3:04PM
Stamp
Jan _2 15:04:05
StampMilli
Jan _2 15:04:05.000
StampMicro
Jan _2 15:04:05.000000
StampNano
Jan _2 15:04:05.000000000
#### Exemple
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
	fmt.Println(current.Day(), current.Month(), current.Year())
}
### Разбор значений времени из строк
#### Theory
Parse(layout, str)
Эта функция анализирует строку, используя указанный макет, чтобы создать значение Time. Возвращается error, указывающая на проблемы с разбором строки.
ParseInLocation(layout, str, location)
Эта функция анализирует строку, используя указанный макет и Location, если в строку не включен часовой пояс. Возвращается error, указывающая на проблемы с разбором строки.
#### Exemple
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}
### предопределенных макетов даты
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	dates := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
	}
	for _, d := range dates {
		time, err := time.Parse(time.RFC822, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}
### Указание разбора местоположения
#### Theory
LoadLocation(name)
Эта функция возвращает *Location для указанного имени и error, указывающую на наличие проблем.
LoadLocationFromTZData(name, data)
Эта функция возвращает *Location из байтового среза, содержащего отформатированную базу данных часовых поясов.
FixedZone(name, offset)
Эта функция возвращает *Location, который всегда использует указанное имя и смещение от UTC.
#### Exemple
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
}
### Встраивание базы данных часовых поясов
База данных часовых поясов, используемая для создания значений Location, устанавливается вместе с инструментами Go, что означает, что она может быть недоступна при развертывании скомпилированного приложения. Пакет time/tzdata содержит встроенную версию базы данных, загружаемую функцией инициализации пакета (как описано в главе 12). Чтобы гарантировать, что данные часового пояса всегда будут доступны, объявите зависимость от пакета следующим образом:
...
import (
    "fmt"
    "time"
       _ "time/tzdata"
)
...
В пакете нет экспортированных функций, поэтому пустой идентификатор должен использоваться для объявления зависимости без создания ошибки компилятора.
### Использование локального местоположения
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		localTime, _ := time.ParseInLocation(layout, date, local)
		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
}
### Непосредственное указание часовых поясов
func PrintTime(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}
func main() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london := time.FixedZone("BST", 1*60*60)
	newyork := time.FixedZone("EDT", -4*60*60)
	local := time.FixedZone("Local", 0)
	//if lonerr == nil && nycerr == nil {
	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)
	PrintTime("No location:", &nolocation)
	PrintTime("London:", &londonTime)
	PrintTime("New York:", &newyorkTime)
	PrintTime("Local:", &localTime)
	// } else {
	// 	fmt.Println(lonerr.Error(), nycerr.Error())
	// }
}
## Управление значениями времени time.Duration
### Theory 
Add(duration)
Этот метод добавляет указанную Duration к Time и возвращает результат.
Sub(time)
Этот метод возвращает значение Duration, выражающее разницу между Time вызова метода и Time, указанным в качестве аргумента.
AddDate(y, m, d)
Этот метод добавляет к Time указанное количество лет, месяцев и дней и возвращает результат.
After(time)
Этот метод возвращает значение true, если Time, в которое был вызван метод, наступает после Time, указанного в качестве аргумента.
Before(time)
Этот метод возвращает значение true, если время, в которое был вызван метод, предшествует Time, указанному в качестве аргумента.
Equal(time)
Этот метод возвращает значение true, если Time, в которое был вызван метод, равно Time, указанному в качестве аргумента.
IsZero()
Этот метод возвращает значение true, если Time, в которое был вызван метод, представляет момент нулевого времени, то есть 1 января 1 года, 00:00:00 UTC.
In(loc)
Этот метод возвращает значение Time, выраженное в указанном Location.
Location()
Этот метод возвращает Location, связанный с Time, что фактически позволяет выразить время в другом часовом поясе.
Round(duration)
Этот метод округляет Time до ближайшего интервала, представленного значением Duration.
Truncate(duration)
Этот метод округляет Time до ближайшего интервала, представленного значением Duration.
### Practice
#### Exemple 1 
func main() {
	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if err == nil {
		Printfln("After: %v", t.After(time.Now()))
		Printfln("Round: %v", t.Round(time.Hour))
		Printfln("Truncate: %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}
}
#### Exemple 2 Equal
func main() {
	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)
}
## Представление продолжительности time.Duration part2
### Theory
**func**
Hour
Эта константа представляет 1 час.
Minute
Эта константа представляет 1 минуту.
Second
Эта константа представляет 1 секунду.
Millisecond
Эта константа представляет 1 миллисекунду.
Microsecond
Эта константа представляет 1 миллисекунду.
Nanosecond
Эта константа представляет 1 наносекунду.
**Методы продолжительности**
Hours()
Этот метод возвращает float64, который представляет Duration в часах.
Minutes()
Этот метод возвращает float64, который представляет Duration в минутах.
Seconds()
Этот метод возвращает float64, который представляет Duration в секундах.
Milliseconds()
Этот метод возвращает float64, который представляет Duration в миллисекундах.
Microseconds()
Этот метод возвращает float64, который представляет Duration в микросекундах.
Nanoseconds()
Этот метод возвращает float64, который представляет Duration в наносекундах.
Round(duration)
Этот метод возвращает Duration, которая округляется до ближайшего кратного указанной Duration.
Truncate(duration)
Этот метод возвращает Duration, которая округляется в меньшую сторону до ближайшего кратного указанной Duration.
### Example 1
func main() {
	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())
}
## Создание продолжительности относительно времени
### Theory
Since(time)
Эта функция возвращает Duration, выражающую время, прошедшее с момента указанного значения Time.
Until(time)
Эта функция возвращает Duration, выражающую время, прошедшее до указанного значения Time.
### Example 
func main() {
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))
}
## Создание длительности из строк ParseDuration(str)
### Theory
ParseDuration(str)
Эта функция возвращает значение Duration и error, указывающую на наличие проблем при синтаксическом анализе указанной строки.
h
Эта единица обозначает часы.
m
Эта единица обозначает минуты.
s
Эта единица обозначает секунды.
ms
Эта единица обозначает миллисекунды.
us или μs
Эта единица обозначает микросекунды.
ns
Эта единица обозначает наносекунды.
### Example
func main() {
	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}
## Использование функций времени для горутин и каналов
### Theory
Sleep(duration)
Эта функция приостанавливает текущую горутину по крайней мере на указанное время.
AfterFunc(duration, func)
Эта функция выполняет указанную функцию в своей собственной горутине по истечении указанного времени. Результатом является *Timer, метод Stop которого можно использовать для отмены выполнения функции до истечения продолжительности.
After(duration)
Эта функция возвращает канал, который блокируется на указанное время, а затем возвращает значение Time. Подробнее см. в разделе «Получение уведомлений по времени».
Tick(duration)
Эта функция возвращает канал, который периодически отправляет значение Time, где период указан как продолжительность.
### Example 
func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
### Example 2 Отсрочка выполнения функции time.AfterFunc
func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 1)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	time.AfterFunc(time.Second*5, func() { writeToChannel(nameChannel) })
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
## Получение уведомлений по времени
func writeToChannel(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
## Использование уведомлений в качестве тайм-аутов в операторах Select
func writeToChannel(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 3)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	channelOpen := true
	for channelOpen {
		Printfln("Starting channel read")
		select {
		case name, ok := <-nameChannel:
			if !ok {
				channelOpen = false
				break
			} else {
				Printfln("Read name: %v", name)
			}
		case <-time.After(time.Second * 2):
			Printfln("Timeout")
		}
	}
}
## Остановка и сброс таймеров NewTimer(duration)
### Theory
NewTimer(duration) - Эта функция возвращает *Timer с указанным периодом.
C - Это поле возвращает канал, по которому Time будет отправлять свое значение Time.
Stop() - Этот метод останавливает таймер. Результатом является логическое значение, которое будет true, если таймер был остановлен, и false, если таймер уже отправил свое сообщение.
Reset(duration) - Этот метод останавливает таймер и сбрасывает его так, чтобы его интервал был заданным значением Duration.
### Example
func writeToChannel(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	timer := time.NewTimer(time.Minute * 10)
	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second)
	}()

	Printfln("Initial duration elapsed.")
	<-timer.C
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 3)
	}
	close(channel)
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
## Получение повторяющихся уведомлений time.Tick
### Theory
NewTicker(duration)
Эта функция возвращает *Ticker с указанным периодом.
C
Это поле возвращает канал, по которому Ticker будет отправлять значения Time.
Stop()
Этот метод останавливает тикер (но не закрывает канал, возвращаемый полем C).
Reset(duration)
Этот метод останавливает тикер и сбрасывает его так, чтобы его интервал был равен указанной Duration.
### Example 1
func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	tickChannel := time.Tick(time.Second)
	index := 0
	for {
		<-tickChannel
		nameChannel <- names[index]
		index++
		if index == len(names) {
			index = 0
		}
	}
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
### Example 2
func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	ticker := time.NewTicker(time.Second / 10)
	index := 0
	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
		}
	}
}
func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
# 20 Чтение и запись данных
## Theory Read
Read(byteSlice)
Этот метод считывает данные в указанный []byte. Метод возвращает количество прочитанных байтов, выраженное как int, и error.
### Example
func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	r := strings.NewReader("Kayak Яматоусуки")
	processData(r)
}
## Theory Writer
Write(byteSlice)
Этот метод записывает данные из указанного byte среза. Метод возвращает количество записанных байтов и error. Ошибка будет ненулевой, если количество записанных байтов меньше длины среза.
### Example
func processDataRead(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	r := strings.NewReader("Kayak Яматоусуки")
	//processData(r)
	//
	var builder strings.Builder
	processDataRead(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}
## Theory Read and Write (Copy, ReadAll, WriteString)
Copy(w, r)
Эта функция копирует данные из Reader в Writer до тех пор, пока не будет возвращен EOF или не будет обнаружена другая ошибка. Результатом является количество копий байтов и error, используемая для описания любых проблем.
CopyBuffer(w, r, buffer)
Эта функция выполняет ту же задачу, что и Copy, но считывает данные в указанный буфер перед их передачей во Writer.
CopyN(w, r, count)
Эта функция копирует count байтов из Reader в Writer. Результатом является количество копий байтов и error, используемая для описания любых проблем.
ReadAll(r)
Эта функция считывает данные из указанного Reader до тех пор, пока не будет достигнут EOF. Результатом является байтовый срез, содержащий считанные данные и error, которая используется для описания любых проблем.
ReadAtLeast(r, byteSlice, min)
Эта функция считывает как минимум указанное количество байтов из устройства чтения, помещая их в байтовый срез. Сообщается об ошибке, если считано меньше байтов, чем указано.
ReadFull(r, byteSlice)
Эта функция заполняет указанный байтовый срез данными. Результатом является количество прочитанных байтов и error. Будет сообщено об ошибке, если EOF будет обнаружен до того, как будет прочитано достаточно байтов для заполнения среза.
WriteString(w, str)
Эта функция записывает указанную строку в модуль записи.
### Example copy
func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
func main() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}
## специализированных средств чтения и записи
Pipe()
Эта функция возвращает PipeReader и PipeWriter, которые можно использовать для соединения функций, требующих Reader и Writer, как описано в разделе «Использование каналов».
MultiReader(...readers)
Эта функция определяет переменный параметр, который позволяет указать произвольное количество значений Reader. В результате получается Reader, которое передает содержимое каждого из своих параметров в той последовательности, в которой они определены, как описано в разделе «Объединение нескольких средств чтения».
MultiWriter(...writers)
Эта функция определяет переменный параметр, который позволяет указать произвольное количество значений Writer. Результатом является Writer, который отправляет одни и те же данные всем указанным модулям записи, как описано в разделе «Объединение нескольких средств записи».
LimitReader(r, limit)
Эта функция создает Reader, который будет завершать работу после указанного количества байтов, как описано в разделе «Ограничение чтения данных».
### Example Pipe
func GenerateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	writeSize := 4
	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error: %v", err.Error())
		}
	}
}
func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			Printfln("Read data: %v", string(slice[0:count]))
			data = append(data, slice[0:count]...)
		}
		if err == io.EOF {
			break
		}
	}
	Printfln("Read data: %v", string(data))
}

func FinalReader() {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		pipeWriter.Close()
	}()
	ConsumeData(pipeReader)
}
### Example MultyReader
func main() { 

	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)

}
### Объединение нескольких средств записи MultiWriter
func main() {

	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

}
### 














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