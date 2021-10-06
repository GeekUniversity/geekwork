package main

import "fmt"

// Leetcode66: https://leetcode-cn.com/problems/plus-one/
func plusOne(digits []int) []int {
	length := len(digits)
	increment := 1
	for i := length - 1; i >= 0; i-- {
		digits[i] += increment
		if digits[i] == 10 {
			digits[i] = 0
			continue
		}
		increment = 0
		break
	}
	// 如果总位数不够需要进一
	if increment == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits

}

func main() {
	input := []int{9}
	res := plusOne(input)
	fmt.Println(res)
}
