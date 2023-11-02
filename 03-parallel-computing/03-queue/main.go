package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 3
const MESSAGES = 10

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Queue lenght: %d\n", N)
	queue := make(chan struct{}, N)

	wg.Add(MESSAGES)

	for work := 0; work < MESSAGES; work++ {
		process(work, queue, &wg)
	}
	wg.Wait()
	// close the channel for the purpose of simplifying the example
	// and using additional synchronization means.
	close(queue)
	fmt.Println("Processing completed")
}

func process(payload int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		fmt.Printf("Start processing of %d\n", payload)
		time.Sleep(200 * time.Millisecond) // do work
		fmt.Printf("Complete processing of %d\n", payload)

		<-queue
	}()
}
