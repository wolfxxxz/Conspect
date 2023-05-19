package main

import "fmt"

type Chair struct {
	category string
	name     string
	leg      int
	price    float64
}

func NewChair(newName string, newLeg int, newPrice float64) Chair {
	return Chair{"Chair", newName, newLeg, newPrice}
}

type Table struct {
	category string
	name     string
	leg      int
	price    float64
	id       int
}

func NewTable(newName string, newLeg int, newPrice float64, id int) *Table {
	return &Table{"Table", newName, newLeg, newPrice, id}
}

// -------------------------------------------------------------
func (t *Table) GetName() (name string) {
	return t.name
}
func (t *Table) GetCathegory() (categoty string) {
	return t.category
}
func (t *Chair) GetName() (name string) {
	return t.name
}
func (t *Chair) GetCathegory() (categoty string) {
	return t.category
}

// -------------------------------------------------------------
type onlySliceOrMapOrArr interface {
	GetName() string
	GetCathegory() string
	QuantityLeg(int) int
}

func main() {
	table1 := NewTable("stul", 4, 22.5, 1)
	chair1 := NewChair("ThreeLeg", 3, 12.0)
	products := map[string]onlySliceOrMapOrArr{
		table1.name: table1,
		chair1.name: &chair1,
	}

	for key, v := range products {
		switch item := v.(type) {
		case onlySliceOrMapOrArr:
			fmt.Println("Name:", item.GetName(), item.GetCathegory(), item.QuantityLeg(1))
		default:
			fmt.Println(key)
		}
	}

}

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

// ----------------------
func NewRoom(c Chair, t *Table) *Room {
	return &Room{c, t}
}

// Or
/*
func NewRoomChairPlTable(newChLeg, newTabLeg int, newChPrice, newTabPrice float64) *Room {
	return &Room{NewChair(newChLeg, newChPrice), NewTable(newTabLeg, newTabPrice)}
}*/

// ---------------------
func (r *Chair) QuantityLeg(quantoty int) (total int) {
	total = r.leg * quantoty
	return
}

//------------------------------

func (t *Table) QuantityLeg(quantity int) (total int) {
	total = t.leg * quantity
	return
}

//--------------------------------------
