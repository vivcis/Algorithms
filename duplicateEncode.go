package main

import "strings"

func DuplicateEncode(word string) string {
	var result []string
	//var w []string = result
	for i := 0; i < len(word); i++ {
		w := strings.ToLower(word)
		if strings.Count(strings.ToLower(word), string(w[i])) <= 1 {
			result = append(result, "(")
		}
		if strings.Count(strings.ToLower(word), string(w[i])) > 1 {
			result = append(result, ")")
		}
	}
	x := strings.Join(result, "")
	return x
}
