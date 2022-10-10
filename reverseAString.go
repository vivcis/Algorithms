package main

import (
	"fmt"
	"strings"
)

func main() {
	reverse1("meal")
	// Reversing the string.
	str := "the sky is blue"

	// returns the reversed string.
	strRev := reverseWords(str)
	fmt.Println(str)
	fmt.Println(strRev)
	//reverseWords("the sky is blue")
}

func reverse1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//Given a string s, reverse the order of characters in each word within a sentence while
//still preserving whitespace and initial word order.
//Example 1:
//Input: s = "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
func reverseWords2(s string) (result string) {
	sArr := strings.Split(s, " ")
	for i := 0; i < len(sArr); i++ {
		sArr[i] = reverseaStr(sArr[i], len(sArr[i]))
	}
	return strings.Join(sArr, " ")
}

func reverseaStr1(s string, k int) string {
	s = s[:k]
	var sr []rune
	for i := len(s) - 1; i >= 0; i-- {
		sr = append(sr, rune(s[i]))
	}
	return string(sr)
}

func reverseaStr(s string, k int) string {
	s = s[:k]
	var sr []rune
	for i := len(s) - 1; i >= 0; i-- {
		sr = append(sr, rune(s[i]))
	}
	return string(sr)
}

//541. Reverse String II(Leetcode)
//Given a string s and an integer k, reverse the first k characters for every 2k characters counting
//from the start of the string.
//If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater
//than or equal to k characters, then reverse the first k characters and leave the other as original.
//Example 1:
//Input: s = "abcdefg", k = 2
//Output: "bacdfeg"
func reverseStr(s string, k int) (result string) {
	//loop through the string given
	//if input is less than k, reverse all
	//if input is greater than k, reverse first k characters n leave the rest

	if len(s) >= (k * 2) {
		revStr := strings.Split(s[:(k*2)], "")
		remStr := strings.Split(s[k*2:], "")
		revStr[0], revStr[1] = revStr[1], revStr[0]

		if len(remStr) >= (k*2)-1 {
			remStr[0], remStr[1] = remStr[1], remStr[0]

		}
		return strings.Join(revStr, "") + strings.Join(remStr, "")
	}
	return s
}

func reverseString(s []byte) {
	var left = 0
	var right = len(s) - 1
	if left < right {
		var tempStart = s[left]
		var tempEnd = s[right]
		s[left] = tempEnd
		s[right] = tempStart
		left++
		right--
	}
}

func reverseString1(s []byte) {
	//Write a function that reverses a string. The input string is given as an array of characters s.
	//You must do this by modifying the input array in-place with O(1) extra memory.
	//Example 1:
	//Input: s = ["h","e","l","l","o"]
	//Output: ["o","l","l","e","h"]
	m := s
	for i := len(s) - 1; i > 0; i-- {
		s[(len(s)-1)-i] = m[i]
	}
}

//Given an input string s, reverse the order of the words.
//A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
//Return a string of the words in reverse order concatenated by a single space.
//Input: s = "the sky is blue"
//Output: "blue is sky the"
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var result []string
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) != 0 {
			if len(result) != 0 {
				result = append(result, " ")
			}
			result = append(result, words[i])
		}
	}
	str := strings.Join(result, "")
	return str
}
