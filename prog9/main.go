package main

import "fmt"

func main() {
	s := "Hello, World! Welcome to Go programming."
	chars := []rune(s)
	length := len(chars)
	for i := 0; i < length/2; i++ {
		chars[i], chars[length-1-i] = chars[length-1-i], chars[i]
	}
	fmt.Println(string(chars))
}
