package main

import "fmt"

func main() {
	// Empty main function
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for num := range c {
		fmt.Println(num)
	}
}
