package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int) // value → last seen index

	for i, v := range nums {
		if j, ok := seen[v]; ok && i-j <= k {
			return true
		}
		seen[v] = i // always update to the latest index
	}

	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       // true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) // false
}
