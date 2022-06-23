package main

import (
	"strconv"
	"strings"
)

func main() {

}

//maximum element on hackerrank
func getMax1(operations []string) []int32 {
	stack := make([]int32, 0)
	result := make([]int32, 0)
	for _, v := range operations {
		innerVal := strings.Split(v, " ")
		if innerVal[0] == "1" {
			intVal, _ := strconv.Atoi(innerVal[1])
			stack = append(stack, int32(intVal))
		} else if innerVal[0] == "2" {

			// remove the top element
			stack = stack[:len(stack)-1]
		} else if innerVal[0] == "3" {
			// find the maximum
			var temp int32
			for i := 0; i < len(stack); i++ {
				if stack[i] > temp {
					temp = stack[i]
				}
			}
			// append the max to the result
			result = append(result, int32(temp))
		}
	}
	return result
}

//OR

func getMax(operations []string) []int32 {

	stack := make([]int32, 0)
	result := make([]int32, 0)
	for _, v := range operations {
		innerVal := strings.Split(v, " ")
		if innerVal[0] == "1" {
			intVal, _ := strconv.Atoi(innerVal[1])
			stack = append(stack, int32(intVal))
		} else if innerVal[0] == "2" {
			// remove the top element
			stack = stack[:len(stack)-1]
		} else if innerVal[0] == "3" {
			// find the maximum
			var temp int32
			for i := 0; i < len(stack); i++ {
				if stack[i] > temp {
					temp = stack[i]
				}
			}
			// append the max to the result
			result = append(result, int32(temp))
		}
	}
	return result
}
