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
