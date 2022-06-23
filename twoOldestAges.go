package main

import "math"

func main() {

}

//The two oldest ages function/method needs to be completed. It should take an array
//of numbers as its argument and return the two highest numbers within the array.
//The returned value should be an array in the format [second oldest age, oldest age].
//The order of the numbers passed in could be any order. The array will always include at least 2 items.
//If there are two or more oldest age, then return both of them in array format.
//E.G TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}

func TwoOldestAges(ages []int) [2]int {
	var result [2]int

	oldestAge := -math.MaxInt32
	secondOldest := oldestAge

	for _, age := range ages {
		if age > oldestAge {
			secondOldest = oldestAge
			oldestAge = age
		} else if age > secondOldest {
			secondOldest = age
		}
	}
	result[0] = secondOldest
	result[1] = oldestAge

	return result
}
