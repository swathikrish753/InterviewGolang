package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums)-2; i++ {

		// Skip duplicate first elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{
					nums[i],
					nums[left],
					nums[right],
				})

				left++
				right--

				// Skip duplicate left values
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicate right values
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}
