package main

import "fmt"

// The number of workers can be greater than the logical cores of the processor,
// as worker pools are usually used for I/O-bound tasks such as network calls.
// The specific value should be set depending on the input data flow in a way
// that workers don't idle in case of low data volume and
// manage to process in case of high data volume.
const totalJobs = 10
const totalWorkers = 5

func main() {
	// channels is buffered to prevent blocking during write and read operations.
	jobs, results := make(chan int, totalJobs), make(chan int, totalWorkers)

	for workerId := 0; workerId <= totalWorkers; workerId++ {
		go worker(workerId, jobs, results)

	}

	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}

	// close the channel for the purpose of simplifying the example
	// and using additional synchronization means.
	close(jobs)

	for a := 1; a <= totalJobs; a++ {
		<-results
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		go func(job int) {
			fmt.Printf("Worker %d started job %d\n", id, job)

			result := job * 2 //do some work
			results <- result

			fmt.Printf("Worker %d finished job %d\n", id, job)
		}(j)
	}
}
