package main

import "math"

func main() {
	n := int64(20)
	p := int64(3)
	result := pthFactor1(n, p)
	println(result)
}

//determine the factors of a number(i.e, all positive integer values that evenly divide into a number)
//and then return the pth element of the list, sorted ascending. if there is no pth element, return 0.
//example: n=20, p=3
//the factors of 20 in ascending order are (1, 2, 4, 5, 10, 20). Using 1-based indexing, if p = 3, then 4 is returned. if p > 6, 0 would be returned.
//function description: complete the function pthFactor
//pthFactor has the following parameters:
//long int n: the integer whose factors are to be found
//long int p: the index of the factor to be returned
//Returns:
//long int: the value of the pth integer factor of n or, if there is no factor at that index, then 0 is returned
//Constraints:
//1 <= n <= 10^15
//1 <= p <= 10^9

func pthFactor(n int64, p int64) int64 {
	var factors []int64
	for i := int64(1); i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	if int64(len(factors)) < p {
		return 0
	}
	return factors[p-1]
}

//OR

func pthFactor1(n int64, p int64) int64 {
	factors := []int64{}
	sqrtN := int64(math.Sqrt(float64(n)))

	for i := int64(1); i <= sqrtN; i++ {
		if n%i == 0 {
			factors = append(factors, i)

			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}

	sortFactors(factors)

	if p <= int64(len(factors)) {
		return factors[p-1]
	} else {
		return 0
	}
}

func sortFactors(factors []int64) {
	for i := 0; i < len(factors)-1; i++ {
		for j := 0; j < len(factors)-i-1; j++ {
			if factors[j] > factors[j+1] {
				factors[j], factors[j+1] = factors[j+1], factors[j]
			}
		}
	}
}
