package main

import (
	"math"
)
/*
	Name: Longest Substring Without Repeating Characters
	Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters
*/
func main() {
	str := "pwwkwwewwewkew"

	println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	i := 0
	max := 0
	for j := 0; j < len(s); j++ {
		if v, ok := m[s[j]]; ok {
			i = int(math.Max(float64(v), float64(i)))
		}
		max = int(math.Max(float64(max), float64(j - i + 1)))
		m[s[j]] = j + 1
	}

	return max
}
