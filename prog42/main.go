package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 45, 67, 91}
	target := 23

	fmt.Println("Array:", arr)
	fmt.Printf("Searching for: %d\n", target)

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found in the array\n", target)
	}
}
