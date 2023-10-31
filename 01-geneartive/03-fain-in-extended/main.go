package main

import (
	"fmt"
	"sync"
	"time"
)

type payload struct {
	name  string
	value int
}

func producer(pname string, done <-chan struct{}, wg *sync.WaitGroup) <-chan payload {
	ch := make(chan payload)
	i := 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println(pname, "completed")
				return
			case ch <- payload{
				name:  pname,
				value: i,
			}:
				fmt.Println(pname, "produced", i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	return ch
}

func consumer(cname string, channels []<-chan payload, done <-chan struct{}, wg *sync.WaitGroup, fanIn chan<- payload) {
	for i, ch := range channels {
		i := i + 1
		ch := ch
		go func() {
			defer wg.Done()
			fmt.Println("started consumer", i)
			for {
				select {
				case <-done:
					fmt.Println("consumer ", cname, "stopped")
					return
				case v := <-ch:
					fmt.Println("consumer", cname, i, "got value", v.value, "from", v.name)
					fanIn <- v
				}
			}
		}()
	}
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(3)
	producers := make([]<-chan payload, 0, 3)
	producers = append(producers, producer("Alice", done, &wg))
	producers = append(producers, producer("Bob", done, &wg))
	producers = append(producers, producer("Jack", done, &wg))

	fanIn1 := make(chan payload, 0)
	fanIn2 := make(chan payload, 0)

	wg.Add(3) //consumer start a number of goroutines equal to the number of channels in producers
	consumer("C1", producers, done, &wg, fanIn1)

	wg.Add(3)
	consumer("C2", producers, done, &wg, fanIn2)

	go func() {
		f1Done, f2Done := false, false
		for {
			select {
			case <-done:
				if f1Done && f2Done {
					return
				}
			case v, ok := <-fanIn1:
				if !ok {
					f1Done = true
					continue
				}
				fmt.Printf("fanIn1 got value %v\n", v)
			case v, ok := <-fanIn2:
				if !ok {
					f2Done = true
					continue
				}
				fmt.Printf("fanIn2 got value %v\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}
