package main

import "fmt"

func main() {
	str := "liril"
	fmt.Printf("Original string: %s\n", str)
	// Reverse the string
	reversed := reverseString(str)
	fmt.Printf(" String is palindrome: %t\n", reversed)
}

func reverseString(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}
