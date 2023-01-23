package main

//have the function ArrayChallenge(arr) take the array of  positive integers stored in arr and return the length of the longest
//increasing subsequence(LIS). A LIS is a subset of the original list where the numbers are in sorted order, from lowest to highest,
//and are increasing in order. The sequence does not need to be contiguous or unique, and there can be several different subsequences.
//e.g if arr is [4, 3, 5, 1, 6] then a possible LIS is [3, 5, 6], and another is [1, 6]. For this input, your program should return 3
//because that is the length of the longest LIS.

func ArrayChallenge(arr []int) int {

	lenNums := len(arr)
	lis := make([]int, lenNums)

	for i := 0; i < lenNums; i++ {
		lis[i] = 1
	}

	for i := 1; i < lenNums; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && lis[i] < (lis[j]+1) {
				lis[i] = lis[j] + 1
			}
		}
	}

	max := 0

	for i := 0; i < lenNums; i++ {
		if lis[i] > max {
			max = lis[i]
		}
	}

	return max
}
