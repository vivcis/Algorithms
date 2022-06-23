package main

import (
	"fmt"
	"strconv"
)

//staircase on hackerrank
func staircase(n int32) {
	// Write your code here
	str := ""
	for i := int32(0); i < n; i++ {
		str += "#"

		// Use %v for printing with padding
		// Example: printing with 10 space padding to right: Printf("%10v", str)

		fmt.Printf("%"+strconv.FormatInt(int64(n), 10)+"v\n", str)
	}

}
