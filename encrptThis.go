package main

import (
	"fmt"
	"strings"
)

//Encrypt this! on codewars
func EncryptThis(text string) string {

	if text == "" {
		return ""
	}
	words := strings.Fields(text)
	result := ""
	for _, v := range words {
		each_word := ""
		firstChar := ""
		secondChar := ""
		lastChar := ""

		for i := 0; i < len(v); i++ {
			if i == 0 {
				//first char
				firstChar = fmt.Sprintf("%d", v[i])
			} else if i == 1 {
				//second char
				secondChar = string(v[i])
			} else if i == len(v)-1 {
				//last char
				lastChar = string(v[i])
			} else {
				each_word += string(v[i])
			}

		}
		each_word = lastChar + each_word + secondChar
		each_word = firstChar + each_word

		result += each_word + " "
	}
	return result[:len(result)-1]
}
