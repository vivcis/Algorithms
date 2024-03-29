package main

import "strings"

//HackerRank in a String!
func hackerrankInString(s string) string {
	// Write your code here
	//create an empty string to push each letter found
	//create a count to count for each letter found

	var count = 0
	strTest := ""
	str := "hackerrank"
	strArr := strings.Split(str, "")

	for _, v := range s {

		if string(v) == strArr[count] {
			strTest += string(v)
			count++
		}
		if count == 10 {
			return "YES"
		}

	}
	return "NO"
}

func hackerrankInString1(s string) string {
	// create a variable to store hackerrank
	result := "hackerrank"

	if len(s) < len(result) {
		return "NO"
	}
	for i := 0; i < len(s); i++ {
		if s[i] == result[0] {
			if len(result) == 1 {
				return "YES"
			} else {
				result = result[1:]
			}
		}
	}
	return "NO"
}
