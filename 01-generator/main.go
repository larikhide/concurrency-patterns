package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{}) // empty struct is lighter than a context for graceful shutdown
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := makeGenerator(done, &wg)

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Printf("Value %v\n", v)
		}
	}()

	time.Sleep(time.Millisecond * 500)
	close(done)
	wg.Wait()
}

func makeGenerator(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int, 5)
	i := 0

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				fmt.Println("done")
				return
			default:
				time.Sleep(time.Millisecond * 100) //cpu-bound imitation
				ch <- i                            // read from queue imitation
				i++
			}
		}
	}()
	return ch
}
