package main

//Divisible sum pairs
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here

	var (
		count int32
		i     int32
		j     int32
	)
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {

			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}
