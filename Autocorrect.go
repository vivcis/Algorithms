package main

import "strings"

func Autocorrect(text string) string {

	if strings.HasPrefix(text, "you") {
		return text
	}
	if strings.HasSuffix(text, "uuuu") {
		return strings.Trim(text, "uuuu")
	}
	if strings.HasSuffix(text, "youuuuu") {
		return strings.Replace(text, "youuuuu", "your client", 0)
	}

	return strings.Replace(text, "you", "your client", -1)
}

//You've been tasked with writing an autocorrect service for messages sent by your legal team. The messages which are sent to other lawyers need to be updated
//so that each message sent references the lawyer's client for clarity. To do this you need to replace all instances of "you" and its misspellings with "your client".
//Complete the function so that it takes a string and replaces all instances of "you", "youuu", or "u" (not case sensitive) with "your client" (always lower case).
//"youuuuu" with any number of "u" characters tacked onto the end should be replaced with "your client"
//"u" at the beginning, middle, or end of a string, but NOT part of a word
//"you" but NOT as part of another word like "young" or "Bayou"
func Autocorrect1(s string) string {
	return strings.ReplaceAll(s, "youuuuu", "your client")
}
