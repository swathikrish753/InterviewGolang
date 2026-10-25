package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "héllo, wörld"
	fmt.Println("len (bytes):", len(s))                   // 14
	fmt.Println("rune count:", utf8.RuneCountInString(s)) // 12
	fmt.Printf("s[1] = %d (%T) -- a byte, not 'é'\n", s[1], s[1])
	fmt.Print("range gives runes: ")
	for i, r := range s {
		fmt.Printf("[%d:%c] ", i, r) // note the byte offsets jump
	}
	fmt.Println()
	// Correct conversions
	runes := []rune(s)
	bytes := []byte(s)
	fmt.Println("runes[1]:", runes[1], "len(runes):", len(runes))
	fmt.Println("len(bytes):", len(bytes))
	// Reversing a string correctly (rune-aware)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("reversed:", string(runes))
}
