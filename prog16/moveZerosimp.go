package main

import "fmt"

func shiftZerosRight(nums []int) []int {
	nonZero := []int{}
	zeroCount := 0

	for _, v := range nums {
		if v != 0 {
			nonZero = append(nonZero, v) // push only non-zero
		} else {
			zeroCount++
		}
	}

	// append zeros at the end
	for i := 0; i < zeroCount; i++ {
		nonZero = append(nonZero, 0)
	}

	return nonZero
}

func main() {
	arr := []int{3, 0, 4, 0, 5, 0, 6}
	result := shiftZerosRight(arr)
	fmt.Println(result)
}
