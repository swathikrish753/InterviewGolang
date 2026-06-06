package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	// build first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	// slide the window
	for i := k; i < len(nums); i++ {
		windowSum += nums[i]   // right comes in
		windowSum -= nums[i-k] // left goes out
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	tests := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4}, // 12.75
		{[]int{5}, 1},                    // 5.0
		{[]int{0, 1, 1, 3, 3}, 4},        // 2.0
		{[]int{-1, -2, -3, -4}, 2},       // -1.5
	}

	for _, t := range tests {
		fmt.Printf("findMaxAverage(%v, %d) = %.2f\n", t.nums, t.k, findMaxAverage(t.nums, t.k))
	}
}
