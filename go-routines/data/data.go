package data

import (
	"fmt"
	"log"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{0, "Dune", false},
	{1, "Ready Player One", false},
	{2, "1989", false},
	{3, "Foundation", false},
	{4, "Wonder", false},
	{5, "Mist born I", false},
	{6, "Mist born II", false},
	{7, "Mist born III", false},
	{8, "Mist born IV", false},
	{9, "Mist born V", false},
}

func findBook(id int, mutex *sync.RWMutex) (int, *Book) {
	mutex.RLock()
	var found = false
	var index = -1
	var book *Book

	for i := 0; i < len(books) && !found; i++ {
		if id == books[i].ID {
			found = true
			book = books[i]
			index = i
		}
	}
	mutex.RUnlock()
	return index, book
}

func FinishBook(id int, mutex *sync.RWMutex) bool {
	i, book := findBook(id, mutex)
	if i < 0 {
		return false
	}
	mutex.Lock()
	book.Finished = true
	books[id] = book
	mutex.Unlock()
	log.Printf("Finished book: %s\n", book.Title)
	return true

}

func ToString() {
	for _, book := range books {
		fmt.Printf("%s[%v]->%v\n", book.Title, book.ID, book.Finished)
	}
}
