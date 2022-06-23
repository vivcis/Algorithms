package main

import (
	"fmt"
	"strconv"
	"strings"
)

//There is an array with some numbers. All numbers are equal except for one. Try to find it!
//E.G findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
//findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
func FindUniq(arr []float32) float32 {
	//create variables and assign them float32 cause automatically it would take float64
	var duplicateNum float32
	var result float32

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				duplicateNum = arr[i]
				break
			}
		}
	}
	for _, num := range arr {
		if num != duplicateNum {
			result = num
			break
		}
	}
	return result
}

//Solution 2
func FindUniq1(arr []float32) float32 {
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for i, v := range arr[1:] {
		if v != arr[i] {
			return v
		}
	}
	return 0
}

//Time conversion on Hackerrank
func timeConversion(s string) string {
	hourStr := string(s[0]) + string(s[1])
	minuteStr := string(s[3]) + string(s[4])
	secondsStr := string(s[6]) + string(s[7])
	hour, _ := strconv.Atoi(hourStr)
	if strings.HasSuffix(s, "AM") && hour == 12 {
		hour = 0
	}
	if strings.HasSuffix(s, "PM") && hour != 12 {
		hour += 12
	}
	return fmt.Sprintf("%02d:%s:%s", hour, minuteStr, secondsStr)
}
