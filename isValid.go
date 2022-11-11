package main

//leetcode - 20. Valid Parentheses
//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//An input string is valid if:
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Every close bracket has a corresponding open bracket of the same type.
//Example 1:
//Input: s = "()"
//Output: true

func isValid(s string) bool {
	var braces []string

	for _, v := range s {
		t := string(v)

		if t == "(" || t == "{" || t == "[" {
			braces = append(braces, t)
		} else {
			if len(braces) == 0 {
				return false
			}
			endingBrace := braces[len(braces)-1]
			if t == ")" && endingBrace == "(" {
				braces = braces[:len(braces)-1]
			} else if t == "}" && endingBrace == "{" {
				braces = braces[:len(braces)-1]
			} else if t == "]" && endingBrace == "[" {
				braces = braces[:len(braces)-1]
			}

		}
		if len(braces) == 0 {
			return true
		}
	}
	return false
}

func isValid1(s string) bool {
	stack := []rune{}
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if len(stack) != 0 && (r == ')' && stack[len(stack)-1] == '(' || r == ']' && stack[len(stack)-1] == '[' || r == '}' && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
