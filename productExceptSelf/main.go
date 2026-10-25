package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// result[i] starts as the product of everything to the left of i
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// multiply in the product of everything to the right of i
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func main() {
	testCases := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
		{2, 3},
	}

	for _, tc := range testCases {
		result := productExceptSelf(tc)
		fmt.Printf("productExceptSelf(%v) = %v\n", tc, result)
	}
}
