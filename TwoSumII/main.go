package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {
	tests := []struct {
		numbers []int
		target  int
	}{
		{[]int{2, 7, 11, 15}, 9},  // [1, 2]
		{[]int{2, 3, 4}, 6},       // [1, 3]
		{[]int{-1, 0}, -1},        // [1, 2]
		{[]int{1, 2, 3, 4, 5}, 8}, // [3, 5]
	}

	for _, t := range tests {
		fmt.Printf("twoSum(%v, %d) = %v\n", t.numbers, t.target, twoSum(t.numbers, t.target))
	}
}
