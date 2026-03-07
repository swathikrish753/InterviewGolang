package main

import "fmt"

func main() {
	arr := []int{3, 7, 2, 9, 4}
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)
}
