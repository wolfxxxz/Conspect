package main

import "fmt"

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
