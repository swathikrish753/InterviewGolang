package main

import "fmt"

func isAlnum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlnum(s[left]) {
			left++
		}
		for left < right && !isAlnum(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	tests := []string{
		"A man, a plan, a canal: Panama", // true
		"race a car",                     // false
		"",                               // true
		"racecar",                        // true
	}

	for _, t := range tests {
		fmt.Printf("isPalindrome(%q) = %v\n", t, isPalindrome(t))
	}
}
