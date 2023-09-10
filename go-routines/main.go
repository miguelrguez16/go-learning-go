package main

import (
	"fmt"
	"go-routines/data"
	"log"
	"math/rand"
	"sync"
	"time"
)

//var wg = sync.WaitGroup{}

func main() {
	log.Println("Application start")
	tStart := time.Now()

	var wg = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, mutex)
	}

	wg.Wait()
	data.ToString()
	log.Printf("Total time %v ms", time.Since(tStart).Milliseconds())

	log.Println("Application End")
}

func readBook(id int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	data.FinishBook(id, mutex)
	delay := rand.Intn(500)
	fmt.Printf("GOROUTINE read book #%d with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
