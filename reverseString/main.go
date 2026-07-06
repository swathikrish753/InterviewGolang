package main

import "strings"

func main() {
	s := "Hello, World! Welcome to Go programming."
	words := strings.Fields(s)
	low := 0
	high := len(words) - 1
	for low < high {
		words[low], words[high] = words[high], words[low]
		low++
		high--
	}
	result := strings.Join(words, " ")
	println(result)
}
