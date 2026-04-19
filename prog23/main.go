package main

import "fmt"

func main() {
	haystack := "sanfrancisco"
	needle := "san"
	isPresent := isSubstringPresent(haystack, needle)
	fmt.Printf("the substring presence bool value is %v", isPresent)

}
func isSubstringPresent(mainStr string, subStr string) bool {
	n := len(subStr)
	for i := 0; i < (len(mainStr) - n); i++ {
		fmt.Println(mainStr[i : i+n])
		if mainStr[i:i+n] == subStr {
			return true
		}
	}
	return false
}
