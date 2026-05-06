package main

import "fmt"

func main() {
	arr := []int{0, 12, 0, 34, 34, 898, 0, 0, 34}
	n := len(arr)
	fmt.Println(result(arr, n))
}

func result(arr []int, n int) []int {
	left := 0

	for right := 0; right < n; right++ {
		if arr[right] != 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
	}
	return arr
}
