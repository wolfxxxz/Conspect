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
 