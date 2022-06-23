package main

import "fmt"

//Apple and Oranges
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var newApple []int32
	var newOranges []int32
	var orangeCount int32
	var appleCount int32

	for _, v := range apples {
		newApple = append(newApple, v+a)
	}
	for _, w := range oranges {
		newOranges = append(newOranges, w+b)
	}
	for _, x := range newApple {
		if x >= s && x <= t {
			appleCount++
		}
	}
	for _, x := range newOranges {
		if x >= s && x <= t {
			orangeCount++
		}
	}
	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}
