package main

import (
	"fmt"
	"sync"
)

type Result struct {
	num  int
	kind string
}

func main() {
	oddArr := []int{1, 3, 5, 7}
	evenArr := []int{2, 4, 6, 8}

	oddTurn := make(chan struct{})
	evenTurn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for _, n := range oddArr {
			<-oddTurn
			fmt.Println(Result{n, "odd"})
			evenTurn <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()

		for _, n := range evenArr {
			<-evenTurn
			fmt.Println(Result{n, "even"})

			// Don't signal odd after the last even number
			if n != evenArr[len(evenArr)-1] {
				oddTurn <- struct{}{}
			}
		}
	}()

	// Start with odd
	oddTurn <- struct{}{}

	wg.Wait()
}
