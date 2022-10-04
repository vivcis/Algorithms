package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
	"unicode"
)

func main() {
	//fmt.Println(Numerals(900))
	fmt.Println(missingNumbers([]int{1, 4, 5, 6, 7, 8, 9, 0, 3, 3}))
	//fmt.Println(PrintLowerCase("I am GOING Home"))
	fmt.Println(lowercase("I am GOING Home"))
	fmt.Println(countInRange([]int{1, 3, 5, 7, 8, 9, 0}, 0, 9))
}

//input => [2,3,4,5], 9
//output => [2,3]
func twoSum(nums []int, target int) []int {
	result := []int{}
	var sum int

	//iterate through the numbers
	//find indexes of the array that adds up to get the target
	//if its true return i, j
	//else continue and return nil

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]

			if sum == target {
				result = append(result, i, j)
			} else {
				continue
			}
		}
	}
	return result
}

func extraLongFactorials1(n int32) {
	// Write your code here
	fac := big.NewInt(1)
	for ; n >= 1; n-- {
		fac = fac.Mul(fac, big.NewInt(int64(n)))
	}
	fmt.Println(fac)
}

//Extra Long Factorials on hackerrank
func extraLongFactorials(n int32) {
	//Declare a variable to hold the factorial
	var fact = big.NewInt(1)
	//Iterate from 1 to the given number
	for i := int32(1); i <= n; i++ {
		//Multiply the current factorial by the current number
		fact = fact.Mul(fact, big.NewInt(int64(i)))
	}
	//Print the factorial
	fmt.Println(fact)
	//extraLongFactorials(25) => 15511210043330985984000000
}

func Factorials(n int32) {
	var fact int32 = 1
	//Iterate from 1 to the given number
	for i := int32(1); i <= n; i++ {
		//Multiply the current factorial by the current number
		fact *= i
	}
	fmt.Println(fact)
}

//1990 => MCMXC
func Numerals(num int) string {
	//create a var to hold the result or append to
	//divide the given num by 1000, 500, 100 and so on till there's no remainder

	//I := 1
	//V := 5
	//X := 10
	//L := 50
	XC := 90
	//C := 100
	//D := 500
	CM := 900
	M := 1000
	var result []string
	for i := 0; i < num; i++ {
		if num%1000 == 900 {
			result = append(result, string(M))
		}
		if num%500 == 90 {
			result = append(result, string(CM))
		}
		if num%90 == 0 {
			result = append(result, string(XC))
		}
		//if strings.HasPrefix(string(num), string(2)) {
		//
		//}
	}
	x := strings.Join(result, "")
	return x
}

//Implement a function that returns the missing number in a given integer array between 0 and 9.
//If there is no missing number, the function should return -1.

func missingNumbers(a []int) int {
	var no = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		if a[i] != no[i] {
			return a[i] - 1
		}
	}
	return -1
}

//write a function that returns the missing number in a given integer array between 0 and 9.
//if there is no missing number, the function should return -1.
//func missingNumber (arr []int) int {
//	// loop through arr
//	for i:=0; i<len(arr); i++ {
//		// if arr[i] is between 0 and 9
//		if arr[i] >= 0 && arr[i] <= 9 {
//			// if arr[i] is not in result
//			if !contains(result, arr[i]) {
//				// append arr[i] to result
//				result = append(result, arr[i])
//			}
//		}
//	}
//}
func missing(a []int) int {
	m := map[string]int{"1": 0, "2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "10": 9}

	for _, v := range a {
		if _, ok := m[string(v)]; ok {
			fmt.Println(-1)
		}
	}
	return 0
}

//Return all the lowercase letters in any given string.

func PrintLowerCase(a string) string {
	var str string
	for _, v := range a {
		if unicode.IsLower(v) {
			str = string(v)
		}
	}
	return str
}

//write a function that returns all the lowercase in any given string
func lowercase(str string) []string {
	var result []string
	// loop through str
	for i := 0; i < len(str); i++ {
		// if str[i] is lowercase
		if str[i] >= 'a' && str[i] <= 'z' {
			// append str[i] to result
			result = append(result, string(str[i]))
		}
	}
	return result
}

//write a function that returns the number of elements in an array of numbers that are within the range specified
//of a number lower range and another number upper range, inclusive or exclusive
func countInRange(arr []int, lower, upper int) int {
	// loop through arr
	var count int
	for i := 0; i < len(arr); i++ {
		// if arr[i] is between lower and upper
		if arr[i] >= lower && arr[i] <= upper {
			// increment count
			count++
		}
	}
	return count
}
