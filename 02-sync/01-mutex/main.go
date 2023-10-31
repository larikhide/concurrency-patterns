package main

import (
	"fmt"
	"sync"
)

type mutexType chan struct{}
type mutex struct {
	s mutexType
}

func NewMutex() mutex {
	return mutex{s: make(chan struct{}, 1)}
}

func (m *mutex) Lock() {
	e := struct{}{}
	m.s <- e
}

func (m *mutex) Unlock() {
	<-m.s
}

const N = 10000

func main() {
	mu := NewMutex()
	wg := sync.WaitGroup{}
	count := 0
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock() // better if we want do non atomic operation and handle panic
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
