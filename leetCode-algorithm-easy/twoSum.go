package main

import (
	"fmt"
)
/*
	My LeetCode first problem :)
	Name: Two Sum
	Problem: https://leetcode.com/problems/two-sum
 */
func main() {
	nums := []int{9, 75, 46, 98, 13, 1}
	fmt.Println(twoSum(nums, 10))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var compliment = target - nums[i]
		if v, ok := m[compliment]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	fmt.Println("No two sum solution")

	return []int{0, 0}
}
