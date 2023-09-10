package data

import "log"

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "Ready Player One", true},
	{3, "1989", false},
	{4, "Foundation", false},
	{5, "Wonder", false},
	{6, "Mist born I", false},
	{7, "Mist born II", false},
}

func findBook(id int) (int, *Book) {
	for _, book := range books {
		if book.ID == id {
			return book.ID, book
		}
	}
	return -1, nil
}

func finishBook(id int) bool {
	id, book := findBook(id)
	if id < 0 {
		return false
	}
	book.Finished = true
	books[id] = book
	log.Printf("Finished book: %s\n", book.Title)
	return true

}
