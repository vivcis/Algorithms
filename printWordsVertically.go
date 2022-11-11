package main

import "strings"

//leetcode - 1324. Print Words Vertically
//Given a string s. Return all the words vertically in the same order in which they appear in s.
//Words are returned as a list of strings, complete with spaces when is necessary. (Trailing spaces are not allowed).
//Each word would be put on only one column and that in one column there will be only one word.
//Example 1:
//Input: s = "HOW ARE YOU"
//Output: ["HAY","ORO","WEU"]
//Explanation: Each word is printed vertically.
// "HAY"
// "ORO"
// "WEU"
func printVertically(s string) []string {
	words := strings.Split(s, " ")
	max := 0
	for _, word := range words {
		if len(word) > max {
			max = len(word)
		}
	}
	res := make([]string, max)
	for i := 0; i < max; i++ {
		for _, word := range words {
			if len(word) > i {
				res[i] += string(word[i])
			} else {
				res[i] += " "
			}
		}
	}
	for i := 0; i < len(res); i++ {
		res[i] = strings.TrimRight(res[i], " ")
	}
	return res
}
