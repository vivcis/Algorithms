package main

//Migratory birds on hackerrank
func migratoryBirds(arr []int32) int32 {
	// Write your code here

	max1st := int32(0)
	mappedNum := make(map[int32]int32) // NILL
	for _, char := range arr {
		if char > max1st {
			max1st = char
		}
		value, ok := mappedNum[char]
		if ok {
			mappedNum[char] = value + 1
		} else {
			mappedNum[char] = 1
		}
	}
	// map[int32]int32{1:2, 2:1, 3:2, 4:3}
	var count, max int32
	for key, value := range mappedNum {
		if value > count {
			count = value
			max = key
		} else if value == count {
			if key < max {
				max = key
			}
		}
	}
	return max
}

//OR
func migratoryBirds2(arr []int32) int32 {
	// Write your code here
	mappedNum := make(map[int32]int32) // NILL
	for _, char := range arr {
		mappedNum[char]++
	}
	var count, max int32
	for key, value := range mappedNum {
		if value > count {
			count = value
			max = key
		} else if value == count {
			if key < max {
				max = key
			}
		}
	}
	return max
}
