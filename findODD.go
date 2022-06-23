package main

func FindOdd(seq []int) int {
	result := 0
	for _, num := range seq {
		result ^= num
	}
	return result
}

func FindOdd1(seq []int) int {
	// your code here

	freqMap := make(map[int]int)
	for _, num := range seq {
		if v, ok := freqMap[num]; ok {
			freqMap[num] = v + 1
		} else {
			freqMap[num] = 1
		}
	}

	result := 0
	for k, v := range freqMap {
		if v%2 == 1 {
			result = k
			break
		}
	}
	return result
}
