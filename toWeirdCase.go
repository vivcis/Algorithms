package main

import "strings"

func main() {

}

//Write a function toWeirdCase (weirdcase in Ruby) that accepts a string, and returns the same string with all even
//indexed characters in each word upper cased, and all odd indexed characters in each word lower cased.
//The indexing just explained is zero based, so the zero-ith index is even, therefore that character should be upper-
//cased and you need to start over for each word. The passed in string will only consist of alphabetical characters
//and spaces(' '). Spaces will only be present if there are multiple words. Words will be separated by a single space(' ').
//E.G toWeirdCase("String") // => returns "StRiNg"
//toWeirdCase("Weird string case") // => returns "WeIrD StRiNg CaSe"

func toWeirdCase(str string) string {
	// Your code here and happy coding!

	if str == "" {
		return ""
	}
	result := ""
	index := 0
	for _, value := range str {
		v := string(value)
		if v == " " {
			result += v
			index = 0
			continue
		}
		if index%2 == 0 {
			result += strings.ToUpper(v)
		} else {
			result += strings.ToLower(v)
		}
		index++
	}
	return result
}
