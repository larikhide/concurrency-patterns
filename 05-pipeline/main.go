package main

import (
	"fmt"
	"math"
)

func main() {
	in := generateWork([]int{1, 2, 3, 4, 5, 6, 7, 8})

	out := filterOdd(in) //imitate 1st step
	out = square(out)    //2nd
	out = half(out)      //3rd

	for v := range out {
		fmt.Println(v)
	}
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, w := range work {
			ch <- w
		}
	}()
	return ch
}

func filterOdd(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			v := math.Pow(float64(i), 2)
			out <- int(v)
		}
	}()
	return out
}

func half(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			value := i / 2
			out <- value
			// i = i / 2
			// out <- i
		}
	}()
	return out
}
