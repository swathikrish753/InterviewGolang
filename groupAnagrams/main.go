package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := sortedKey(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortedKey(s string) string {
	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
	return string(chars)
}

func main() {
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}

	for _, tc := range testCases {
		result := groupAnagrams(tc)
		fmt.Printf("groupAnagrams(%v) = %v\n", tc, result)
	}
}
