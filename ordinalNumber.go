package main

import (
	"strconv"
	"strings"
)

func main() {

}

//Ordinal numbers{codewars} are used to tell the position of something in a list. Unlike regular numbers,
//they have a special suffix added to the end of them.
//Notation for brief notation
//Example
//1 - 1st
//2 - 2d
//3 - 3d
//5 - 5th
//11- 11th
//149 - 149th
//903 - 903d
func ordinal(number int) string {
	var x string
	if number == 0 {
		x = "0"
	}
	if number%100 <= 10 {
		if number%10 == 1 {
			num := strconv.Itoa(number)
			x = num + "st"
		} else if number%10 == 2 {
			num := strconv.Itoa(number)
			x = num + "nd"
		} else if number%10 == 3 {
			num := strconv.Itoa(number)
			x = num + "rd"
		} else {
			num := strconv.Itoa(number)
			x = num + "th"
		}
	}

	if number%100 > 10 && number%100 < 20 {
		num := strconv.Itoa(number)
		x = num + "th"
	}
	//20th, 21st, 22nd, 23rd, 24th, 25th
	if number%100 > 20 {
		if number%10 == 1 {
			num := strconv.Itoa(number)
			x = num + "st"
		} else if number%10 == 2 {
			num := strconv.Itoa(number)
			x = num + "nd"
		} else if number%10 == 3 {
			num := strconv.Itoa(number)
			x = num + "rd"
		} else {
			num := strconv.Itoa(number)
			x = num + "th"
		}
	}

	return x
}

//OR

func NumberToOrdinal(number int) string {
	num := strconv.Itoa(number)
	if num >= "11" && num <= "13" {
		return num + "th"
	} else if strings.HasSuffix(num, "1") {
		return num + "st"
	} else if strings.HasSuffix(num, "2") {
		return num + "nd"
	} else if strings.HasSuffix(num, "3") {
		return num + "rd"
	} else if num == "0" {
		return num
	}
	return num + "th"
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
