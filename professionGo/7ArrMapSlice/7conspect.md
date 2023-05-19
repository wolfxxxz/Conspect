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
### String = []rune(string) => string([]rune) - Круговорот стринг в байт и ...
func main() {
	var wordString = "Бибизянка Dusja"
	var wordRune []rune = []rune(wordString)
	fmt.Println(string(wordRune[2:14])) //бизянка Dusj
}