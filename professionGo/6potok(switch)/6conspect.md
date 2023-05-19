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
