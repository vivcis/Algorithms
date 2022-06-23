package main

import (
	"strings"
	"unicode"
)

func main() {

}

//Implement a function that determines whether a string is an isogram. Assume an empty string is an isogram, and ignore letter case.
func IsIsogram(s string) bool {
	m := make(map[rune]bool)
	for _, r := range s {
		if m[unicode.ToLower(r)] {
			return false
		}
		m[unicode.ToLower(r)] = true
	}
	return true
}

//isIsogram return true
//is not should return false e.g aba
//Dermatoglyphic
func IsIsogram1(str string) bool {
	if str == "" {
		return true
	}
	word := strings.ToLower(str)
	for _, v := range word {
		if strings.ContainsAny(word, string(v)) {
			return false
		}

	}
	return true
}
