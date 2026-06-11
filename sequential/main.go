package main

import (
	"fmt"
	"sync"
)

type Result struct {
	num   int
	label string
}

// producer — sends numbers 1..10 into jobs, then closes
func producer(jobs chan<- int) {
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
}

// consumer — reads from jobs, classifies, sends to results
func consumer(jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs { // exits when jobs is closed
		label := "odd"
		if n%2 == 0 {
			label = "even"
		}
		results <- Result{n, label}
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan Result, 10)

	// start producer
	go producer(jobs)

	// start 2 consumers
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consumer(jobs, results, &wg)
	}

	// close results once both consumers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// collect and sort by number
	output := make([]Result, 10)
	for r := range results {
		output[r.num-1] = r
	}

	for _, r := range output {
		fmt.Printf("%d → %s\n", r.num, r.label)
	}
}
