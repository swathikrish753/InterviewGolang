package main

import "fmt"

func bubbleSort(arr []int, descending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			shouldSwap := false
			if descending {
				shouldSwap = arr[j] < arr[j+1] // bigger first
			} else {
				shouldSwap = arr[j] > arr[j+1] // smaller first
			}
			if shouldSwap {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func arrangeOddEven(arr []int) []int {
	odds := []int{}
	evens := []int{}

	// separate
	for _, v := range arr {
		if v%2 != 0 {
			odds = append(odds, v)
		} else {
			evens = append(evens, v)
		}
	}

	bubbleSort(odds, false) // ascending
	bubbleSort(evens, true) // descending

	// merge
	result := []int{}
	result = append(result, odds...)
	result = append(result, evens...)
	return result
}

func main() {
	arr := []int{2, 3, 5, 6, 7, 8}
	fmt.Println("Input: ", arr)
	fmt.Println("Output:", arrangeOddEven(arr))
}
