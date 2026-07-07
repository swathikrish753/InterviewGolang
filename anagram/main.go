package main

import "fmt"

func isAnagramGeneral(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freq := make(map[rune]int)

	for _, ch := range s1 {
		freq[ch]++
	}
	for _, ch := range s2 {
		freq[ch]--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		s1, s2 string
	}{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"Ño", "oÑ"},
		{"Listen", "Silent"},
		{"", ""},
		{"a", "ab"},
	}

	for _, tc := range testCases {
		result := isAnagramGeneral(tc.s1, tc.s2)
		fmt.Printf("isAnagramGeneral(%q, %q) = %v\n", tc.s1, tc.s2, result)
	}
}
