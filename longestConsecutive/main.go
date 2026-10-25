package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	longest := 0
	for n := range set {
		// only start counting from the beginning of a sequence
		if set[n-1] {
			continue
		}

		length := 1
		for set[n+length] {
			length++
		}
		if length > longest {
			longest = length
		}
	}
	return longest
}

func main() {
	testCases := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		{},
	}

	for _, tc := range testCases {
		result := longestConsecutive(tc)
		fmt.Printf("longestConsecutive(%v) = %v\n", tc, result)
	}
}
