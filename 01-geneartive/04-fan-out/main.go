package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}                //only a WaitGroup is sufficient, cause the volume of input data is limited
	work := []int{1, 2, 3, 4, 5, 6, 7, 8} // Otherwise, a signaling channel chan struct{} will be needed for graceful shutdown

	wg.Add(1)
	in := generateWork(work, &wg)

	wg.Add(3)
	fanOut("Alice", in, &wg)
	fanOut("Bob", in, &wg)
	fanOut("Jack", in, &wg)

	wg.Wait()
}

func generateWork(work []int, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch) //It will serve as a signal for the reading goroutines in fanOut
		//when ch is closed and all data is read, the for range loop in fanOut will automatically exit
		for _, v := range work {
			ch <- v
		}
		fmt.Println("All data written")
	}()

	return ch
}

func fanOut(name string, in <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		//when in is closed and all data is read, the for range loop will automatically exit
		for v := range in {
			fmt.Println("fanOut", name, "got value", v)
		}
	}()
}
