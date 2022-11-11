package main

//Given a string s consisting of words and spaces, return the length of the last word in the string.
//A word is a maximal substring
// consisting of non-space characters only.
//Example 1:
//Input: s = "Hello World"
//Output: 5
//Explanation: The last word is "World" with length 5.
//Example 2:
//Input: s = "   fly me   to   the moon  "
//Output: 4
//Explanation: The last word is "moon" with length 4.
//Example 3:
//Input: s = "luffy is still joyboy"
//Output: 6
//Explanation: The last word is "joyboy" with length 6.
//Constraints:
//1 <= s.length <= 104
//s consists of only English letters and spaces ' '.
//There will be at least one word in s.
func lengthOfLastWord(s string) int {
	var count int
	var prev rune
	for _, r := range s {
		if r == ' ' {
			if prev != ' ' {
				count = 0
			}
		} else {
			count++
		}
		prev = r
	}
	return count
}

//func lengthOfLastWord(s string) int {
//	//create a count variable
//	//iterate through s from the back
//	//my count increments as the loop happens
//	//once it encounters an empty space it should stop
//	//then return my count
//	var length int
//	var lastWord string
//	var arr = strings.Split(s, "")
//	for i := 0; i < len(arr); i++ {
//		if len(arr[i]) > 0 {
//			lastWord = arr[i]
//		}
//	}
//	length = len(lastWord)
//	return length
//}
