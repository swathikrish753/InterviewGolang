package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i]%2 == 0 {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}
	fmt.Println(arr)
}
