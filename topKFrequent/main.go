package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// buckets[i] = numbers that occur exactly i times; i ranges 0..len(nums)
	buckets := make([][]int, len(nums)+1)
	for n, count := range freq {
		buckets[count] = append(buckets[count], n)
	}

	result := make([]int, 0, k)
	for count := len(buckets) - 1; count >= 0 && len(result) < k; count-- {
		for _, n := range buckets[count] {
			result = append(result, n)
			if len(result) == k {
				break
			}
		}
	}
	return result
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2},
		{[]int{1}, 1},
		{[]int{4, 1, -1, 2, -1, 2, 3}, 2},
	}

	for _, tc := range testCases {
		result := topKFrequent(tc.nums, tc.k)
		fmt.Printf("topKFrequent(%v, %d) = %v\n", tc.nums, tc.k, result)
	}
}
