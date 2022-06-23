package main

import "strings"

func main() {

}

//Write a function that takes a string of braces, and determines if the order of the braces is valid.
//It should return true if the string is valid, and false if it's invalid.
//Examples
//"(){}[]"   =>  True
//"([{}])"   =>  True
//"(}"       =>  False
//"[(])"     =>  False
//"[({})](]" =>  False

func ValidBraces(str string) bool {

	if str == "" {
		return true
	}
	var stack []string
	for _, t := range str {
		s := string(t)
		if s == "(" || s == "{" || s == "[" {
			stack = append(stack, s)
		} else {
			if len(stack) == 0 {
				return false
			}
		}
		lastStr := stack[len(stack)-1]
		if s == ")" {
			if lastStr == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if s == "]" {
			if lastStr == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if s == "}" {
			if lastStr == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

//OR
func ValidBraces1(stringBraces string) bool {

	for strings.Contains(stringBraces, "()") || strings.Contains(stringBraces, "[]") || strings.Contains(stringBraces, "{}") {
		stringBraces = strings.Replace(stringBraces, "()", "", -1)
		stringBraces = strings.Replace(stringBraces, "[]", "", -1)
		stringBraces = strings.Replace(stringBraces, "{}", "", -1)
	}

	if stringBraces == "" {
		return true
	} else {
		return false
	}
}
