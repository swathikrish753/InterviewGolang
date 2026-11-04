package main

import "fmt"

// lengthOfLongestSubstring returns the length of the longest substring
// with all-unique characters. Variable sliding window: O(n) time.
func lengthOfLongestSubstring(s string) int {
	// letters currently inside the window
	window := make(map[byte]struct{})
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		// If s[right] is already inside, shrink from the left
		// until the duplicate is removed.
		for {
			if _, dup := window[s[right]]; !dup {
				break
			}
			delete(window, s[left]) // drop the leftmost letter
			left++
		}

		window[s[right]] = struct{}{} // now safe to add the new letter

		if size := right - left + 1; size > best {
			best = size // record the widest window so far
		}
	}

	return best
}

func main() {
	tests := []string{
		"abcabcbb", // 3
		"bbbbb",    // 1
		"pwwkew",   // 3
		"",         // 0
		"au",       // 2
		"dvdf",     // 3
	}

	for _, s := range tests {
		fmt.Printf("lengthOfLongestSubstring(%q) = %d\n", s, lengthOfLongestSubstring(s))
	}
}
