package main

import (
	"math"
	"strconv"
	"strings"
)

//In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.
//codewars - Highest and Lowest

func HighAndLow(in string) string {

	maxNum := -math.MaxInt32
	minNum := math.MaxInt32

	numbers := strings.Fields(in)

	for _, v := range numbers {
		currentNum, _ := strconv.Atoi(v)
		if maxNum < currentNum {
			maxNum = currentNum
		}
		if minNum > currentNum {
			minNum = currentNum
		}

	}
	return strconv.Itoa(maxNum) + " " + strconv.Itoa(minNum)
}
