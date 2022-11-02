package main

import (
	"fmt"
	"time"
)

func main() {
	//fizzBuzz(10)
	//fizz()
	//buzz()
	//FIZZbuzz()

	fmt.Println("Main function")
	go countNumbers(20)

	fmt.Println("End main function")
}

//Write a short program that prints each number from 1 to 100 on a new line.
//For each multiple of 3, print "Fizz" instead of the number.
//For each multiple of 5, print "Buzz" instead of the number.
//For numbers which are multiples of both 3 and 5, print "FizzBuzz" instead of the number.
func fizz() {
	for i := 1; i < 10; i++ {
		s := fmt.Sprintf("%d", i)
		if i%3 == 0 {
			s = "Fizz"
			if i%5 == 0 {
				s += "Buzz"
			}
		} else if i%5 == 0 {
			s = "Buzz"
		}
		fmt.Println(s)
	}
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s = "Buzz"
		}
		//if i%3 == 0 && i%5 == 0 {
		//	s = "FizzBuzz"
		//}
		if s == "" {
			s = fmt.Sprintf("%d", i)
		}
		fmt.Printf("%s\n", s)
	}
}

func buzz() {
	for i := 1; i <= 10; i++ {
		s := ""
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s = "Buzz"
		}
		if s == "" {
			s = fmt.Sprintf("%d", i)
		}
		fmt.Printf("%s\n", s)
	}
}

func FIZZbuzz() {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func countNumbers(limit int) {
	time.Sleep(1 * time.Second)
	num := 0
	for i := 1; i < limit; i++ {
		num += i
	}
	fmt.Println("Num: ", num)
}
