package main

import "strings"

func main() {

}

//Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number
//of characters then it should replace the missing second character of the final pair with an underscore ('_').
//e.g Solution("abc") //should return ["ab", "c_"]
//Solution("abcdef") //should return ["ab", "cd", "ef"]

func Solution(str string) []string {

	//to append result
	var result []string

	//set counter to zero since the task is to put the characters in pairs
	counter := 0

	//a temp string, empty
	tempStr := ""

	//last character too, empty
	lastChar := ' '

	//if str is empty, return an empty slice
	if str == "" {
		return result
	}
	//loop over each of the characters
	for _, v := range str {

		//increment counter
		counter++
		tempStr += string(v)
		lastChar = v

		//conditional statement
		if counter == 2 {

			//append to result
			result = append(result, tempStr)
			//set back the counter to 0 and temporary string empty
			counter = 0
			tempStr = ""
		}
	}
	//if counter is 1 or odd, append to result with an underscore
	if counter == 1 {
		result = append(result, string(lastChar)+"_")
	}
	return result
}

//Strings end with?
func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
