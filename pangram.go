package main

import "strings"

func main() {

}

//A pangram is a string that contains every letter of the alphabet. Given a sentence determine whether it is a pangram in the English alphabet.
//Ignore case. Return either pangram or not pangram as appropriate.
//E.G We promptly judged antique ivory buckles for the next prize => WILL RETURN pangram cause all the English letters are contained
//We promptly judged antique ivory buckles for the prize => not pangram cause The string lacks an x.

func pangrams(s string) string {
	// Write your code here
	ss := strings.Split(strings.ToLower(s), " ")
	s = strings.Join(ss, "")
	if len(s) < 26 {
		return "not pangram"
	}
	m := map[rune]bool{}
	for _, v := range s {
		m[v] = true
	}
	if len(m) == 26 {
		return "pangram"
	}
	return "not pangram"
}
