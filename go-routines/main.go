package main

import (
	"fmt"
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

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go show(i, wg)
	}

	wg.Wait()

	log.Printf("Total time %v ms", time.Since(tStart).Milliseconds())
	log.Println("Application End")
}

func show(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("GOROUTINE #%d with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
