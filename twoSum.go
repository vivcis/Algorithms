package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println(Numerals(900))
}

//input => [2,3,4,5], 9
//output => [2,3]
func twoSum(nums []int, target int) []int {
	result := []int{}
	var sum int

	//iterate through the numbers
	//find indeces of the array that adds up to get the target
	//if its true return i, j
	//else continue and return nill

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
