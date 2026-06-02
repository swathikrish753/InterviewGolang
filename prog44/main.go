package main

import (
	"fmt"
	"sort"
)

func sequentialDigits(low int, high int) []int {

	var result []int

	for start := 1; start <= 9; start++ {
		num := start

		for next := start + 1; next <= 9; next++ {
			num = num*10 + next

			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	low := 100
	high := 300

	fmt.Println(sequentialDigits(low, high))
}
