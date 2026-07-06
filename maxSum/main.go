package main

import "fmt"

func maxSum(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}

	// first window sum
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	maxSum := sum

	// slide window
	for i := k; i < len(arr); i++ {
		sum = sum + arr[i] - arr[i-k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	arr := []int{1, 3, 5, 7, -3}
	k := 3

	fmt.Println(maxSum(arr, k))
}
