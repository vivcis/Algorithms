package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println(hackerrankInString("cvbhavvcketrrrrrran"))
	//fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
	//fmt.Println(saveThePrisoner(5, 2, 1))
	//fmt.Println()

	//name := "Cecilia"
	//fmt.Printf("name: %v, %v\n\n", name, &name)
	//
	//myname := &name
	//fmt.Printf("%v\n", *myname)

	//l := []string{"Joseph", "Ebube", "Cecilia"}
	//for i, _ := range l {
	//	fmt.Println(&l[i])
	//}
	//change(l)
	//fmt.Println("address of l: %l")

	fmt.Println(Solution("abvcfdergt"))
	fmt.Println(superReducedString("aaabccddd"))
	//fmt.Println(staircase(6))

	s := ordinal(200357)

	fmt.Println(s)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2022", 2021-input)
}

func change(list []string) {
	//list[0] = "changed"
	list = append(list, "changed")
	fmt.Printf("address of list: %p\n\n", &list)
}
