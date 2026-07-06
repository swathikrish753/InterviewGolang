package main

import "fmt"

func main() {
	numbers := []int{22, 1, 56, 7, 3, 9, 12}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Printf("the sorted numbers are %v \n", numbers)
	println("the second largest number is", numbers[len(numbers)-2])
}
