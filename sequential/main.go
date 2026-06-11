package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	results := make([]string, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%2 == 0 {
				results[n-1] = "even"
			} else {
				results[n-1] = "odd"
			}
		}(i)
	}

	wg.Wait() // wait for all goroutines to finish

	for i, label := range results {
		fmt.Printf("%d → %s\n", i+1, label)
	}
}
