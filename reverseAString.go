package main

import "fmt"

func main() {
	reverse1("meal")
	// Reversing the string.
	str := "ate a meal"

	// returns the reversed string.
	strRev := reverse2(str)
	fmt.Println(str)
	fmt.Println(strRev)
}

func reverse1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse2(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
