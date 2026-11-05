package main

import "fmt"

// minSubArrayLen returns the length of the shortest contiguous subarray
// with sum >= target, or 0 if none exists. Assumes positive numbers.
// Variable sliding window: O(n) time, O(1) space.
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	best := len(nums) + 1 // sentinel: "bigger than any real window"

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // right hand grabs the next item

		// While the window already meets the goal, shrink it from the left
		// to find the tightest qualifying window.
		for sum >= target {
			if size := right - left + 1; size < best {
				best = size
			}
			sum -= nums[left] // drop the leftmost item
			left++
		}
	}

	if best == len(nums)+1 {
		return 0 // never reached target
	}
	return best
}

func main() {
	tests := []struct {
		target int
		nums   []int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}},  // 2  -> [4,3]
		{4, []int{1, 4, 4}},           // 1  -> [4]
		{11, []int{1, 1, 1, 1, 1, 1}}, // 0  -> impossible
		{15, []int{1, 2, 3, 4, 5}},    // 5  -> whole array
	}

	for _, tc := range tests {
		fmt.Printf("minSubArrayLen(%d, %v) = %d\n", tc.target, tc.nums, minSubArrayLen(tc.target, tc.nums))
	}
}
