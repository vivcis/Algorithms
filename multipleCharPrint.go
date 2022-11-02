package main

import (
	"fmt"
	"strings"
)

//Create a program that prints the following (up to 25 characters):
//G GG GGG GGGG GGGGG GGGGGG GGGGGGG ...
//The program should print the G's on a single line, with a space between each G.
//use 2 nested for loops and not a string method
func g() {
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Print(" ")
	}
}

//using a string method
func Gs() {
	for i := 1; i <= 25; i++ {
		fmt.Printf("%s ", strings.Repeat("G", i))
	}
}
