package main

import "fmt"

func moveZeroes(nums []int) {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

func main() {
	tests := [][]int{
		{0, 1, 0, 3, 12}, // [1 3 12 0 0]
		{0, 0, 1},        // [1 0 0]
		{0},              // [0]
		{1, 2, 3},        // [1 2 3]
		{0, 0, 0, 4},     // [4 0 0 0]
	}

	for _, t := range tests {
		moveZeroes(t)
		fmt.Println(t)
	}
}
