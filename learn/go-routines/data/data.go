package data

import (
	"fmt"
	"log"
	"strings"
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
	{6, "El Silmarillon", false},
	{7, "Ubik", false},
	{8, "La Sombra del vinto", false},
	{9, "En las Montañas de la locura", false},
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
	log.Printf("Finished book: %s [%v]\n", book.Title, book.ID)
	return true
}

func ToString() string {
	var sb strings.Builder
	for _, book := range books {
		sb.WriteString(fmt.Sprintf("%s [%v]->%v\n", book.Title, book.ID, book.Finished))
	}
	return sb.String()
}
