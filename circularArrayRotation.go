package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here

	//create a variable to store the numbers when first rotated
	//create another variable that will return the final result
	//loop through the slice of array
	//push the last element forward
	//check for k to know the rotation count
	//at queries, push the indices at q forward and the other behind
	//input [3,4,5] will frst rotate [5,3,4]
	//output [4,5,3]

	var firstRotate []int32
	var result []int32

	for i := 0; i < len(a); i++ {
		firstRotate = append(firstRotate, a[i-1])
		fmt.Println(firstRotate)
	}
	return result
}
