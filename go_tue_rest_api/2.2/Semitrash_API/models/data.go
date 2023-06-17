package models

var DB []Book

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		Id:            1,
		Title:         "Lord of the Rings. Vol.1",
		YearPublished: 1978,
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkin",
			BornYear: 1982,
		},
	}
	DB = append(DB, book1)
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool
	for _, b := range DB {
		if b.Id == id {
			book = b
			found = true
		}
	}
	return book, found
}

// нужно заменить oldBook на newBook v DB
func ReWriteBook(id int, book Book) {
	book.Id = id
	for i, b := range DB {
		if b.Id == id {
			DB[i] = book
		}
	}
}

// Удалить book from DB
func DelBookFromDB(f int) {
	for i, _ := range DB {
		if i != len(DB)-1 && i >= f-1 {
			DB[i] = DB[i+1]
		}
		if i == len(DB)-1 {
			DB = DB[:i]
		}
	}
}
