package main

import "math/bits"

//Write a function that takes an integer as input, and returns the number of bits that are equal to
//one in the binary representation of that number. You can guarantee that input is non-negative.
//Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

//Bit Counting on codewars
func CountBits(num uint) int {
	if num == 0 {
		return 0
	}
	var counter uint
	for num > 0 {
		counter += num % 2
		num /= 2
	}
	return int(counter)
}

func CountBits1(n uint) int {
	return bits.OnesCount(n)
}

//OR
var CountBits2 = bits.OnesCount

func CountBits3(b uint) int {
	bits := 0
	for b > 0 {
		b = b & (b - 1)
		bits++
	}
	return bits
}
