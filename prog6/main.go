package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{2, 7, 11, 13, 5}
	target := 12
	ind1, ind2, err := TwoSum(arr, target)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("the indices are %d and %d ", ind1, ind2)

}

func TwoSum(arr []int, target int) (i int, j int, err error) {
	hashM := make(map[int]int)
	for i, val := range arr {
		hashM[val] = i
	}
	for i, val := range arr {
		diff := target - val
		if j, ok := hashM[diff]; ok && i != j {
			return i, j, nil
		}
	}
	return 0, 0, errors.New("the indices not found for the sum")
}
