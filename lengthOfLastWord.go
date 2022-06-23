package main

import "strings"

func main() {

}

func lengthOfLastWord(s string) int {
	//create a count variable
	//iterate through s from the back
	//my count increments as the loop happens
	//once it encounters an empty space it should stop
	//then return my count
	var length int
	var lastWord string
	var arr = strings.Split(s, "")
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			lastWord = arr[i]
		}
	}
	length = len(lastWord)
	return length
}
