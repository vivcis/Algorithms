package main

//Number of integer partitions on codewars
//An integer partition of n is a weakly decreasing list of positive integers which sum to n.
//For example, there are 7 integer partitions of 5:
//[5], [4,1], [3,2], [3,1,1], [2,2,1], [2,1,1,1], [1,1,1,1,1].
func Partitions(n int) int {
	if n <= 0 {
		return 0
	}
	res := make([]int, n+1, n+1)
	res[0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			res[j] += res[j-i]
		}
	}
	return res[len(res)-1]
}

func Partitions1(n int) int {
	parts := make([]int, n+1)
	parts[0] = 1
	counter := 0
	for i := 1; i < n+1; i++ {
		counter = 0
		for j := i; j < n+1; j++ {
			parts[j] += parts[counter]
			counter++
		}
	}
	return parts[n]
}
