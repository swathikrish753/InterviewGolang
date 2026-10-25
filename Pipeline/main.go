package main

import "fmt"

//gen(2,3)
//sq(nums)
//print sq

func main() {
	genCh := gen(2, 3, 5, 7, 11)
	sqCh := sq(genCh)
	for n := range sqCh {
		fmt.Println(n)
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(input chan int) chan int {
	sqCh := make(chan int)
	go func() {
		for n := range input {
			sqCh <- n * n
		}
		close(sqCh)
	}()
	return sqCh
}
