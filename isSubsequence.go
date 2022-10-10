package main

//392. Is Subsequence(Leetcode)
//Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//A subsequence of a string is a new string that is formed from the original string by deleting
//some (can be none) of the characters without disturbing the relative positions of the remaining characters.
//(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//Example 1:
//Input: s = "abc", t = "ahbgdc"
//Output: true
func isSubsequence(s string, t string) bool {
	//loop through the string given
	//another loop to check for similarities
	//check at indexes if the substring is a subset of the s given

	result := ""
	for _, v := range s {
		for j, val := range t {
			if v == val {
				result += string(v)
				t = t[j+1:]
				break
			}
		}
	}
	if result == s {
		return true
	}

	return false
}
