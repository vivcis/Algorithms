package main

import (
	"regexp"
	"unicode"
)

//Not very secure on codewars
//In this example you have to validate if a user input string is alphanumeric.
//The given string is not nil/null/NULL/None, so you don't have to check that.
//The string has the following conditions to be alphanumeric:
//At least one character ("" is not valid)
//Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
//No whitespaces / underscore
func alphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, v := range str {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			continue
		} else {
			return false
		}
	}
	return true
}

func alphanumeric1(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return r.MatchString(s)
}
