package main

import (
	"math"
)
/*
	Name: Reverse Integer
	Problem: https://leetcode.com/problems/reverse-integer
*/
func main() {
	var number = -1534236469

	println(reverse(number))
}

func reverse(x int) int {
	isMinus := 1
	if x < 0 {
		isMinus = -1
		x = x * -1
	}

	number := 0
	for x > 0 {
		tmp := x % 10
		number = number * 10 + tmp

		x = x / 10
	}

	number = number * isMinus
	
	if number > math.MaxInt32 || number < math.MinInt32 {
		return 0
	}

	return number
}