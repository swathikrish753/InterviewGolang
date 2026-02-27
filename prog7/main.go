package main

import "fmt"

func main() {
	arr := []int{0, 2, 3, 4, 1}
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	for i, val := range arr {
		if i != val {
			println(i)
			break
		}
		if val == length-1 {
			println(val + 1)
			return
		}
	}
}
