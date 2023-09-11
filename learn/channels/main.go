package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	// create a new channel, type string
	channelIds := make(chan string)
	fakeChannel := make(chan string)
	closeChannel := make(chan int)

	wg.Add(3)
	go generateIDS(channelIds, wg, closeChannel)
	go generateFakeIDs(fakeChannel, wg, closeChannel)
	go logIDs(channelIds, fakeChannel, wg, closeChannel)
	wg.Wait()
}

// generateFakeIDs
func generateFakeIDs(fakeChannel chan<- string, wait *sync.WaitGroup, closeChannel chan<- int) {
	for i := 0; i < 50; i++ {
		id := uuid.New()
		fakeChannel <- fmt.Sprintf("FAKE %d: - %v", i+1, id.String())
	}
	close(fakeChannel) // end channel
	closeChannel <- 1
	wait.Done()
}

func generateIDS(channel chan<- string, wait *sync.WaitGroup, closeChannel chan<- int) {
	for i := 0; i < 100; i++ {
		id := uuid.New()
		channel <- fmt.Sprintf("%d: - %v", i+1, id.String())
	}
	/*
		channel receive the information
		id.String() send the information to the channel
	*/
	close(channel) // end channel
	closeChannel <- 1
	wait.Done()
}

func logIDs(channel <-chan string, fakeChannel <-chan string, wait *sync.WaitGroup, closeChannel chan int) {
	closedChannels := 0
	for {
		select {
		case id, ok := <-channel:
			if ok {
				fmt.Println("ID", id)
			}

		case id, ok := <-fakeChannel:
			if ok {
				fmt.Println("Fake ID", id)
			}

		case _, ok := <-closeChannel:
			if ok {
				closedChannels++
			}
		}
		if closedChannels == 2 {
			close(closeChannel)
			break
		}
	}
	wait.Done()
	fmt.Println("END")

}
