# Лекция 1. Установка и запуск первой программы
append
## Шаг 1. Установка компилятора
* Переход по ссылке: https://golang.org/dl/
* Установили для своей ОС компилятор GoLang
***Важно*** : на данном курсе желательно, чтобы у вас была версия компилятра > 1.08

## Шаг 2. GoRoot GoPath
**Определение** : ```GOROOT``` - это файловый путь, указывающий расположение ***КОМПИЛЯТОРА*** Go.
**Определение**: ```GOPATH``` - это файловый путь , указывающий на расположение ***РАБОЧЕГО ОКРУЖЕНИЯ*** (Там где пишем код и мазюкаем проекты). По умолчанию, на курсе мы создали ```GOPATH``` по адресу ```C:\Users\<username>\go```

## Шаг 3. Инициализация рабочего окружения
Чтобы создать рабочее окрудение нам надо в ```GOPATH``` определить 3 диреткории:
* ```src``` - место, где будут лежать исходники проектов (скрипты .go)
* ```bin``` - место, где будут лежать скомпилированные бинарники, после выполнения компиляции проектов
* ```pkg``` - мсто, где будут жить сторонние пакеты для наших проектов

## Шаг 4. Первая программа
В ```GOPATH```/src создадим файл ```main.go```
со следующей начинкой:
```
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}

```

## Шаг 5. Запуск и компиляция
Go - компилируемый язык.
Для того, чтобы скомпилировать исполняемый файл, можно выполнить команду
```
go build <path/to/go/file.go>
```
Данная команда создает исполнямый файл по месту ее вызова. Это удобно,когда мы в пылу битвы и хочется посмотреть, скомпилируется ли оно вообще или в случае каких-либо тестов - позволяет на месте все проверять.

Другая команда, которая также позволяет создать исполняемый файл:
```
go install <path/to/go/file.go>
```
Данная команда создает исполняемый файл по пути ```GOPATH/bin```.

Третья команда, которая будет часто использоваться на курсе - ```go run```
```
go run <path/to/go/file.go>
```
Данная команда делает следующие действия:
* Создает исполняемый файл в временном хранилище
* Запускает этот файл
* И "устарняет его" (зависит ОС)

Для того, чтобы узнать, где это временное хранилище ```--work```

## Шаг 6. Правильная структуризация рабочего окружения
Очень рекомендую создать следующий путь ```GOPATH/src/github.com/<your_github_username>/<github_repo_name>```
# Лекция 2. Декларирование переменных и I/O

## Шаг 1. Какая типизация?
***В языке Go*** принята полустрогая статическая типизация.

## Шаг 2. Способы декларирования переменных
**Декларирование** - это процесс связывания имени переменной с типом потенциального значения
**При декларировании переменной** автоматически происходит ее **инициализация** **НУЛЕВЫМ ЗНАЧЕНИЕМ ДЛЯ ЭТОГО ТИПА***.
Пример:
```
var age int
fmt.Println("My age is:", age)
```
Будет выведен 0.


## Шаг 3. Декларирование и инициализация
Простейший случай единичной инициализации
```
//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем "полустрогость" типзации?
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)
```

## Короткое присваивание
Оператор ```:=``` слева от себя требует ***КАК МИНИМУМ ОДНУ НОВУЮ ПЕРЕМЕННУЮ***.
```
//Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 10, 30
	// fmt.Println(aArg, bArg)

	//Исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg)

```
## fmt.Fscan  fmt.Scan  fmt.Printf fmt.Fscan

func main() {
	var (
		age  int
		name string
	)
	//fmt.Scan (&age) // Функция Сканирует (ищет) данные
	//fmt.Scan (&name)
	fmt.Scan(&age, &name)

	fmt.Printf("My name is : %s\n, My age is : %v\n", name, age)

	fmt.Fscan(os.Stdin, &age)    //Fscan чтение данных из других источников (нужно указывать явно) (os.Stdin - клавиатура)
	fmt.Println("New age:", age) // существует команда Fprint - для отправки вывода на другое устройство
}
# Типы данных 


## Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

## Numerics. Integers
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)

	//Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)
	//Узнаем, сколько байт занимает переемнная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент. При использовании короткого объявления - тип для целого числа - int платформо-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//Эксперимент 2. Используйте явное приведение типов при необходимости если уверены что не произойдет коллиззии
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Эксперимент 3. Если проводятся арифметические операции
	// над int и intX , то обязательно нужно использовать механизм приведения. Т.к. int != int64
	var third64 int64 = 16123414
	var fourthInt int = 156234
	fmt.Println(third64 + int64(fourthInt))
	// + - * / %

	// Аналогичным образом утсроены unit8, uint16, uint32, uint64, uint
	//Numerics. Float
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)

## Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

## Strings. Rune

	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string :", name, len(name)) // Функция len() возвращает количество элементов в наборе
	
        fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))

        //Rune - руна. Это один utf-ный символ.
	
        //Поиск подстроки в строке
	totalString, subString := "ABCDEDFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))
        //из strings v logic 
        fmt.Println(strings.Compare("abcd", "a")) // -1 if first < second, 0 if first == second, 1 if first > second
             //strings.EqualFold(i, I) нашёл в нете
	
        //rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' // Для инициализации руны символьным значением - используйте ''
	var thirdRune rune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)
	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "a")) // -1 if first < second, 0 if first == second, 1 if first > second

	var aByte byte // alias uint8
	fmt.Println("Byte:", aByte)

### String convert

// Получить срез float64 со строки через strings.Split
func ArrFloat(A string, N string) []float64 { //N - количество цифр "N" string
	a := strings.Split(A, " ")
	n, _ := strconv.Atoi(N)
	b := make([]int, n)
	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}
	dd := []float64{}
	for _, v := range b {
		dd = append(dd, float64(v))
	}
	return dd
}

// regexp.MustCompile("[0-9]+")
func ArrFloat2(a string) []float64 {
	re := regexp.MustCompile("[0-9]+")
	s := re.FindAllString(a, -1)
	d := []float64{}
	for _, v := range s {
		sv, _ := strconv.Atoi(v)
		vv := float64(sv)
		d = append(d, vv)
	}
	return d
}

// Сортировка через тип
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	//STRING
	var String string = "qwerty"
	Str := "qwerty"
	AddStr := String + " " + Str
	fmt.Println("Concat", AddStr)

	//Сравнение строк
	fmt.Println(strings.Contains(String, Str)) // true or false
	fmt.Println(strings.Compare("ifgy", "fg")) // 1 if false or 0 if true

	//Convert
	//string to int
	StrInt, _ := strconv.Atoi("1234")
	fmt.Println(StrInt)
	StrInt64, _ := strconv.ParseInt("456", 10, 64)
	fmt.Println(StrInt64)

	//int to string
	Str2 := strconv.Itoa(StrInt)
	fmt.Println(Str2)
	//string to float
	var StrNum string = "25"
	StrFloat64, _ := strconv.ParseFloat(StrNum, 64)
	fmt.Println(StrFloat64)

	// string to rune
	d := []rune(String)
	fmt.Printf("string to rune: %c \n", d)
	// rune to string
	stringD := string(d)
	fmt.Println("rune to string:", stringD)
	// Длина RuneCounter - количество !рун! в СТРОКЕ
	fmt.Println("Length of String:", utf8.RuneCountInString(String), "runes")

	// []string{}
	s := []string{"1234", "54678", "789"}
	sort.Strings(s)        // по алфавиту
	fmt.Println(s)         // [1234 546 789]
	sort.Sort(byLength(s)) // по количеству, через тип
	fmt.Println(s)         //[789 1234 54678]

	// string to []byte{}
	wordString := "Тестовая строка"
	wordByte := []byte(wordString) // преобразование
	fmt.Println(wordByte)
	// []byte{} to string
	newString := string(wordByte)
	fmt.Println(newString)

	// string to float64
	StringIntov := "10 15 25 64 87"
	// Массив флоат со стринги
	Arrfloat64 := ArrFloat(StringIntov, "5")
	fmt.Println(Arrfloat64)
	// Масси флоат со стринги
	Arr2Float64 := ArrFloat2(StringIntov)
	fmt.Println(Arr2Float64)

	// []int{} to string
	SliseInt := []int{23, 5, 16, 32, 77}
	SString := []string{}
	for _, v := range SliseInt {
		d := strconv.Itoa(v)
		SString = append(SString, d)
	}
	var frt string = strings.Join(SString, " ")
	fmt.Println(frt)
}
# Условный оператор  if else

## Классический условный оператор
	// 	if condition {
	// 		//body
	// 	}
	// }
	//	Условный оператор с блоком else
	// if condition {

	// } else {

	// }
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}
	// if condition1 {

	// } else if condition2 {

	// } else if ... {

	// } else {

	// }

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

## Good Инициализация в блоке условного оператора
	//Блок присваивания - только :=
	//Инициализируемая переменная видна ТОЛЬКО внутри области виидимости условного оператора (в телах if, else if, или else)
	// Но не за его пределами
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	//Ущербно
	/*
		var age int = 10
		if age > 7 {
			fmt.Println("Go to school")
		} //По факту, сюда подставляется ; компилятором, и дальнейший код уже не имеет связи с предыдущим if
		else {
			fmt.Println("Another case")
		}
	*/

	//НЕ ИДЕОМАТИЧНО
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("Width <= 100")
	}
	//Странное правило номер 1: в Go стараются избегать блоков ELSE

	//Идеоматичность
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("Height <= 100")
}

## HOME WORK
### Чётные не чётные
//https://contest.yandex.ru/contest/25606/problems/J/
func main() {

	var (
		a float32
		b float32
	)

	fmt.Scan(&a, &b)
	if sum := int64(a + b); sum%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}

}

### strings.Contains
K tekst analiz
func main() {
	var a string
	var b string = "чат"
	fmt.Scan(&a)
	if strings.Contains(a, b) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}

}

### Валидатор при регистрации strings.Contains

func main() {
	var (
		a string
		b string = "@"
		c string
		d string = "."
		e int    = 10
	)
	fmt.Scan(&a)

	if strings.Contains(a, b) || len(a) < e {
		fmt.Println("Некорректный логин")
	} else {
		fmt.Scan(&c)
		if strings.Contains(c, d) || strings.Contains(a, d) {
			fmt.Println("ОК")
		} else {
			fmt.Println("Некорректная почта")
		}
	}

}

### strings.Compare
Наш отряд
func main() {
	var a string
	var b string
	var c string

	fmt.Scan(&a)
	if strings.Compare(a, "раз") == 0 || strings.Compare(a, "один") == 0 {
		fmt.Scan(&b)
		if strings.Compare(b, "два") == 0 {
			fmt.Scan(&c)
			if strings.Compare(c, "три") == 0 {
				fmt.Println("ОК")
			} else {
				fmt.Println("НЕ ПРАВИЛЬНО")
			}
		} else {
			fmt.Println("НЕ ПРАВИЛЬНО")
		}

	} else if strings.Compare(a, "1") == 0 {
		fmt.Scan(&b)
		if strings.Compare(b, "2") == 0 {
			fmt.Scan(&c)
			if strings.Compare(c, "3") == 0 {
				fmt.Println("ОК")
			} else {
				fmt.Println("НЕ ПРАВИЛЬНО")
			}
		} else {
			fmt.Println("НЕ ПРАВИЛЬНО")
		}

	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}

}

### Ход коня 
package main

import "fmt"

func main() {
	var (
		x1 int16
		y1 int16
		x2 int16
		y2 int16
	)
	fmt.Scan(&x1, &y1, &x2, &y2)
	if ((x1 - x2) == 2) && ((y1 - y2) == 1) {
		fmt.Println("ДА")
	} else if ((x1 - x2) == 2) && ((y1 - y2) == -1) {
		fmt.Println("ДА")
	} else if ((x1 - x2) == -1) && ((y1 - y2) == -2) {
		fmt.Println("ДА")
	} else if ((x1 - x2) == -2) && ((y1 - y2) == -1) {
		fmt.Println("ДА")
	} else if ((x1 - x2) == -2) && ((y1 - y2) == 1) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
# Цыклы
## for

func main() {
	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - условие (если верно - то тело цикла выполняется, если нет - то цикл завершается)
	// post - изменение переменной цикла (инкрементарное действие, декрементарное действие)
	// }

	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}
	//Важный момент : в качестве init может быть использовано выражение присваивания ТОЛЬКО через :=

### break - команда, прерывающая текущее выполнение тела цикла и передающая управление инструкциям, следующим за циклом
	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK")

### continue - команда, прерывающая текущее выполнение тела цикла и передающая управления СЛЕДУЮЩЕЙ ИТЕРАЦИИ ЦИКЛА
	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with CONTINUE")

### Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идее выше треугольник")

### Лейблы - это синтаксический сахар. Иногда бывает плохо. С лейблами по лучше. 
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer // Хочу чтобы вообще все циклы (внешние тоже остановились)
			}
		}
	}
	//Модификации цикла for.
### 1. Классчиеский цикл while do
	var loopVar int = 0
	// while (loopVar < 10){
	// 	....
	// 	loopVar++
	// }
	for loopVar < 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}
### 2. Классический бесконечный цикл
	var password string
outer2:
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password . Try again")
		} else {
			fmt.Println("Password Accepted")
			break outer2
		}
	}

### 3. Цикл с множественными переменными цикла
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}
}

## switch-case

func main() {
	//Switch!
	var price int
	fmt.Scan(&price)
	//В switch-case запрещены дублирующиеся условия в case"ах
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")
	default:
		//Отрабатывает только в том случае, если не один из выше перечисленных кейсов - не сработал
		fmt.Println("Default case")

	}

### Case с множеством вариантов
	var ageGroup string = "q" //Возрастнаые группы : "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You are too yong/old")
	}

### Case с выражениями
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("Too yong")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 30 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you")
	}

### Case с проваливаниями. Проваливания выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ
	//В момент выполнения fallthroug у следующего кейса не проверяется условие,
	//а сразу выполняется тело
	var number int
	fmt.Scan(&number)
outer:
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		if number%2 == 0 {
			break outer
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d GREATER then 200\n", number)
		fallthrough
	case number > 1000:
		fmt.Printf("%d GREATER then 1000\n", number)
		fallthrough
	default:
		fmt.Printf("%d DEFAULT\n", number)
	}

### Гадость с терминацией цикла for из switchv
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}

	fmt.Println("END")
}

## Home Work
### Пирамидка for
//https://contest.yandex.ru/contest/25622/problems/K/

//   1
//  1 1
// 1 2 1

func main() {
	var rows int
	temp := 1
	fmt.Scan(&rows)
	if rows > 0 {
		for i := 0; i < rows; i++ { //
			for j := 2; j <= rows-i; j++ { //пробелы
				fmt.Print(" ")
			}
			for k := 0; k <= i; k++ {
				if k == 0 || i == 0 {
					temp = 1
				} else {
					temp = temp * (i - k + 1) / k
				}
				fmt.Printf("%d ", temp)
			}
			fmt.Println("")
		}
	}
}

### Дроби if + math.Modf
////https://contest.yandex.ru/contest/25622/problems/H/

func main() {
	var a float64
	var b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		c, d := math.Modf(a / b)
		if d == 0 {
			fmt.Print(c)
		} else {
			fmt.Printf("%.2f")
		}

	} else if b > a {
		c, d := math.Modf(b / a)
		if d == 0 {
			fmt.Print(c)
		} else {
			fmt.Printf("%.2f")
		}

	} else {
		fmt.Printf("%.2f")
	}
}

### Платный лого utf8.RuneCountInString if math.Modf
func main() {
	var m string
	fmt.Scan(&m)
	b := (utf8.RuneCountInString(m))
	d := float64(b) * 0.23
	d1, d2 := math.Modf(d)
	d3 := d2 * 100
	if d1 < 1 {
		fmt.Printf("%.f коп.", d3)
	} else {
		fmt.Printf("%.0f р. %.0f коп.", d1, d3)

	}
}

### M kvadrat if 
func main() {
	var (
		a float32
		b float32
		c float32
		d float32
	)
	fmt.Scan(&a, &b, &c)
	d = (b * b) - (4 * a * c)
	if a != 0 {
		if d > 0 {
			fmt.Println("два корня")
		} else if d == 0 {
			fmt.Println("один корень")
		} else {
			fmt.Println("корней нет")
		}
	} else if a == 0 {
		if b != 0 {
			fmt.Println("один корень")
		} else {
			fmt.Println("корней нет")
		}
	} else {
		fmt.Println("корней нет")
	}
}

###  Бесконечный цикл + other:
Polnuj nul
func main() {
other:
	for i := 1; i >= 1; {
		fmt.Scan(&i)
		if i == 0 {
			break other
		}
		fmt.Printf("%d\n", i)
	}
}

Vacuum
func main() {
other:
	for {
		var i string
		var j int
		fmt.Scanln(&i)
		j = len(i)
		if j == 0 {
			break other
		}
		fmt.Println(i)
	}
}

### Validacija utf8.RuneCountInString

func main() {
outher:
	for {
		var a string
		var b string
		fmt.Scan(&a, &b)
		if utf8.RuneCountInString(a) < 8 {
			fmt.Println("Слишком короткий пароль!")
		} else if strings.Contains(a, "123") || strings.Contains(a, "qwe") {
			fmt.Println("Слишком простой пароль!")
		} else if a != b {
			fmt.Println("Введенные пароли различаются!")
		} else {
			fmt.Println("Ну наконец-то!")
			break outher
		}
	}
}

### J vusokaja mat
func main() {
	var s int
	fmt.Scan(&s)
	for n := 0 - s; n <= s; n++ {
		m := n * n
		fmt.Printf("Квадрат числа %d равен %d \n", n, m)
	}
}
# Массивы. Основа
## 1. Определение массива.
	//Создадим массив под хранение 5-ти целочисленных элементов
	var arr [5]int // При инициализации массива важно передать информацию -сколько элементов в нем будет
	fmt.Println("This is my array:", arr)
## 2. Определение элементов массива (после предварительной инициализации)
	// Необходимо обратиться к элементу массива через синтаксис arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[3] = -500
	arr[4] = 720
	fmt.Println("After elemtns init:", arr)
## 3. Определние массива с указанием элементов на месте
	// Если при инициализации количество элементов меньше номинальной длины массива
	// то недостающие элементы инициализируются нулями
	newArr := [5]int{10, 20, 30}
	fmt.Println("Short declaration and init:", newArr)
## 4. Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))
	arrWithValues[0] = 10000
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))
## 5. Массив - это набор ЗНАЧЕНИЙ. То есть при всех манипуляциях - массив копируется (жестко, на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First arr:", first)
	fmt.Println("Second arr:", second)
## 6. Массив и его размер - это две составляющие одного типа (Размер массив - часть типа)
	// var aArr [5]int
	// var bArr [6]int
	// aArr[0] = 100
	// bArr = aArr
## 7. Итерирование по массиву
	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}
## 8. Итерирование по массиву через оператор range
	var sum float64
	for id, val := range floatArr {
		fmt.Printf("%d element of arr is %.2f\n", id, val)
		sum += val         //или sum = sum + val
	}
	fmt.Println("Total sum is:", sum)

      max := floatArr[0]
	for _, val := range floatArr {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max: ", max)

      min := floatArr[0]
	for _, val := range floatArr {
		if val < min {
			min = val
		}
	}
	fmt.Println("Min: ", min)



###  Игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}
## 9. Многомерные массивы
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}
	fmt.Println("Multidimensional array:", words)
### Итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}
## Slice append
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10   //изминение элемента 
	slice[1] = 200
	slice = append(slice, 200) // Добавление нового элемента
	for i, v := range slice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	emptySlice = append(emptySlice, 200)
## Slice Prepend // можно добавить только слайс, елемент в начало слайса хз как
// PREPEND
	bcd := []string{"b", "c", "d"}
	a := []string{"a"}

	bcd = append(a, bcd...)
	fmt.Println(bcd)
	// [a b c d]
## some
func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	s := len(a)              //количество элементов в массиве
	fmt.Println(a[0])        //вывод первого элемента в массиве
	fmt.Println(a[len(a)-1]) // выввод последнего элемента в массиве
	fmt.Println(a[0], a[1], a[2], s)
## range arr

	//вывод значений индексов и элементов
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//вывод только элементов
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q)
	var r [3]int = [3]int{1, 2} // по умолчанию r[2] = 0;( r[0, 1, 2])
	fmt.Println(r)
	fmt.Println(r[2])

	// Если троеточие в литерале масива то длинна массива определяется количеством инициализаторов
	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n %d\n", p, len(p))
##  Можно указать список пар ИНДЕКС - ЗНАЧЕНИЕ

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RUR
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "₤", RUR: "₽"}
	fmt.Println(RUR, symbol[RUR])

	// r := [...]int{99: -1} // массив 100 элементов один из которых -1
## Сравнение массивов ==, !=

	f := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(f == b, f == c, b == c) //true false false
	//d := [3]int{1, 2}
	//fmt.Println(f == d) // Ошибка компиляции разные типпы [2]int и [3]int
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
## Home work

### zad ARR G
func main() {
	var arr []float64
outher:
	for {
		val := 0.00
		fmt.Scan(&val)
		if val < 0 {
			max := arr[0]
			for _, val := range arr {
				if val > max {
					max = val
				}
			}
			min := arr[0]
			for _, val := range arr {
				if val < min {
					min = val
				}
			}
			var sum int
			for _, val := range arr {
				if val > 0 {
					sum++
				}
			}
			fmt.Printf("%d\n%.1f %.1f", sum, min, max)
			break outher
		}
		if val >= 100 && val <= 140 {
			arr = append(arr, val)
		}

	}
}

### zad ARR I matematik math.Modf
func main() {
	var c float64
	fmt.Scan(&c)
	s := []float64{}
	var i float64
	for i = c; i > 0; i-- {
		_, g := math.Modf(c / i)

		if g == 0 {
			s = append(s, i)
		}
	}
	//Перевернуть срез
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	if len(s) == 2 {
		for _, r := range s {
			fmt.Printf("%.0f ", r)
		}
		fmt.Println("\nACHTUNG")

	} else {
		for _, r := range s {
			fmt.Printf("%.0f ", r)
		}
	}
}






### L Sum Rjad Arr + for
func main() {
	var a int
	s := []int{}
	fmt.Scan(&a)
	for i := 0; i <= a-1; i++ {
		var c int
		fmt.Scan(&c)
		s = append(s, c)
	}
	//fmt.Println(s)
	var d int
	for i, val := range s {
		if i%2 == 0 {
			d = d + val
		} else {
			d = d - val
		}
	}
	fmt.Println(d)
}

### Mkvadrat koordinatu ARR
func main() {

	MKvadrat := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	for i, val := range MKvadrat {
		fmt.Scan(&val)
		MKvadrat[i] = val
	}
	x1 := MKvadrat[0]
	y1 := MKvadrat[1]

	x3 := MKvadrat[5]
	y3 := MKvadrat[6]

	//fmt.Println(MKvadrat)
	tochki := []int{}
	var a int
	fmt.Scan(&a)
	for i := 0; i <= (a*2)-1; i++ {
		var a int
		fmt.Scan(&a)
		tochki = append(tochki, a)
	}
	//fmt.Print(tochki)
	var x0 int
	var y0 int

	for i, val := range tochki {
		if i%2 == 0 {
			x0 = val
		} else if i%2 != 0 {
			y0 = val

			if x0 >= x1 && y0 >= y1 && x0 <= x3 && y0 <= y3 {
				fmt.Printf("Точка с координатами %d,%d принадлежит исследуемому квадрату\n", x0, y0)
			} else {
				fmt.Printf("Точка с координатами %d,%d не принадлежит исследуемому квадрату\n", x0, y0)
			}
		}
	}
}
# Slice

## 1. Слайсы (они же - срезы)
func main() {
  Обьявить слайс 
        var slice []int 
        slice := []int{}
  Добавить значение b
	var b int = 5
	arr = append(slice, b)
	*/
	// Слайс - это динамическая обвязка над массивом.
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] // Слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]:", startSlice)
	// Создали слайс, основываясь уже на существующем массиве

## 2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)

        //2.1 Общение со срезом == как массив только без указания начального размера
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10   //изминение элемента 
	slice[1] = 200
	slice = append(slice, 200) // Добавление нового элемента
	for i, v := range slice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	emptySlice = append(emptySlice, 200)

## 3. Измнение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4] // Это набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

## 4. Один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)

## 5. Срез как встроенный тип
	// type slice struct {
	// 	Length int
	// 	Capacity int
	// 	ZeroElement *byte
	// }

## 6. Длина и емкость слайса
	wordsSilce := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce))
	wordsSilce = append(wordsSilce, "four")
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce))
	// Capacity (cap) или ёмкость слайса - это значение, показывающее СКОЛЬКО ЭЛЕМЕНТОВ В ПРИНЦИПЕ
	// можно добавить в срез БЕЗ ВЫДЕЛЕНИЯ ДОПОЛНИТЕЛЬНОЙ ПАМЯТИ ПОД НИЖЕЛЕЖАЩИЙ МАССИВ.
	// Допустим у нас есть срез на 3 элемента (инициализировали без явного указания массива)
	// Компилятор при создании этого среза СНАЧАЛА создал массив ровно на 3 элемента
	// После этого компилятор вернул адрес, где этот масив живет
	// Срез запомнил этот адрес и теперь ссылается на него
	// Потом
	// Начинаем деформировать слайс (увеличим длину /увеличим количество элементов)
	// Проблема - в нижележащем массиве (на котором основан слайс) все 3 ячейки. Что делать?
	// Компилятор ищет в памяти место для массива размера 3*2 (в общем случае n*2, где n - первоначальный размер)
	// После того как место найдено (в нашем случае найдено место для 6 элементов), в это место копируются
	// старые 3 элемента на свои позиции. На 4-ую позицию мы добавляем новый элемент
	// После этого компилятор возвращает нашему слайсу новый адрес в памяти, где находится массив под 6 элементов.

	//Емкость всегда будет изменять как n*2.
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	//Важно: после выделения памяти под новый массив, ссылки со старым будут перенсены в новый
	// Пример
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

## 7. Как создавать слайсы наиболее эффективно.. make
	// make() - это функция, позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)
	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	//Сначала инициализируется arr = [15]int
	//Затем по нему делается срез arr[0:10]
	//После чего он возаращается
	fmt.Println(sl)

## 8. Добавление элементов в СРЕЗ
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherSlice := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlice...)
	myWords = append(myWords, "seven", "eight")
	fmt.Println("myWords:", myWords)

## 9. Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)
}

## 10. Mix random

import (
	"fmt"
	"math/rand"
	"slovarik/library"
	"time"
)

func main() {
	a := []string{"1", "2", "3"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a),
		func(i, j int) { a[i], a[j] = a[j], a[i] })

	fmt.Println(a)
}

func MixUp(l []library.Library) []library.Library {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(l),
		func(i, j int) { l[i], l[j] = l[j], l[i] })
	return l
}

## Useful func
### Del element Arr
func DelSliceElem(f int, c []int) ([]int, bool) {
	var tru bool
	if f > len(c) {
		tru = false
		c = c[:]
		return c, tru
	}

	for i, _ := range c {
		if i != len(c)-1 && i >= f-1 {
			c[i] = c[i+1]
		}
		if i == len(c)-1 {
			c = c[:i]
			tru = true
		}
	}
	return c, tru
}
### Delete Dublicat
func ReturnArr(arr []int) []int {
        sort.Ints(arr)
	Arc := []int{}
	for i, v := range arr {
		if i == len(arr)-1 { //Если конец то добавить последний элемент и фсё
			Arc = append(Arc, v)
		} else if v != arr[i+1] {
			Arc = append(Arc, v)
		}
	}
	return Arc
}
### Delete Dublicat 2
func DelDublikat(s []int) []int {
	c := []int{}
	count := 0
	for ii, v := range s {
		for i := ii; i <= len(s)-1; i++ {
			if v == s[i] {
				count++
			}
		}
		if count == 1 {
			c = append(c, v)
			count = 0
		} else {
			count = 0
		}
	}
	return c
}
### Del String DUblicat (Не упорядочено)
func DelDublikatString(s []string) []string {
	c := []string{}
	count := 0
	for ii, v := range s {
		for i := ii; i <= len(s)-1; i++ {
			if v == s[i] {
				count++
			}
		}
		if count == 1 {
			c = append(c, v)
			count = 0
		} else {
			count = 0
		}
	}
	return c
}
### Дубликаты удаление
func main() {
	s := []int{5, 1, 2, 3, 4, 5, 6, 6, 1}
	sort.Ints(s)
	d := DelDublikat(s)
	fmt.Println(d)

	dd := ReturnArr(s)
	fmt.Println(dd)
}

func ReturnArr(arr []int) []int {
	Arc := []int{}
	for i, v := range arr {
		if i == len(arr)-1 { //Последний элемент просто копируется
			Arc = append(Arc, v)
		} else if v != arr[i+1] {
			Arc = append(Arc, v)
		}
	}
	return Arc
}

func DelDublikat(s []int) []int {
	c := []int{}
	count := 0
	for ii, v := range s {
		for i := ii; i <= len(s)-1; i++ {
			if v == s[i] {
				count++
			}
		}
		if count == 1 {
			c = append(c, v)
			count = 0
		} else {
			count = 0
		}
	}
	return c
}
### Удаление дубликатов в срезе структур с изминением порядка элементов 
// Функция которая пинимает два среза
//   - удаляет дубликаты
//   - если есть дубликаты то перемещает их в начало среза но при этом оставляя их неизменными
import (
	"fmt"
	"strings"
)

type model struct {
	name string
	num  int
}

func main() {
	var originArr = []*model{

		{
			name: "five",
			num:  1,
		},
		{
			name: "six",
			num:  1,
		},
		{
			name: "three",
			num:  1,
		},
		{
			name: "four",
			num:  1,
		},
	}

	var needAdd = []*model{
		{
			name: "one",
			num:  1,
		},
		{
			name: "two",
			num:  1,
		},
		{
			name: "three",
			num:  2,
		},
		{
			name: "four",
			num:  2,
		},
	}
	needAdd = append(needAdd, originArr...)

	withoutDublAndChange := DelDublikat(needAdd)
	for _, v := range withoutDublAndChange {
		fmt.Printf("name: %s, num: %v ", v.name, v.num)
	}

}

// на выходе должен получится такой массив
// [{one 1} {two 1} {three 1} {four 1} {five 1} {six 1} ]
// изменять num в originArr нельзя

func ReverseSlice(s []*model) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func DelDublikat(s []*model) []*model {

	var count = 0
	for i := 0; i <= len(s)-1; i++ {
		for _, v := range s {
			if strings.EqualFold(v.name, s[i].name) {
				count++
			}
			if count == 2 {
				s[i].num = v.num
				count = 0
			}
		}
	}
	fmt.Println("stage1", s)
	ReverseSlice(s)
	withoutDublicat := []*model{}
	for ii, v := range s {
		var count1 int
		for i := ii; i <= len(s)-1; i++ {
			if strings.EqualFold(v.name, s[i].name) {
				count1++
			}
		}
		if count1 == 1 {
			withoutDublicat = append(withoutDublicat, v)
			count1 = 0

		} else {

			count1 = 0
		}
	}
	ReverseSlice(withoutDublicat)
	return withoutDublicat
}
# String как Slice byte, rune
## 1. Строка - это байтовый слайс со своими особенносятми при обращении к нижелажащему массиву
	word := "Тестовая строка"
	fmt.Printf("String %s\n", word)
	// Какие значения байт сейчас находятся в слайсе word?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления 16-ти ричного байта
	}
	fmt.Println()
	// Каким образом получать доступ к отдельно стоящим символам?
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) //%c - формат представления символа
	}
	fmt.Println()
## 2. Строки в Go - хранятся как наборы UTF-8символов. Каждый символ, вообще говоря, может занимать больше чем 1 байт
## 3. Руна (Rune) - это стандартный встроенный тип в Go (alias над int32), позволяющий хранить
	//единый неделимый UTF символ ВНЕ ЗАВИСИМОСТИ ОТ ТОГО сколько байт он занимает
	fmt.Printf("Runes: ")
	runeSlice := []rune(word) // Преобразование слайса байт к слайсу рун []byte(sliceRune)
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()
## 4. Итерирование по строке с использованием рун
	for id, runeVal := range word { // id - это индекс байта, с КОТОРОГО НАЧИНАЕТСЯ РУНА runeVal
		fmt.Printf("%c starts at postion %d\n", runeVal, id)
	}
## 5. Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43} // Исходное представление байтов
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 101, 102, 103} // Синтаксический сахар - можно использовать десятичное представление байтов
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)
## 6. Создание строки из слайса рун
	// Руны как hex
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRune := string(runeHexSlice)
	fmt.Println("From Runes(hex):", myStrFromRune)
	// Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 'y', 'a'} // '' - таким образом обозначается руна
	myStrFromRuneLiterals := string(runeLiteralSlice)
	fmt.Println("From Runes(literals):", myStrFromRuneLiterals)

	fmt.Printf("%s and %T\n", myStrFromRuneLiterals, myStrFromRuneLiterals)
## 7. Длина и емкость строки
	// Длина len() - количество байт в слайсе
	fmt.Println("Length of Вася:", len("Вася"), "bytes")
	// Длина RuneCounter - количество !рун!
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes")
	// Вычисление емкости строки - бессмысленно, т.к. строка базовый тип
	fmt.Println(cap([]rune("Вася")))
## 8. Сравнение строк == и !=. Начиная с go 1.6
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
## 9. Конкатенация
	word3 := word1 + word2
	fmt.Println(word3)

	//10. Строитель строк (java -> StringBuiler)
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s #### %s", firstName, secondName)
	fmt.Println("FullName:", fullName)
## 11. Строки не изменяемые А слайсы изменяемые
    Строки не изменяемые
	// fullName[0] = "Q"

	 А слайсы изменяемые :)
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'
	fullName = string(fullNameSlice)
	fmt.Println("String mutation:", fullName)
## 12. Сравнение рун
	if 'Q' == 'Q' {
		fmt.Println("Runes equal")
	} else {
		fmt.Println("Runes not equal")
	}
## 14. Где живут полезные методы работы со строками? import "strings"
	// import "strings"
## 15. И помни Arr str to String
    Range делит слайс на руны !!!
func main() {
	i := "I love you Baby"
	SliceI := []string{}

	for _, v := range i {
		d := string(v)
		SliceI = append(SliceI, d)
	}

	fmt.Println(SliceI)
	var StringI string
	for _, v := range SliceI {
		StringI = StringI + v
	}
	fmt.Println(StringI)
}
## Home Work
### //https://contest.yandex.ru/contest/25667/problems/C/

func main() {
	//Первое слово
	var str string
	fmt.Scan(&str)
	STR1 := []rune{}
	var f1 rune
	for _, v := range str {
		STR1 = append(STR1, v)
		f1 = STR1[len(STR1)-1] //последняя буква первого слова
	}
	//fmt.Printf("f1 %c\n", f1)
	var str1 string
	for {
		fmt.Scan(&str1)
		STR := []rune{}
		for _, v := range str1 {
			STR = append(STR, v)
		}
		f2 := STR[0] // Первая буква второго слова
		//fmt.Printf("f2 %c\n", f2)
		if f1 == f2 {
			f1 = STR[len(STR)-1]
			continue
		} else {
			fmt.Print(str1)
			break
		}
	}

}
### //https://contest.yandex.ru/contest/25667/problems/D/
func main() {
	//Первое слово
	var str string
	fmt.Scan(&str)
	STR1 := []rune{}
	var f1 rune
	for _, v := range str {
		STR1 = append(STR1, v)
		if STR1[len(STR1)-1] != 'ь' {
			f1 = STR1[len(STR1)-1] //последняя буква первого слова
		} else {
			f1 = STR1[len(STR1)-2]
		}
	}
	//fmt.Printf("f1 %c\n", f1)
	var str1 string
	for {
		fmt.Scan(&str1)
		STR := []rune{}
		for _, v := range str1 {
			STR = append(STR, v)
		}
		f2 := STR[0] // Первая буква второго слова
		//fmt.Printf("f2 %c\n", f2)
		if f1 == f2 { // ? мягкий знак во втором слове ?
			if STR[len(STR)-1] != 'ь' {
				f1 = STR[len(STR)-1]
			} else {
				f1 = STR[len(STR)-2]
			}
			continue
		} else {
			fmt.Print(str1)
			break
		}
	}
}
### //https://contest.yandex.ru/contest/25667/problems/E/
func main() {
	var str string
	STR := []rune{}
	fmt.Scan(&str)
	for _, val := range str {
		STR = append(STR, val)
	}
	for i, val := range STR {
		if i%2 == 0 {
			for i = 0; i <= 2; i++ {
				fmt.Printf("%c", val)
				//STR = append(STR, val)
			}
		}
	}
}
### //https://contest.yandex.ru/contest/25667/problems/F/
func main() {
	var str string
	STR1 := []rune{}
	STR := []rune{}

	fmt.Scan(&str)

	for _, v := range str {
		fmt.Printf("%c", v)
		STR1 = append(STR1, v)
	}
	fmt.Println()

	for {

		for i, val := range STR1 {
			if i >= 2 {
				fmt.Printf("%c", val)
				STR = append(STR, val)
			}
		}

		if len(STR) <= 1 {
			break
		}
		fmt.Println()

		STR1 = STR[:len(STR)-2]
		STR = STR[:0]
		for _, val := range STR1 {
			fmt.Printf("%c", val)
		}
		if len(STR1) <= 1 {
			break
		}
		fmt.Println()
	}
}
### //https://contest.yandex.ru/contest/25667/problems/G/
func main() {
	//https://contest.yandex.ru/contest/25667/problems/G/?success=70411031#2706657/2020_04_14/lSuTYKnvCT
	var str string
	STR := []rune{}
	fmt.Scanln(&str)
	fmt.Print(str) //
	fmt.Println()  //
outher:
	for _, v := range str {
		if v != 'о' && v != 'р' {
			str = ""
			break outher
		}
	}
	for _, v := range str {
		STR = append(STR, v)
	}
	I := []int{}
	var i int
	for _, v := range STR {
		if v == 'о' {
			i++
		} else if v == 'р' {
			I = append(I, i)
			i = 0
		}
		I = append(I, i)
	}
	var max int
	for _, val := range I {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
### //https://contest.yandex.ru/contest/25667/problems/H/
func main() {
	// Скан входящих
	var i int
	gslice := []int{}
	var c int
	var d int
	fmt.Scan(&i)
	for a := 1; a <= i; a++ {
		var b int
		fmt.Scan(&b)
		gslice = append(gslice, b)
	}
	fmt.Scan(&c)
	fmt.Scan(&d)
	var sum int
	for i, val := range gslice {
		if i >= c-1 && i <= d-1 {
			sum = val + sum
		}
	}
	fmt.Print(sum)
}
### //https://contest.yandex.ru/contest/25667/problems/B/
func main() {
	var str string
	fmt.Scan(&str)
	STR := []rune{}
	for _, v := range str {
		STR = append(STR, v)
	}

	f1 := STR[0]
	f2 := STR[len(STR)-1]
	arr := [2]rune{f1, f2}
	arr2 := [2]rune{f2, f1}

	str1 := [2]rune{'Д', 'а'}
	str2 := [2]rune{'д', 'А'}
	str3 := [2]rune{'Д', 'А'}
	str4 := [2]rune{'д', 'а'}

	if arr == str1 || arr == str2 || arr == str3 || arr == str4 || arr2 == str1 || arr2 == str2 || arr2 == str3 || arr2 == str4 {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}

}
# bufio, strconv.Atoi, strconv.ParseInt
## bufio
func main() {
	var name string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() { // Команда захвата потока ввода и сохранения в буфер (захват идет до символа окончания строки)
		name = input.Text() // Команда возвращения элементов, помещенных в буфер (отдаст string)
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}
## strconv.Atoi Преоброазование строкового литерала к чему-нибудь числовому 
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr) // Atoi - Anything to Int (именно int)
	fmt.Printf("%d is %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	numFloat32, _ := strconv.ParseFloat(numStr, 32) // Но это 64-разрядное число будет без проблем ГАРАНТИРОВАНО ПРЕВОДИТЬСЯ К 32
	numInt2:= strconv.Itoa(numInt)
	fmt.Println(numInt64, numFloat32, numInt2)
	fmt.Printf("%.3f and %T\n", numFloat32, float32(numFloat32))
}
## Home Work
### //https://contest.yandex.ru/contest/25667/problems/A/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var m int // Нужно мне
	var mm string
	mArr := []string{}
	var n int // На сервисе
	var nn string
	nArr := []string{}

	in := bufio.NewScanner(os.Stdin)

	// На сервисе
	if in.Scan() {
		nn = in.Text()
	}
	n, _ = strconv.Atoi(nn)

	// Нужно мне
	if in.Scan() {
		mm = in.Text()
	}
	m, _ = strconv.Atoi(mm)

	//Arr На сервисе
	for a := 1; a <= n; a++ {
		var b string
		if in.Scan() {
			b = in.Text()
		}
		nArr = append(nArr, b)
	}

	//Arr Нужно мне
	for a := 1; a <= m; a++ {
		var b string
		if in.Scan() {
			b = in.Text()
		}
		mArr = append(mArr, b)
	}

	for _, val := range mArr {
		var c int
		for _, v := range nArr {
			if val == v {
				fmt.Println("ЕСТЬ")
				c = 1
			}
		}

		if c != 1 {
			fmt.Println("НЕТ В НАЛИЧИИ")
			c = 0
		}
	}
}
### //https://contest.yandex.ru/contest/25667/problems/I/
func main() {

	STR := []string{}
	var i int
	var ii string

	input := bufio.NewScanner(os.Stdin)
	//4
	if input.Scan() {
		ii = input.Text()
	}

	i, _ = strconv.Atoi(ii)
	// arr4
	for a := 1; a <= i; a++ {
		var b string
		if input.Scan() {
			b = input.Text()
		}
		STR = append(STR, b)
	}

	numStr := []int{}
	var cq int
	var c string
	// 2
	if input.Scan() {
		c = input.Text()
	}
	cq, _ = strconv.Atoi(c)
	//arr 2
	for a := 1; a <= cq; a++ {
		var d int
		var dd string
		if input.Scan() {
			dd = input.Text()
			d, _ = strconv.Atoi(dd)
		}
		numStr = append(numStr, d)
	}

	for _, val := range numStr {
		for c, v := range STR {
			if val == c+1 {
				fmt.Println(v)
			}
		}
	}
}
### //https://contest.yandex.ru/contest/25667/problems/J/
func main() {
	//https://contest.yandex.ru/contest/25667/problems/J/?success=70629939#2706657/2020_03_25/7a3dHNDkst
	in := bufio.NewScanner(os.Stdin)
	// Число блюд которое может приготовить столовая
	var m int
	M := []string{}
	var mm string
	if in.Scan() {
		mm = in.Text()
	}
	m, _ = strconv.Atoi(mm)
	for i := 1; i <= m; i++ {
		var mt string
		if in.Scan() {
			mt = in.Text()
		}
		M = append(M, mt)
	}
	//fmt.Println(M)
	//Число дней для которых есть списки блюд
	var n int
	var nn string
	if in.Scan() {
		nn = in.Text()
	}
	n, _ = strconv.Atoi(nn)
	//fmt.Println(n)
	//Многомерный массив
	BB := [][]string{}
	for i := 1; i <= n; i++ {
		B := []string{}
		var b int
		var bb string
		if in.Scan() {
			bb = in.Text()
		}
		b, _ = strconv.Atoi(bb)
		var cc string
		for ii := 1; ii <= b; ii++ {
			if in.Scan() {
				cc = in.Text()
			}
			B = append(B, cc)
		}
		BB = append(BB, B)

	}
	BBB := []string{}
	for _, v := range BB {
		for _, val := range v {
			BBB = append(BBB, val)
		}
	}
	//fmt.Println(BBB)
	B1 := []string{}
	for _, v := range M {
		var c int
		for _, val := range BBB {
			if val == v {
				c = 1
			}
		}
		if c == 1 {
			continue
		}
		B1 = append(B1, v)
	}
	sort.Strings(B1)
	for _, v := range B1 {
		fmt.Println(v)
	}
}
# Map

## Map - это набор пар ключ:значение. Инициализация пустой мапы

## 1. Инициализация пустой мапы
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

## 2. Добавление пар в существующую мапу
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mapper after adding pairs:", mapper)

## 3. Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}
	newMapper["Jo"] = 3000
	fmt.Println("New Mapper:", newMapper)

## 4. Что может быть ключом в мапе?
	//4.1 ВАЖНО: Мапа НЕ УПОРЯДОЧЕНА В GO
	//4.2 КЛЮЧОМ В МАПЕ МОЖЕТ БЫТЬ ТОЛЬКО СРАВНИМЫЙ ТИП (==, !=)

## 5. Нулевое значение для map
	// var mapZeroValue map[string]int // mapZeroValue == nil
	// mapZeroValue["Alice"] = 12

## 6. Получение элементов из map
	//6.1 Получение элемента, который представлен в мапе
	testPerson := "Alice"
	fmt.Println("Salary of testPerson:", newMapper[testPerson])
	//6.2 Получение элемента, который НЕ представлен в мапе
	testPerson = "Derek"
	fmt.Println("Salary of new testPerson:", newMapper[testPerson]) // При обращении по несуществующему ключу - новая пара не добавляется
	fmt.Println(newMapper)

## 7. Проверка вхождения ключа
	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}

	//7.1 При обращении по ключу - возвращается 2 значения
	if value, ok := employee["Den"]; ok {
		fmt.Println("Den and value:", value)
	} else {
		fmt.Println("Den does not exists in map")
	}

	if value, ok := employee["Jo"]; ok {
		fmt.Println("Jo and value:", value)
	} else {
		fmt.Println("Jo does not exists in map")
	}

## 8. Перебор элементов мапы
	fmt.Println("==============================")
	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}

## 9. Как удалять пары
	//9.1 Удаление существующей пары
	fmt.Println("Before deleting:", employee)
	delete(employee, "Den")
	fmt.Println("After first deleting:", employee)

	//9.2 Удаление не существующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna") // ОЧЕНЬ ДОРОГАЯ ОПЕРАЦИЯ
	}

	fmt.Println("After second deleting:", employee)

## 10. Количество пар == длина мапы
	fmt.Println("Pair amount in map:", len(employee))

	//11. Мапа (как и слайс) ссылочный тип
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}

	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("words map:", words)
	fmt.Println("newWords map:", newWords)

## 12 . Сравнение мап
	//12.1 Сравнение массивов (массив можно использовать как ключ в мапе)
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	//12.2 Сравнение слайсов. Слайсы (из-за того что тип ссылочный - можно сравнить на равенство только с nil)
	// if []int{1, 2, 3} == []int{1, 2, 3} {
	// 	fmt.Println()
	// }

	//12.3 Сравнение мап.  Мапа (из-за того, что тип ссылочный - можно сравнивать только с nil)
	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("Zero value map")
	}

	if bMap == nil {
		fmt.Println("Zero value of map bMap")
	}

## 14. Упорядочить Map по значению
	// Если мапа/слайс являются составляющими какой-либо структуры - структура автоматически не сравнима

        //14. Упорядочить Map по значению
        ages := map[string]int{
		"Андрей":  30,
		"Nastyxa": 25,
	}
	ages["Наталья"] = 31
	ages["Mixail"] = ages["Наталья"] + ages["Андрей"]
	//Как упорядочить мап ?
	fmt.Println(ages)
	// С помощью структур
	type kv struct {
		k string
		v int
	}
	kvs := make([]kv, 0, len(ages))
	for k, v := range ages {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v < kvs[j].v
	})
	fmt.Println(kvs)
	// Преобразовать структуру в мап
	mainMap := make(map[string]int)
	for _, v := range kvs {
		mainMap[v.k] = v.v
	}
	fmt.Println(mainMap)
}
# Func

## 1. Явная функция - локально-определенный блок кода , имеющий имя (ЯВНОЕ ОПРЕДЕЛЕНИЕ)
// Функцию необходимо определить + Функциию необходимо вызвать
/*
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mult(a, b int) int {
	return a * b
}
func calcAndReturnValidFunc(command string, a int, b int) int {
	if command == "additions" {
		return add(a, b)
	} else if command == "substraction" {
		return sub(a, b)
	} else {
		return mult(a, b)
	}
}
*/
## 2. Сигнатура функций и их определение
// func functionName(params type) typeReturnValue {
// 	//body
// }

func main() {
	fmt.Println("Hello world")
	//3. Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)
	fmt.Println("Result of mult(1, 2, 3, 4):", mult(1, 2, 3, 4))
	per, area := rectangleParameters(10.5, 10.5)
	newPer, _ := rectangleParameters(10, 10)
	fmt.Println("Area of rect(10.5, 10.5):", area)
	fmt.Println("Perimeter of rect(10.5, 10.5):", per)
	fmt.Println("NewPer:", newPer)
	secondPer, secondArea := namedReturn(10, 20)
	fmt.Println(secondArea, secondPer)
	emptyReutrn(10)
	helloVariadic(10, 20, 30, 40, 50, 60, 60)
	helloVariadic()
	someStrings(2, 3)
	sum1 := sumVariadic(10, 30, 40, 50, 60)
	sliceNumber := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)

	fmt.Println(sumSlice([]int{30, 40, 50, 60, 80, 90, 100}))
	fmt.Println(sumSlice(sliceNumber))

	//12. Анонимная функция (синтаксис)
	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("Anon:", anon(20, 30))
	fmt.Println("BigFunction(10, 20):", bigFunction(10, 20))

}

//13. Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b + 1 }(aArg, bArg)
}

## 3. Простейшая функция (определить функцию можно как до момента ее вызова в функции main,
// так и в любом месте пакета, главное чтобы она была определена в принципе где-нибудь)
func add(a int, b int) int {
	result := a + b
	return result
}

## 4. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

## 5. Возврат больше чем одного значения (returnType1, returnType2.......)
func rectangleParameters(length, width float64) (float64, float64) {
	var perimeter = 2 * (length + width)
	var area = length * width

	return perimeter, area
}

## 6. Именованный возврат значений
func namedReturn(a, b int) (perimeter int32, area int) {
	perimeter = int32(2 * (a + b))
	area = a * b
	return // Не нужно указывать возвращаемые переменные
}

## 7. При вызове оператора return функцию прекращает свое выполнение и возвращает что-то
func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

## 8. Что вернется в случае, если return в принципе не указан (или он пустой)
func emptyReutrn(a int) {
	fmt.Println("I'M emptyReturn with parameter:", a)
}

## 9. Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("value %v and type %T\n", a, a)
}

## 10. Смешение параметров с variadic (
// 	1. Континуальынй параметр всегда самый последний
//  2. Variadic параметр - на всю функцию один (для удобочитаемости)
// )
func someStrings(a, b int, words ...string) {
	fmt.Println("Parameter:", a)
	fmt.Println("Parameter:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Result concat:", result)
}

## 11. Передача слайса или использование variadic parameters?
func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func sumSlice(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

## 12. Анонимная функция (синтаксис)
	anon := func(a, b int) int {
		return a + b
	}

## 14. Возврат функции в качестве значения
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

func main() {
	var command string
	command = "addition"
	res := calcAndReturnValidFunc(command)
	res1 := calcAndReturnValidFunc("substraction")(110, 20)
	command1 := "faction"
	a := 3
	b := 5
	resab := calcAndReturnValidFunc(command1)(a, b)
	fmt.Println(command, res(10, 20), "\nsubstraction", res1, "\n", command1, resab)
}

## 15.Функция как параметр в другой функции

func recieveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 200
	intVarB = 100

	return f(intVarA, intVarB)
}
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("rec(func(a, b int) int { return a - b }", recieveFuncAndReturnValue(func(a, b int) int {
		return a - b
	}))
	fmt.Println("recieveFuncAndReturnValue(add):", recieveFuncAndReturnValue(add)) // !add - isn`t add()
	fmt.Println("rec(func(a, b int) int { return a*a + b*b }", recieveFuncAndReturnValue(func(a, b int) int {
		return a*a + b*b + 2*a*b
	}))
}

## Тип функции 

### 1. Явные функции (в принципе любая функция в Go) - является
//экземлпяром 1-го уровня (функцию можно присваивать в перменную, ее можно передавать в
//качестве параметра и возвращать из других функций)

### 2. Возврат функции в качестве значения
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

### 3. Функция как парметр в другой функции
func recieveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200

	return f(intVarA, intVarB)
}

func add(a, b int) int {
	return a + b
}
func main() {

	var command string
	command = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command :", command, "value:", res(10, 20))
	fmt.Println(res(30, 40))

### 4. Тип функции
	fmt.Printf("Type of func is %T\n", res)
	fmt.Printf("Type of calcAndReturnValidFunc is %T\n", calcAndReturnValidFunc)

### 5. Тип функции в Go определяется как входными парамтерами, так и выходными

	fmt.Println("recieveFuncAndReturnValue(add):", recieveFuncAndReturnValue(add))
	fmt.Println(recieveFuncAndReturnValue(func(a, b int) int {
		return a*a + b*b + 2*a*b
	}))
}
# Constanta 

## 1. Константы - это неизменяемые переменные, которые служат для:
//	1) Более строгого понимания кода
//	2) Для того, чтобы случайно не поменять значение (предполагается что значение константы не изменно)
//	3) Для удобных преобразований

const (
	MAIN_PORT = "8001"
)

## 2. Объвление одной константы
func main() {

	const a = 10
	fmt.Println(a)
## 3. Объявление блока констант с областью видимости внутри функции main
	const (
		ipAddress string = "127.127.00.03"
		port             = "8000"
		dbName           = "postgres"
	)
	fmt.Println("ipAddress value:", ipAddress)
	fmt.Println(checkPortIsValid())

## 4. Константу никак нельзя поменять в ходе работы программы
	// const b = 200
	// b = 30

## 5. Значения констант ДОЛЖНЫ БЫТЬ ИЗВЕСТНЫ на момент компиляции
	var sqrt = math.Sqrt(25)
	//const sqrt = math.Sqrt(25) //Нельзя присвоить в константу что-либо, что является результатом вызова функции, метода
	fmt.Println("Var sqrt:", sqrt)

## 6. Типизированные и нетипизированные константы
	const ADMIN_EMAIL string = "admin@admin.com" // Указание типа - повышение читабельности кода

## 7. Нетипизирвоанные константы и их профит
	//При использовании нетипизированных констант мы разрешаем компилятору
	//исопльзовать неявное приведение типов в момент присваивания значеий констант в перменные
	const NUMERIC = 10
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex complex64 = NUMERIC

	fmt.Printf("numInt8 value %v type %T\n", numInt8, numInt8)
	fmt.Printf("%v + %v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC)
	fmt.Printf("numInt32 value %v type %T\n", numInt32, numInt32)
	fmt.Printf("numInt64 value %v type %T\n", numInt64, numInt64)
	fmt.Printf("numFloat32 value %v type %T\n", numFloat32, numFloat32)
	fmt.Printf("numComplex value %v type %T\n", numComplex, numComplex)
}

## 8. Константы в Go зашиваются в момент компиляции в RUNTIME программы и не выбрасываются до ее окончания

func checkPortIsValid() bool {
	if MAIN_PORT == "8001" {
		return true
	}
	return false
}
# Pointer 

## 1. Указатели - переменная, хранящая в качестве значения - адрес в памяти другой переменной

func main() {
## 2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable //&.... - операция взятия адреса в памяти
	//Выше у нас создан указатель на переменную variable
	//В pointer лежит 18293xcd000132 - это место в памяти, где хранится int значение 30
	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)
## 3. А какое нулевое значение для указатели?
	var zeroPointer *int //zeroValue имеет значение nil (это указатель в никуда)
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initializatoin type %T and value %v\n", zeroPointer, zeroPointer)
	}
## 4. Разыменование указателя (получение значения): *pointer - возвращает значение, хранимое по адресу
	var numericValue int = 32
	pointerToNumeric := &numericValue

	fmt.Printf("Value in numericValue is %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)
## 5. Создание указателей на нулевые занчения типов
	// var zeroVar int
	// var zeroPoint *int = &zeroVar
	zeroPoint := new(int) // Создает под капотом zeroValue для int, и возвращает адрес, где этот 0 хранится
	fmt.Printf("Value in *zeroPointer %v\nAddress is %v\n", *zeroPoint, zeroPoint)
## 6. Изменение значения хранимого по адресу через указатель
	zeroPointerToInt := new(int)
	fmt.Printf("Addres is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Addres is %v and New Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)

	b := 345
	a := &b
	c := &b
	*a++
	*c += 100
	fmt.Println(b)
## 7. Указательная арфиметика ОТСУТСТВУЕТ ПОЛНОСТЬЮ
	// У вас на руках адрес одной ячейки - вы можете через этот адрес продвинуться в другие ячейки
	//
## 8. Передача указателей в функции
    // Значение нужно копировать и в итоге всеравно работать с адрессом 
	// Колоссальный прирост производительности засчет того, что передается не значение (которые должно копироваться)
	// а передается лишь адрес в памяти, за которым уже хранится какое-то значение
	sample := 1
	//samplePointer := &sample

	fmt.Println("Origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("After changing sample is:", sample)
## 9. Возврат поинтера из функции (В С++ результат работы такого механизма - неопределен)
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr2: %T and address %v and value %v\n", ptr2, ptr2, *ptr2)

}

//9.1  Инициализация функции, возвращающей указатель
func returnPointer() *int {
	var numeric int = 321
	return &numeric //В момент возврата Go перемещает данную переменную в кучу
}

//8.1 Определдение фукнции, принимающей параметр как указатель
func changeParam(val *int) {
	*val += 100
}
## 10. Указатели на массивы. Почему так делать не надо
func mutation(arr *[3]int) {
	// (*arr)[1] = 909
	// (*arr)[2] = 100000
	//Можно написать и так, т.к. Go сам разыменует указатель на массив (из-за того,что функци принимает *arr)
	arr[1] = 909
	arr[2] = 10000
}

 Используйте лучше слайсы (это идеоматично с точки зрения Go)
func mutationSlice(sls []int) {
	sls[1] = 909
	sls[2] = 10000
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation:", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation:", arr)

	newArr := [3]int{1, 2, 4}
	fmt.Println("newArr before mutationSlice:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutationSlcie:", newArr)
}
# Struct 

## 1. Структура - заименованный набор полей (состояний), определяющий новый тип данных.
## 2. Определение структуры (явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}
## 3. Если имеется ряд состояний одного типа, можно сделать так
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}
## 11. Структура с анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func PrintStudent(std Student) {
	fmt.Println("==================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName:", std.lastName)
	fmt.Println("Age:", std.age)
}

func main() {
## 4. Создание представителей структуры
	stud1 := Student{
		firstName: "Fedya",
		age:       21,
		lastName:  "Petrov",
	}
	PrintStudent(stud1)
	stud2 := Student{"Petya", "Ivanov", 19} // Порядок указания свойств - такой же как в структуре
	PrintStudent(stud2)
## 5. Что если не все поля структуры определить?
	stud3 := Student{
		firstName: "Vasya",
	}
	PrintStudent(stud3)
## 6. Анонимные структуры (структура без имени)
	anonStudent := struct {
		age           int
		groupID       int
		proffesorName string
	}{
		age:           23,
		groupID:       2,
		proffesorName: "Alexeev",
	}
	fmt.Println("AnonStudent:", anonStudent)
## 7. Доступ к состояниям и их модфикация
	studVova := Student{"Vova", "Ivanov", 19}
	fmt.Println("firstName:", studVova.firstName)
	fmt.Println("lastName:", studVova.lastName)
	fmt.Println("age:", studVova.age)
	studVova.age += 2
	fmt.Println("new age:", studVova.age)
## 8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)
## 9. Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}
	fmt.Println("Value studPointer:", studPointer)
	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value afterPointerModify:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)
## 10. Работа с доступ к полям структур через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	fmt.Println("Age via .age:", studPointer.age) //Неявно происходит разыменование указателя studpointer и запрос соотв поля
## 12. Создание экземпляра с анонимными полями структуры
	human := &Human{
		firstName: "Bob",
		lastName:  "Johnson",
		string:    "Additional Info",
		int:       -1,
		bool:      true,
	}

	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
}
## 14. Массив указателей (отсебятинка)
type Employee struct {
	name   string
	salary int
}

func NewEmployee(newName string, newSalary int) *Employee {
	return &Employee{newName, newSalary}
}

func main() {
	db := []*Employee{}
	for i := 1; i <= 50; i++ {
		d := strconv.Itoa(i)
		a := "Employee " + d
		NewEmployeer := NewEmployee(a, i*2+2)
		db = append(db, NewEmployeer)
	}
	fmt.Print("Name ", db[0].name, " ||Salary ", db[0].salary)
}
# Вложенные структуры

## 1. Вложенные структуры (вложение структур). Это использование одной структуры, как тип поля
//в другой структуре
type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}
## 4. Встроенные структуры (когда мы добавляем поля одной структуры к другой)
type Professor struct {
	firstName string
	lastName  string
	age       int
	greatWork string
	//papers     map[string]string - добавление этого поля делает структуру несравнимой
	University // В этом месте происходит добавление всех полей структуры Uni в Professor
}

func main() {
## 2. Создание экземпляров структур с вложением
	stud := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "cool University",
			infoLong:  "very cool University",
		},
	}
## 3. Получение доступа к вложенным полям структур
	fmt.Println("FirstName:", stud.firstName)
	fmt.Println("LastName:", stud.lastName)
	fmt.Println("Year based Uni:", stud.university.yearBased)
	fmt.Println("Long info:", stud.university.infoLong)
## 5. Создание экземпляра с встраиванием структур
	prof := Professor{
		firstName: "Anatoly",
		lastName:  "Smirnov",
		age:       125,
		greatWork: "Ultimate C programming",
		University: University{
			yearBased: 1734,
			infoShort: "short Info",
			age:       2021 - 1734,
		},
	}
## 6. Обращение к состояниям с встроенной структурой
	fmt.Println("FirstName:", prof.firstName)
	fmt.Println("Year based:", prof.yearBased)
	fmt.Println("Info Short:", prof.infoShort)
	fmt.Println("Age:", prof.University.age) //prof.age - получим доступ к полю ВЫШЕЛЕЖАЩЕЙ СТРУКТУРЫ
## 7. Сравнение экземпляров ==
	//При сравнении экзмеляров происходит сравнение всех их полей друг с другом
	profLeft := Professor{}
	profRight := Professor{}

	fmt.Println(profLeft == profRight)
## 8. Если ХОТЯ БЫ ОДНО ИЗ ПОЛЕЙ СТРУКТУР - НЕ СРАВНИМО - то и вся структура несравнима
}
## Sort Struct 

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{"Sally", 20}, {"David", 40}, {"Jon", 30}, {"Larry", 25}}
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	fmt.Println("Sorted by age:", people)

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})

	fmt.Println("Sorted by name:", people)
}
# Metod
## 1. Методы - функции, привязанные к определенным структурам

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

//1. Методы - функции, привязанные к определенным структурам
//func (s Struct) MethodName(parameters type) type {}
//      Reciever - получатель метода
func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Postion:", e.position)
	fmt.Printf("Salary : %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name:     "Bob",
		position: "Senior Golang developer",
		salary:   3000,
		currency: "USD",
	}

## //2. Вызов метода
	emp.DisplayInfo()
}

## //3. В чем преимущество методов над функциями?
// Во-первых: наличие методов улучшает "консистентность" кода, т.к. напрямую влияет на его понимание.
// Во-вторых: в Go запрещается создавать функции с одинаковыми названиями, в то время как методы для разных структур,
// с одинаковыми названиями - разрешены

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

## //4. Создадим функцию и метод Perimeter  для структуры Circle
func PerimeterCircle(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

## //5. Создадим функцию и метод Perimeter для структуры Rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

## //6. В Go разрешено создавать методы с одинаковыми именами в пределах одной структуры. Главное, чтобы
//получать метода в разных структурах (где реализован мтеод со сходим именем) отличался.
func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function:", PerimeterCircle(c))
	fmt.Println("Call method:", c.Perimeter())

	r := Rectangle{10, 20}
	fmt.Println("Call function for rectangle:", PerimeterRectangle(r))
	fmt.Println("Call method for rectangle:", r.Perimeter())
}

## Metod with POINTER

type Employee struct {
	name   string
	salary int
}

### //1. Метод, в котором получатель копируется и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

### //2. Метод, в котором получатель передается по ссылке (в теле метода работаем с ссылкой на экземпляр)
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

### //4. Используйте методы с PointerReciever'ом в ситуациях когда:
// 1) Изменения в работе метода над экземпляром должны быть видны на вызывающей стороне
// 2) Когда экземпляр достаточно "увесистый", то есть копировать его дорого с точки зрения расходов ресурсов
// 3) С ними может работать любой вид экземпляра

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before setting parameters:", e)
	e.SetName("Bob")
	fmt.Println("After SetName call:", e)
	e.SetSalary(4500) //3. Вызов метода у ссылки на сотрудника
### // 5. Аналогично явному вызову (&e).SetSalary(9999)
	fmt.Println("After SetSalary call:", e)
}

## Metod anonymus

type University struct {
	city string
	name string
}

### //1. Данный метод явно связан только с структурой University
func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

### //2. В структуру Professor встроены поля структуры University (переходят и все методы тоже)
type Professor struct {
	fullName string
	age      int
	University
}

func main() {
	p := Professor{
		fullName: "Boris Bobroff",
		age:      150,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}
### //3. Вызываем метод структуры University через экземпляр профессора
	p.FullInfoUniversity()
}

## Metod value (Изминение значений структуры по методу)

type Rectangle struct {
	length, width int
}

### //1. Реализуем функцию и метод для подсчет периметра прямогуольника
// ВАЖНО: Все параметры передаем как копии

### //4. При вызове этого метода неважно, будет ли он вызываться у экземпляра или у его ссылки
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

### //5. Данную функцию можно вызывать ТОЛЬКО у копии прямоугольника (но не у его ссылки)
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

### //6. Допустим будет метод, который меняет значение состояния структуры на новое, но этот метод - Value Reciever
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
### 	//2. Создаем экземпляр прямоугольника
	rectangleAsValue := Rectangle{10, 10}
	fmt.Println("Call function for rectangleAsValue:", Perimeter(rectangleAsValue))
	fmt.Println("Call method for rectangleAsValue:", rectangleAsValue.Perimeter())

### 	//3. Создадим указатель на прямоугольник
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call method for rectangleAsPointer:", rectangleAsPointer.Perimeter())
	//Perimeter(rectangleAsPointer) -  Иллюстрация к пункту 5

### 	//7. Вызываем метод SetLength у экземпляра rectangleAsValue
	fmt.Println("Before call method SetLength:", rectangleAsValue)
	rectangleAsValue.SetLength(9999)
	fmt.Println("Aftet call method SetLength:", rectangleAsValue)

### 	//8. Вызываем метод SetLength у ссылки на rectangleAsValue
	rectangleAsPointer.SetLength(999999999)
	fmt.Println("After call method SetLength for &rectangle:", *rectangleAsPointer)
}

### продолжение
type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	//Значение исходное
	rectangleAsValue := Rectangle{10, 10}
	//Ссылка на исходное значение
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call Area function for &rectangle:", Area(rectangleAsPointer))
	fmt.Println("Call Area method for &rectangle:", rectangleAsPointer.Area())

####	//1. Вызываем метод у value - исходного значения
	fmt.Println("Call Area method for rectangle:", rectangleAsValue.Area())
####	//2. Вызываем функцию с параметром value - исходное значение
	//Area(rectangleAsValue)

####	//3. Распечатаем исходный прямоугольник
	fmt.Println("Before changing width:", rectangleAsValue)

####	//4. Вызываю метод SetWidth у &rectangle
	rectangleAsPointer.SetWidth(999)
	fmt.Println("After change via method SetWidth for &rectangle:", rectangleAsValue)

####	//5. Вызов метода SetWidth у rectangle
	rectangleAsValue.SetWidth(888)
	fmt.Println("After change via method SetWidth for rectangle:", rectangleAsValue)
}

## Методы и базовые типы

### //1. Методы для стандартных типов
// В Go встроено куча примитивов (int, int32, string, bool)
//Что если очень хочется дописать к стандартному типу какой-то мтеод?

### //2. Наивная попытка. Это невыоплнимо. Копилятор запрещает добавление новых методов к существующим базовым типам
// func (a *int) IsEven() bool {
// 	if *a%2 == 0 {
// 		return true
// 	}
// 	return false
// }

### // 3. Но мне очень хочется, что делать?
// Создайте новый тип - ваш int и делайте с ним что хотите!
//Для создания нового типа используют конструкцию
type MySuperDuperInt int

func (msdi *MySuperDuperInt) IsEven() bool {
	if *msdi%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := MySuperDuperInt(202)
	num2 := MySuperDuperInt(201)
	fmt.Println(num1.IsEven())
	fmt.Println(num2.IsEven())
	//4. Внутренние преобразования
	// var num3 MySuperDuperInt = MySuperDuperInt(10)
	// var num4 int = int(num3)
}
# Constructors

## //1. Создадим структуру Rectangle
type Rectangle struct {
	length, width int
}

## //2. Создадим конструктор для Rectangle
// Для имен конструкторов в Go договорились использовать функцию с следующим названием:
// * New() если данный конструткор на файл один (в файле присутствует описание только одной структуры)
// * New<StructName>() - если в файле присутсвуют еще какие-то структуры

## //3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него

func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

## //4. Добавим 2 метода
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := NewRectangle(10, 20)
	fmt.Printf("Type as %T and value %v\n", r, r)
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Area:", r.Area())

}

type Circle struct {
	radius float64
}

func NewCircle(newRadius float64) *Circle {
	return &Circle{newRadius}
}

## 5. Отсебятинка Метод для массива + Pointer
type Employee struct {
	name   string
	salary int
	stavka int
}

func NewEmployee(newName string, newSalary int, newStavka int) *Employee {
	return &Employee{newName, newSalary, newStavka}
}

func (e *Employee) FullSalary() int {
	return e.salary * e.stavka
}

func main() {
	db := []*Employee{}
	for i := 1; i <= 50; i++ {
		name := "Employee " + strconv.Itoa(i)
		salary := i*2 + 2
		stavka := 0
		if i%2 == 0 {
			stavka = 2
		} else {
			stavka = 1
		}
		NewEmployeer := NewEmployee(name, salary, stavka)
		db = append(db, NewEmployeer)
	}
	i := 0
	fmt.Print("Name ", db[i].name, " ||Salary ", db[i].salary, " ||Stavka ", db[i].stavka, "\n")
	i = 1
	fmt.Print("Name ", db[i].name, " ||Salary ", db[i].salary, " ||Stavka ", db[i].stavka, "\n")
	f := db[1].FullSalary()
	fmt.Println(f)
}
# Interface
## Part 1
### //1. Интерфейсы
Структура -  явно декларированный заименованный набор СОСТОЯНИЙ.
//Структра , исходя из своего описания, отвечает на вопрос - из ЧЕГО я должен состоять,
// чтобы считаться тем ТИПОМ, который декларируется структурой?
// Структура - описывает ПАТТЕР СОСТОЯНИЯ.

//1. Интерфейсы - явно декларированный набор сигнатур ПОВЕДЕНИЙ (чаще всего в виде набора методов), удовлетворив который,
// можно считаться типом, который декларирует интерфейс.
// Интерфейсы в Go декларируют ПОЛУ-АБСТРАКТНЫЙ тип.
// Отвечает на вопрос - что я должен УМЕТЬ ДЕЛАТЬ, чтобы считаться тем ТИПОМ, который декларирует интерфейс?
// Интерфейс - описывает ПАТТЕРН ПОВЕДЕНИЯ.

### //2. Объявление интерфейса
type FigureBuilder interface {
	//Набор сигнатур методов, которые необходимо реализовать в структуре-претенденте
	//Во-первых , у претендента должен быть метод Area() возвращающий int
	Area() int
	//Во-вторых, у претендента должен быть метод Perimeter() возвращающий int
	Perimeter() int
}

### 3. Декларируем претендентов
#### //3.1 Первый претендент - это прямоугольник
// У него есть 2 метода -
//Area() int
//Perimter() int
//Когда эти методы реализованы , говорят, что RECTANGLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
//RECTANGLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

#### // 3.2 Второй претендент - это окружность
// У нее есть 2 метода -
//Area() int
//Perimter() int
//Когда эти методы реализованы , говорят, что CIRCLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
//CIRCLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {

###	//4. Создадим по 3 экземпляра
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

###	//5. Посчитаем общую площадь этих  фигур
        total1 := r1.Area() + c1.Area()
	fmt.Println(total1)
        //
	rectangles := []Rectangle{r1, r2, r3}
	totalSumAreaOfRectangles := 0
	for _, rect := range rectangles {
		totalSumAreaOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalSumAreaOfCircles := 0
	for _, circ := range circles {
		totalSumAreaOfCircles += circ.Area()
	}

	fmt.Println("Total area is:", totalSumAreaOfRectangles+totalSumAreaOfCircles)

###	//6 . Теперь воспользуемся интерфейсом, указанным выше
        //figures := make([]FigureBuilder, 0, 10)
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3} //Объявляю слайс экземпляров, удовлетворяющих интерфейсу FigureBuilder
	// С другой стороны, кажется, что это слайс - каких-то определенных типов

###	//7. Посчитаем общую площадь
	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("Total area:", total)
###	//8 . Пояснение - так как каждый экземпляр удовлетворяет интерфейсу FigureBuilder (объявляющий пол-абстрактный тип)
	// у каждого из слайса figures можно 100% вызывать метод Area() (который точно вернет int), или Perimter()
	// кототрый тоже 100% вернет int
}

## Part 2

### //1. Описание интерфейса (описание того, что должен уметь претендент)
type BigWord interface {
	IsBig() bool
}

### //2. Наш претендент, которого надо научить делать IsBig() bool
type MySuperString string

### //3. Реализация IsBig() у претендента MySuperString
func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("akj")
	var interfaceSample BigWord // Объявление переменной, типа интерфейса BigWord
	interfaceSample = sample    // присваивание значения для переменной тип BigWord возможно,
	//потому что sample (типа MySuperString ) удовлетворяет интерфейсу BigWord
	fmt.Println("IsBig?", interfaceSample.IsBig())

	//4. Попытка присвоить значение с типажом, неудовлетворяющему интерфейсу
	// var interfaceBadSample BigWord
	// interfaceBadSample = "abcdef" // тип string не имеет реализации метода IsBig , поэтому не удовлетворяет интерфейсу

}

## Part 3

### //1. Объявляем интерфейс, декларирующий поведенческий -паттерн (набор методов под реализацию)
type Worker interface {
	Work()
}

### //2. Объявляем структуру. Данная структура - претендент на удовлетворение интерфейса Worker
type Employee struct {
	name string
	age  int
}

### //3. Реализуем метод Work(), чтобы структура Employee удовлетворяла интерфейсу Worker
func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working!")
}

### //4. А давайте создадим функцию, которая будет принимать аргументы типа Worker и что-то с ними делать?
func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}

### //6. Объявляем структура. Данная структура - второй пренедент на удовлетворение интерфейса Worker
type Student struct {
	name         string
	courseNumber int
}

func (s Student) WantToEat() {
	fmt.Println("student with name", s.name, "want to eat!")
}

func (s Student) Work() {
	fmt.Println("Student with name", s.name, "is working!")
}

func main() {
###	//5. Создадим экземпляр Employee
	emp := Employee{"Bob", 34}
	var workerEmployee Worker = emp // Присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee) // В резульатте видим, что тип интерфейса определяется структурой, его реализующей,
	//а значение интерфейса - это соответственно значение состояний структуры

###	//7. Создадим экземпляр Student
	std := Student{"Mike", 19}
	var workerStudent Worker = std
	workerStudent.Work()
	Describer(workerStudent) // Приянл внутренний тип Student, а значение - равно значению полей экземпляра

###	//8. Созаддим набор тех, кто умеет  Work()
	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker) // Данная функция вызывается у разных экземпляров благодря тому, кто для ее вызова
		//экземпляру нужно удовлетворить некому контракту (поведенческому паттерну). Если структура экземпляра поддерживает
		// данный паттерн - то у экземпляра 100% можно вызвать все необходимые для этого методы.
	}
}

## Part 4

### //1. А что если создать интерфейс, в котором в принципе нет никаких требований к поведению?
type Empty interface {
}

### //2. Вопрос - кто удовлетворяет этому интерфейсу? Если интерфейс пустой - то ему удволетворяет вообще кто угодно.

### //3. Реализуем функцию, которая может принимать кого угодно
func Describer(pretendent interface{}) {
	fmt.Printf("Type = %T and value %v\n", pretendent, pretendent)
}

type Student struct {
	name string
}

### //4. Type Assertion
func SemiGeneric(pretendent interface{}) {
	val, ok := pretendent.(int) // Пытаюсь проверить, что претендент - типа int
	fmt.Println("Value:", val, "Ok?:", ok)
}

func main() {
	strSample := "Hello world!"
	//4. Передача параметра без присваивания в промежуточную переменную
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vova"})

###	//5. Работа с полу-дженериком
	SemiGeneric(10)
	SemiGeneric(Student{"Fedya"})
	SemiGeneric("Hello world")
}

## Part 5

 func DoSomething(pretendent interface{}) {
	switch pretendent.(type) { // Пытаемся извлечь нижележащий тип
	case string: // если нижележащий тип string
		fmt.Println("This is string!")
	case int:
		fmt.Println("This is int!")
	default:
		fmt.Println("Unknown type! But i'm working!")
	}
}

func main() {
	DoSomething(10)
	DoSomething("Hello world!")
	DoSomething(func(a, b int) int { return a + b })
}

## case Describer

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("%s and %d y.o\n", std.name, std.age)
}

func TypeFinder(i interface{}) {
	switch v := i.(type) { //Присовим переменной v значение лежащие под предполагаемым интерфейсом
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer: // Вывод - с типом интерфейса можно сравниваться точно так же, как и с любым другим типажом языка
		// это как раз и говорит о том, что интерфейсы - полу-абстрактные типы.
		fmt.Println("I'm describer!")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Alex", 23}
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(nil)
	TypeFinder(std)
}

## Pointer and interface

//1. Вопрос - имеет ли для интерфейса значение тот факт, что метод, реализованный для претендента
// , в качестве получателя принимает значение или указатель?

//0. В Go принято называть интерфейсы , с окончанием на `er` (Describer, Worker, Pooller,....)
type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("Student name: %s and age %d\n", std.name, std.age)
}

type Professor struct {
	name string
	age  int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor name %s and age %d\n", prof.name, prof.age)
}

func main() {
	var descr1 Describer
	stud1 := Student{"Alex", 23}
	descr1 = stud1 //Student реализует интерфейс Describer
	descr1.Describe()

	stud2 := &Student{"Bob", 21} // Поскольку экземпляр -ссылка, разыменовать ее имеет право кто-угодно (в том числе и компилятор)
	descr1 = stud2
	descr1.Describe()

	var descr2 Describer
	prof1 := &Professor{"Viktor Wann", 72}
	descr2 = prof1 // &Professor реализует интерфейс Describer
	descr2.Describe()

	prof2 := Professor{"Bob Brown", 65}
	prof2.Describe() // Здесь ссылку на &prof берет компилятор
	descr2 = prof2   //Professor не реализует интерфейс Describer
	// Дело в том, что сам по себе интерфейс - не референсный тип
	// Выливается это в следующее следствие:
	// Когда мы работаем с обычным методом, то взять референс на экземпляр  может компилятор
	// Но когда мы пытаемся сделать это через интерфейс - в нем в принципе комплятор не видит никаких ссылок!
}

## Обьединение Interfases

### //0. Интерфейсы (с точки зрения ООП) - увеличивают уровень абстракции вашего кода
// Засчет увеличения уровня абстракции - можно решать много сторонних проблем, связанных с поддержкой
//понимание и реюзабельностью кода.
// С другой стороны - добавление интерфейсов не решают проблему уменьшения абстрактности.

### //1. Что делать, если хочется скрестить 2 интерфейса и создать единый уровень абстракции в коде?
package main

import "fmt"

type PerimeterCalculator interface {
	Perimeter() int
}

type AreaCalculator interface {
	Area() int
}

### //2. Соберем новый интерфейс из двух старых
type FigureFeatureCalculator interface {
	PerimeterCalculator // встраиваем интерфейсы
	AreaCalculator
	// Наслдеование интерфейсов
	//Perimeter() int
	//Area() int
}

type Rectangle struct {
	a, b  int
	color string
}

### //2. Реализуем интерфейс PerimeterCalculator
func (r Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

### //3. Реализуем второй интерфейс AreaCalculator
func (r Rectangle) Area() int {
	return r.a * r.b
}

func main() {
	var pCalc PerimeterCalculator
	fmt.Printf("Value %v and Type %T\n", pCalc, pCalc)
	var aCalc AreaCalculator
	rect := Rectangle{10, 20, "green"}
	pCalc = rect // Стурктура Rectangle удовлетворяет интерфейсу PerimeterCalculator
	fmt.Printf("Value %v and Type %T\n", pCalc, pCalc)
	fmt.Println("Perimeter:", pCalc.Perimeter())

	aCalc = rect // Структура Rectangle удовлетворяет интерфейсу AreaCalculator
	fmt.Println("Area:", aCalc.Area())

	var ffCalc FigureFeatureCalculator
	ffCalc = rect // Структура Rectangle удовлетворяет FigureFeatureCalculator
	fmt.Println(ffCalc.Area())
	fmt.Println(ffCalc.Perimeter())

}

## Zero Value

### //0. Почему интерфейс -полу-абстрактный тип в Go?

### //1. Создадим интерфейс Ездилка
type Rider interface {
	Ride()
	Gas()
	Stop()
}

func main() {
###	//2. Создаю экземпляр ездилки
	var r Rider   // ZeroValue - nil, ZeroType - nil
	if r == nil { // Попробуем сравнить сравнить с nil
		fmt.Printf("r is nil. Value %v and type %T\n", r, r)
	}

	r.Ride() // Здесь будет паника, т.к. у экземпляра интерфейса можно вызвать метод Ride()
	//но т.к. значение, которое там лежит - nil - получается nil.Ride()
	//Мораль - если код падает с memory-dereferncing - в 99% случаев - это попытка обратиться к
	//экземпляру интерфейса без присвоенного претендента!
}
# Пакетирование
##  go mod init
func main() {
        // go env // показывает путь ко всем ресурсам (пакеты)
        // GOROOT=C:\Program Files\Go // - стандартные библиотеки
        // GOPATH=C:\Users\Mvmir\go // - левые библиотеки

	a := 25
	b := 5

	fmt.Println(Add(a, b))
	fmt.Println(Sub(a, b))
	fmt.Println(Mult(a, b))
	fmt.Println(Div(a, b))

	//4. Для того, чтобы запустить все файлы в пакете надо выполнить:
	//* go run main.go calculator.go -- запускает выполнение пакета
	//* go build main.go calculator.go -- создаёт main.exe файл //./main - Запустить файл
	//* go install main.go calculator.go -- хз что оно делает
	//* ./main - Запустить файл


	//5. Создание стороннего ПАКЕТА rectangle
	// папка внутри папки pack с файлом rectangle.go
	//  команда - go mod init pack -- иннициализация пути к rectangle - создаёт файл go.mod
        //* go build - создаёт файл pack.exe // ./pack
        //* ./pack - Запустить весь пакет

	// 8 Import Rectangle
	/*import (
		"fmt"
		"pack/rectangle"
	)*/
            |пакет!   | метод
	r := rectangle.New(10, 20, "green")
	fmt.Println("Perimeter:", r.Perimeter())

	// 9 Маленькие заглавные буквы имён ВИДНЫ ТОЛЬКО ВНУТРИ ПАКЕТА
	newR := rectangle.Rectangle{
		A:     10,
		B:     7,
		Color: "red",
	}
	fmt.Println(newR)
	//10 go build обновляет инфо
	//go install куда то складывает бинарник?.. и удаляет файл pack.exe

	//11 Всегда используйте
	//  команда - go mod init pack -- иннициализация пути к rectangle - создаёт файл go.mod
}

## Заглушенный импорт

import (
	_ "Lec23/circle" //_ оставляет ссылку на пакет без использования пакета
	"Lec23/rectangle"
	"fmt"
)

//1. Функци init() - данная функция отрпбатывает единожды при первом импортировании пакета
//2. Данных фукнций в пакете может быть несколько штук! (не в одном модуле, т.к. в одном модуле нельзя создать
// более одной функции с каким-то определенным именем)

//3. init() вызывает в момент инициализации пакета:
//* Данный процесс выглядит следующим образом:
// ** сначала компилятор смотрит на содержимое пакета
// ** затем компилятор смотрит на пути импорта (если что-то импортируется, компилятор уходит туда)
// ** затем компилятор инициализурет переменные уровня пакета
// ** затем компилятор запускает функцию init() для данного пакета
// ** повторяет данную процедуру для всех пакетов проекта
// ** после чего вызывается функция main()

//4. Что произойдет , если запустить go run main.go
// * Сначала смотрим в main.go на предмет синтаксических ошибок ,но ничего не инициализируется
// * Затем импорты : сначала импоритруем Lec23/rectangle
// ** Компилятор идет в rectangle
// ** Смотрим в пакет на предмет синтаксических ошибок
// ** Затем импорт fmt
// ** Затем инициализируем переменные уровня пакета
// ** Затем запускаем функцию init() пакета rectangle
// ** Затем подружаем (определяем) все имена пакета rectangle
// **Функции main тут нет, возвращаемся назад
// * Пытаемся импоритровать fmt (т.к. он уже был импортирован одним из пакетов - повторноый импорт не требуется)
// * Инициализируем переменные уровня пакета main
// * Запускаем функцию init() в main
// * Затем определяем имена (тут дополнительных имен нет, тут ничего не делаем)
// * Затем запускаем функцию main()

//5. Все импорты (вне зависимости , стандартные или пользовательские) сортируются по алфавиту
func init() {
	fmt.Println("Init function for main package!")
	fmt.Println("Name:", name, "Age:", age)
}

var (
	name string = "Bob"
	age  int    = 99
)

func main() {
	r := rectangle.New(10, 10)
	fmt.Println(r)
}

## Тестировка

import (
	"log"
	"testing"
)

//1. Файл с модульными тестами приянто называть:
// * <script_name>_test.go
// * <package_name>_test.go

//2. Для того, чтобы тестировать функции (методы, стркутуры, интерфейсы и т.д.)
// На каждый юнит создаем по своей тестирующей функции (Test)
// Приянто каждую такую функию начинать со слова Test....
func TestAdd(t *testing.T) {
	//1. 1-ый test-case
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10, 30) fail. expected %d, got %d\n", 30, res) // log.Fatal провоцирует завершение работы кода
	}

	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add(30, 30) fail. expected %d, got %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {}

func TestMult(t *testing.T) {}

//3. Для запуска тестов используем команду go test
// Детально : go test -v

//4. Coverage (покрытие) - показывает сколько % кода покрыто модульными тестами
// go test -cover -v
// 75~80% coverage - этого бывает более чем достаточно!

//5. Напоследок :
// Все что начинается с слова Test - будет запущено командой go test
// В Go приянто, что создается 1 модуль с тестами на весь пакет (вне зависимости от количества модулей в нем)
// Не тестируйте Getattr/Setattr в общем пайплайне (только специфика)
// Обязательно посмотрите, как происходит связка с CI для Go (TravisCI/CircleCI)

## Defer

//1. DEFER - оператор отложенного ВЫЗОВА функции/метода.

//2. Создадим отложенную фукнцию.

//3. С входными параметрами

func CheckDBCloseConncetion(a int) {
	fmt.Println("Check started......Value numIP in deferred:", a)
	fmt.Println("Check done! Connection refused!")
}

//4. Если функция принимает входные параметры и данная функция используется в связки с defer
// то :
// параметры расчитываются в момент передачи их в функцию
// А вызов функции с уже давно рассчитанным параметром осуществляется в момент
// завершения вышележащей функции

//5. В какой момент defer вызывается?
func OuterFunc() int {
	defer fmt.Println("I'm deferred print function!")
	fmt.Println("OuterFunc started!")
	var result = 10
	fmt.Println("OuterFunc finished. Ready to return value!")
	return result
}

func main() {
	defer fmt.Println("Step 1 defer")
	defer fmt.Println("Step 2 defer")
	defer fmt.Println("Step 3 defer")
	defer fmt.Println("Step 4 defer")

	var numIP = 10
	p := &numIP
	//defer CheckDBCloseConncetion(numIP) // defer означает, что данная функция будет вызвана при завершении main() с параметром,
	//значение которого расчитывается на момент 25-ой строки.
	*p++
	fmt.Println("Value numIP in main():", numIP)
	fmt.Println("Main function started")
	fmt.Println("Main function ended")
	fmt.Println("Value form OuterFunc on main side is:", OuterFunc())
}

## Panica

// В Go существует 2 механизма сигнализирования анаомального поведения
// 1-ый механизм это ошибки Error (ЯВЛЯЕТСЯ КАНОНИЧНЫМ ИСПОЛНЕНИЕМ НА СЛУЧАЙ НЕНОРМАЛЬНОГО ПОВЕДЕНИЯ)
// 2-ой механизм - это паника (не лучший вариант, так как приводит сразу к аварийному завершению, и в принципе
// мог быть обычной ошибкой)

package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
)

func funcWuthError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}
	return "", errors.New("THIS IS ERRRO!")
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied:", r)
		debug.PrintStack()
	}
}

func panicExample(firstName *string, lastName *string) {
	defer PanicRecover() // даже в случае возникновения паники - первым дело будут вызваны все deferred функци.
	if firstName == nil {
		panic("runtime errror: firstname can not be nil!")

	}

	if lastName == nil {
		panic("runtime error: lastname can not be nil!")
	}

	fmt.Println("Full name:", *firstName, *lastName)
}

func main() {
	numLiteral := "12"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		fmt.Println("can not convert this value to int:", err)
		return
	}
	fmt.Println("Convertion done:", num)

	var name = "Bob"
	panicExample(&name, nil)

	ans, err := funcWuthError(5)
	if err != nil {
		fmt.Println("not use odd values", err)
		return
	}
	fmt.Println(ans)
}

## Многопоточность

//Немного общих слов

//Go из коробки - язык конкурентный , а не параллельный.

// 1. Что такое конкурентность?
// Конкурентность - это подвид многопоточного исоплнения программ, где различные задачи ("работники") конкурируют за ресурсы.
// Конкурентность навязывается различными факторами :
// * приорететы исполнения
// * простой ресурсов
// * ручная передача

// 2. Пример конкурентности
// Екатерина бегает утром. Во время пробежки развязываются шнурки. Екатерина останавливается. Завязывает шнурки.
// Затем продолжает пробежку.
// Это классический пример конкурентности.

// 3. А параллелизм?
// Параллелизм - это подвид многопоточного исполнения программ, в котором множество задач ("работников") используют ресурсы ОДНОВРЕМЕННО.

// 4. Пример параллелизма
// Екатерина также бегает по утрам. Но в этот раз она еще и слушает музыку. В этот раз Екатерина ОДНОВРЕМЕННО и бегает и слушает музыку.

// 5. В чем разница и что лучше?
// Рассмотрим простой пример - браузер.
// Когжа мы заходим на какую-нибудь страницу : должно быть выполнено 2 действия
// * загрузка html страницы (файла)
// * отрисовка (рендеринг) в окне браузера

// Если данные задачи выполняются конкуретно, то сначала вы загрузите необходимый объем файлов , а уже затем выполните отрисовку.
// Процессор в этой ситуации будет осуществлять переключение контекста (context switch) в нужный момент (по завершении загрузки) и
// результат будет ожидаем.

// С другой стороны, если бы эти 2 задачи выполнялись параллельно, результат был бы немного шокирующим и непредсказуемым.

// 6. Как реализована конкуретность в языке?
// В Go поддержк конкуретности реализована с исползованием под-программ (со-программ) , т.н. "горутин" (или corutines/gorutines)

//7 . Горутина, это кто?
// Горутина - это фукнция или метод, которая запускает другие функции/методы или выполняет какие-то действия.
// Горутина , с технической точки зрения, может восприниматься как легковесный тред. На одном системном потоке может одновременно
// находиться огромное количество конкурирующих за ресурсы горутин.

// 8. Преимущества горутин над классическими тредами
// * горутина легковесная (размер горутины в миллионы раз меньше, чем размер классического треда в С++/Java)
// * исопльзование большого количество горутин занимает меньшее количество потоков ОС (в отличе от Java/C++, где отдельный тред
// требует выделения отдельного потока в ОС)
// * горутины могут общатсья друг с другом используя каналы

### ------------------------------------------------
package main

import "fmt"

//1. Данная функция будет запущена в качестве горутины.
// Важно: горутины никогда ничего не возвращают через явное использование return
func newGoRoutine() {
	fmt.Println("Hey, I'm new Gorutine!")
}

//2. функция main - на самом деле тоже горутина.
// Особенность в том - что если эта горутина завершается - все остальные запущенные убиваются сразу!
func main() {
	go newGoRoutine() // в этот момент происходит формирование запроса на вызов функции в отдельной горутины.
	// соответственно код основной горутины main продолжает сразу же выполняться и не ждет завершения newGoRoutine()
	fmt.Println("Main goroutine work!")
	//Запустим код и....

}
### ----------------------------------------------------
package main

import (
	"fmt"
	"time"
)

func newGoRoutine() {
	fmt.Println("Hey, I'm new Gorutine!")
}

func main() {
	go newGoRoutine()
	time.Sleep(1 * time.Second) // немного тормозим выполнение main горутины, таким образом даем время для того,
	// чтобы newGoRoutine успела выполниться
	fmt.Println("Main goroutine work!")
}
### -----------------------------------------------------
package main

import (
	"fmt"
	"time"
)

//1. Запустим несколько горутин и посмотрим, как они бьются за ресурсы

// 2. Определим первую горутину
func printEvenNumbers() {
	for i := 1000; i < 1020; i += 2 {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

// 3. Определим вторую горутину
func printOddNumbers() {
	for i := 1; i < 20; i += 2 {
		time.Sleep(450 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

// 4. Определим main горутину
func main() {
	go printEvenNumbers()               // сразу идем дальше, запуск функции будет происходит в отдельной горутине
	go printOddNumbers()                // также идем дальше, запуск функции будет происходит потом
	time.Sleep(7000 * time.Millisecond) // тормозим основную горутину, чтобы остальные успели что-то сделать
	fmt.Println("main goroutine died")
}

//5. Таким образом горутины работают следующим образом : (нарисуем прямоугольник с палочками!)
### ---------------------------------------------------------
package main

import "fmt"

//1. Каналы - средство для коммуникации между горутинами.
// Каналы можно рассматривать как соединетильные трубы, через которые горутины между собой общаются (аналогично тому,
// как вода течет по трубам, данные передаются через каналы)

//2. Объявление канала.
// Каналы по умолчанию имеют zeroValue - nil. Поэтому их создают через фукнцию make.

func main() {
	var a chan int // объявляем канал, через который будут передаваться данные типа int
	if a == nil {
		fmt.Println("channel is nil, Let's define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
}
### -----------------------------------------------------------
package main

import "fmt"

//1. Для отправки данных в канал a (chan int) используем синтаксис
//a <- dataInt

//2. Для получения данных из канала используем синактсис
// dataInt := <- a

//3. Отправка и получения данных из канала - блокирующая операция!
// Это означает, что если данные отправлены в канал, то выполнение текущей программы останавливается до тех пор, пока с другой
// стороны из этого канала кто-то не считает данные.

// Аналогично и в обратную сторону. Если кто-то читает из канала, то выполнение текущей программы (горутины) останавливается до тех пор,
// пока кто-то в этот канал не отправит данные.

//4. Пример использования каналов.
func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I'm new Gorutine!")
	done <- true
}

func main() {
	done := make(chan bool)
	go newGoRoutine(done)
	<-done // в этой точке выполнение main горутины останавливается до тех пор, пока в канал кто-нибудь не запишет данные!
	fmt.Println("Main goroutine work!")
}
### --------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

//1. Немного перепишем последнюю программу, чтобы лучше увидеть как устроен процесс блокирования

func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I'm newGoRoutine and I'm going sleep!")
	time.Sleep(4 * time.Second)
	fmt.Println("newGoRoutine awake and going to send data to channel")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("I'm main goroutine and I wanna call newGoRoutine")
	go newGoRoutine(done)
	<-done
	fmt.Println("Ok, Main goroutine recieved data and gonna die!")
}
### -------------------------------------------------------------------
package main

import "fmt"

//1. Создадим чуть более полезную программу, которая будет делать следующие действия:
// берем число, например 456
// (4^2 + 5^2 + 6^2) + (4^3 + 5^3 + 6^3)
// Подсчитаем сумму квадартов цифр и сумму кубов, а затем сложим полученные результаты
//Делать будем следующим образом:
// * main gorutine получает число и вызывает 2 другие горутины, по итогу, получив от них результаты,
// она просто их сложит и выведет на консоль
// ** squaresGoRoutine - будет запущена main и подсчитает сумму квадратов всех цифр, результат положит в канал
// ** cubesGoRoutine - будет запущена main и подсчиатет сумму кубов всех цифр , результат полужит в канал

func squaresGoRoutine(num int, squareChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10 // остаток от деления на 10 // 456%10 -> 6 // 45%10 -> 5 // 4%10 ->4
		sum += digit * digit
		num /= 10
	}
	squareChan <- sum // <- типа return
	fmt.Println("squareChan ferst")
}

func cubesGoRoutine(num int, cubeChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cubeChan <- sum
	fmt.Println("cubeChan ferst")
}

func main() {
	number := 7835343
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go squaresGoRoutine(number, squareChan)
	go cubesGoRoutine(number, cubeChan)
	squaresSum, cubesSum := <-squareChan, <-cubeChan
	fmt.Println("Total result is:", squaresSum+cubesSum)
}
### ------------------------------------------------------------------------------------
package main

//1. Deadlock - ситуация, когда кто-то пишет в канал НО ИЗ НЕГО НИКОГДА НИКТО НИЧЕГО НЕ ПРОЧИТАЕТ, или когда кто-то читает из канала
// НО В НЕГО НИКТО НИКОГДА НЕ ЗАПИШЕТ!

// По сути ситуация означает, что для отправляющей стороны отсутствует получатель (с другой стороны никто не ждет данных). И наоборот.

func main() {
	ch := make(chan int)
	ch <- 10
	// <-ch
}
### ---------------------------------------------------------------------------------------
package main

//1. Каналы могут иметь направление.
// То что инициализировалось до этого момента инициализировало ДВУНАПРАВЛЕННЫЙ КАНАЛ (в него можно и писать и из него можно читать)
// Можно создать канал только на отправку:
// sendChan := make(chan<- int)

// Можно создать канал только на чтение
// readChan := make(<-chan int)

func sender(sendChan chan<- int) {
	sendChan <- 10 // Тут все ок
}

func main() {
	sendChan := make(chan<- int)
	go sender(sendChan)
	<-sendChan // тут не ок, т.к. канал только на отправку данных, но не чтение.
}

//2. Использование однонаправленных каналов никак не сказывается на производительности, а служит лишь для логического разделения кода
### ---------------------------------------------------------------------------------------------
package main

import "fmt"

//1. Закрытие каналов и итерирование
// Со стороны получателя можно использоваться синтаксис
// val, ok := <- ch
//где val - значение помещенное в канал отправителем
// ok - true/false в зависимости от того, открыт канал или уже закрыт отправителем.
// если канал закрыт то в val будет помещено zeroValue для типа канала

func generator(ch chan int) {
	for i := 0; i < 25; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generator(ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Recieved from channel", val)
	}

	// Конструкцию можно упростить и использовать
	// for val := range ch {
	// 	fmt.Println("Recieved from channel:", val)
	// }
}
### -------------------------------------------------------------------------------------
package main

import "fmt"

//1. Попробуем решить старую задачу с подсчетом суммы квадратов и кубов полученного числа с использованием закрытия каналов

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch) // своя горутина для генерации цифр
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
	close(squareop)
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch) // еще одна горутина для генерации цифр
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
	close(cubeop)
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
### ------------------------------------------------------------------------
package main

import "fmt"

//1. Буферезированный канал - канал с буфером, в который можно напихать со стороны отправителя не 1 пак данных, а столько, сколько
// позволяет в буфер.

// ch := make(chan int, capacityIntValue)

func main() {
	ch := make(chan string, 5) // Создадим канал вместимостью 2
	ch <- "Bob"                // Не блокиуремся , т.к. можно запихнуть в буфер еще 4 элемента
	ch <- "Alex"               // Не блокируемся, т.к. можно еще запихнуть в буфер 3 элемента
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
### --------------------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

//1. Как блокируется буферезированный канал?
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

//2. Как только буфер заполняется - канал блокируется до тех пор, пока не будет освобождено место (буфер может быть освобожден не до конца!)

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}

}

// 3. Длина и вместимость канала.
// Длина канала len(ch) - сколько сейчас элементов в канале
// Вместимость cap(ch) - какой размер буфера у канала
### -----------------------------------------------------------------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

//1. Еще один инструмент для оркестрирования горутинами - это WaitGroup
// По сути WaitGroup - это счетчик горутин.
// Когда горутина запускается делается WaitGroup++
// Когда горутина завершается делается WaitGroup--
//Таким образом когда WaitGroup == 0 делаем вывод, что все горутины отработали!

//Пример

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //WaitGroup--
}

func main() {
	no := 5
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) // WaitGroup++
		go process(i, &wg)
	}
	wg.Wait() // if WaitGroup == 0 ? До тех пор, пока это условие не выполнено - мы заблокированы в данной строке для main горутины
	fmt.Println("All go routines finished executing")
}
### --------------------------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

//1. Select - это инструмент, позволяющий выбирать из множества канальных операций (чтение/запись) для множества каналов.
// Если из 10 каналов что-то пришло в один - select выбирает его
// Если из 10 каналов что-то пришло сразу в два и более - select выбирает случайный

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2: // выбирается этот кейс, т.к. в этот канал будут помещены данные быстрее
		fmt.Println(s2)
	}
}
### -------------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

//1. На практике select чаще всего используется для того, чтобы предпринимать какие-то действия,
// пока в каналы еще не пришли данные

//Пример
func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
### ----------------------------------------------------------------------
package main

//1. Select как инструмент защиты от deadlock
// Добавление default страхует от появление deadlock в ходе выполнения и берет работу на себя (попробуйте убрать default и все умрет)

import "fmt"

func main() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")

	}
}
### -----------------------------------------------------------------------
package main

//1. Опрос каналов. Как было сказано ранее - в случае если готовы более чем один канал - выбирается случайный.
import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
### -----------------------------------------------------------------------
package main

//1. Мьютексы - средства защиты от "Resource Sharing"
// Что это такое?
// Во время работы конкурентных программ, главной точкой является тот факт, что множества горутин
// не должны одновременно использовать какой-то общий экземпляр (файл, переменную, бд) для одновременных модификаций.

// 2. Пример RS
// Допустим у нас код
//     x = x + 1
// Пока работает 1 горутина - проблем никаких нет и иннкремент стандартный
// Теперь представим что горутин 2
// ПОСКЛЬКУ ДЛЯ ОБЕИХ ГОРУТИН ДАННЫЙ РЕСУРС БУДЕТ ИСПОЛЬЗОВАТЬСЯ КОНКУРЕТНО, ОЖИДАТЬ ЧТО ПО ИТОГУ РАБОТУ ОБЕИХ ГОРУТИН
// (При начальном X=0) X будет равен 2 - НЕЛЬЗЯ.
//*  Первая и вторая горутина могут начать работать с парметром x =0 (вторая не дождется , пока первая увеличит X на единицу)
// * В итоге будет X = 1
//* Первая горутина начнает работать, увеличивает x на единицу , но не успевает присвоить его
//* Первая горутина начинает работать, завершается, потом стартуер вторая (этот вариант оптимистичный, но сработает с вероятность 1/3!)

// Для того, чтобы избежать этой пробелмы использую мьютексы (мьютекс блокирует ресурс до тех пор , пока его не осводит одна из горутин)
// В таком случае код с инкрементом выглядел бы как :
// mutex.Lock()
// x = x + 1
// mutex.Unlock()

import (
	"fmt"
	"sync"
)

// 2. Пример. Возникновение RS часто именую как Race Condition (состояние гонки)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
### --------------------------------------------------------------------------------
package main

// Разрешим состояние гоник при помощи введение пары мьютексов!
import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
### ------------------------------------------------------------------------------
package main

//Также состояние гонки может быть разрешено через использование каналов (т.к. каналы это более детальный инстурмент
// коммуницирования)
import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
### ----------------------------------------------------------------------------------
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
# Format flag

ar integer int64 = 32500
var floatNum float64 = 22000.456

func main() {

	// Обычный способ вывода десятичного числа
	fmt.Printf("%d \n", integer)

	// Всегда показывает знак
	fmt.Printf("%+d \n", integer)

	// Вывод с другим основанием x -16, o-8, b -2, d - 10
	fmt.Printf("%x \n", integer)
	fmt.Printf("%#x \n", integer)

	// Отступ перед нулями
	fmt.Printf("%010d \n", integer)

	// Оставляет отступ с пробелами
	fmt.Printf("% 10d \n", integer)

	// Отступ с правой стороны
	fmt.Printf("% -10d \n", integer)

	// Вывод вещественного значения
	// с плавающей запятой
	fmt.Printf("%f \n", floatNum)

	// Вещественное число
	// с ограниченной точностью = 5 (после запятой)
	fmt.Printf("%.5f \n", floatNum)

	// Вещественное число
	// в научной заметке
	fmt.Printf("%e \n", floatNum)

	// Вещественное число
	// %e для крупной экспоненты
	// или %f в противном случае
	fmt.Printf("%g \n", floatNum)

}
# Byte string

 //13 - начало строки
	//46 - точка
	//32 - пробел
	//10 - ентер

# 