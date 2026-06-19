package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[2] // seed with first triplet

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			switch {
			case sum == target:
				return sum // can't get closer than exact
			case sum < target:
				left++
			default:
				right--
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))  // expect 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))       // expect 0
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100)) // expect 2
}
