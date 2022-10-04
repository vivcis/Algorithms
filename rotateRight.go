package main

import (
	"fmt"
	"unicode"
)

func rotateRight(nums []int, k int) []int {
	// Write your code here
	k = k % len(nums)
	nums = append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	return nums

}

func main() {
	//fmt.Println(rotateRight([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	fmt.Println(subsets([]int{0}))
	fmt.Println(isValid("abc123F@"))
}

// [1,2,3] => [[], [1], [2], [1,2], [3], [2,3], [1,2,3]]
func subsets(nums []int) [][]int {
	// Write your code here
	result := [][]int{}
	result = append(result, []int{})
	for _, val := range nums {
		for _, v := range result {
			newV := append(v, val)
			result = append(result, newV)
		}
	}
	return result
}

func IsValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
