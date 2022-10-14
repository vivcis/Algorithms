package main

import "strings"

//First non-repeating character on codewars
//Write a function named first_non_repeating_letter that takes a string input,
//and returns the first character that is not repeated anywhere in the string.
//For example, if given the input 'stress', the function should return 't',
//since the letter t only occurs once in the string, and occurs first in the string.
//As an added challenge, upper- and lowercase letters are considered the same character, but the function should
//return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.
//If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.
func FirstNonRepeating(str string) string {
	for _, c := range str {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
			return string(c)
		}
	}
	return ""
}

func FirstNonRepeating1(str string) string {
	//your code here

	result := ""
	for i := 0; i < len(str); i++ {
		found := false

		for j := i - 1; j >= 0; j-- {
			if strings.ToUpper(string(str[j])) == strings.ToUpper(string(str[i])) {
				found = true
				break
			}
		}
		if found {
			continue
		}
		for j := i + 1; j < len(str); j++ {
			if strings.ToUpper(string(str[j])) == strings.ToUpper(string(str[i])) {
				found = true
				break
			}
		}
		if !found {
			result = string(str[i])
			break
		}
	}
	return result
}

func FirstNonRepeating2(str string) string {
	s := strings.ToLower(str)

	for i := 0; i < len(s); i++ {
		st := string(s[i])
		if strings.Index(s, st) == strings.LastIndex(s, st) {
			return string(str[i])
		}
	}
	return ""
}
