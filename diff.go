package main

import (
	"fmt"
	"strings"
)

func main() {
	//var name string
	//fmt.Printf("Hi, %v", name)
	//fmt.Println("My name is")
	//fmt.Print("Rock")
	//fmt.Print("hello world")
	//fmt.Printf("Today is the:")
	//fmt.Print("Enter your name here: ")
	//fmt.Scan(&name)
	//fmt.Printf("my name is %s", name)
	//fmt.Print("HEY ")
	//fmt.Print("WASSUP")
	//variable := fmt.Sprintf("S/N: %v            NAME: %s", 6, "smith")
	//fmt.Print(variable)
	fmt.Print(FibonacciRecursion(10))

	s := []string{"ade", "faith", "jimmy", "gina", "sarah", "lizzy"}
	fmt.Printf("Number\t%s\tName\n", "|")
	for i := 0; i < len(s); i++ {
		fmt.Println(strings.Repeat("-", 20))
		fmt.Printf("%v\t%s\t%s\n", i, "|", s[i])
	}

	fmt.Println(isValid("(]"))
}
