package main

//Find The Parity Outlier on codewars
//[2, 4, 0, 100, 4, 11, 2602, 36]
//Should return: 11 (the only odd number)
//[160, 3, 1719, 19, 11, 13, -21]
//Should return: 160 (the only even number)

func FindOutlier(integers []int) int {
	odd_majority := 0
	even_majority := 0

	if integers[0]%2 == 0 {
		even_majority++
	} else {
		odd_majority++
	}
	if integers[1]%2 == 0 {
		even_majority++
	} else {
		odd_majority++
	}
	if integers[len(integers)-1]%2 == 0 {
		even_majority++
	} else {
		odd_majority++
	}
	result := 0
	if even_majority > odd_majority {
		//outlier should be just odd number
		for _, num := range integers {
			if num%2 != 0 {
				result = num
				break
			}
		}
	} else {
		//outlier is even no
		for _, num := range integers {
			if num%2 == 0 {
				result = num
				break
			}
		}
	}
	return result
}

func FindOutlier1(integers []int) int {
	var evens []int
	var odds []int

	for _, el := range integers {
		if el%2 == 0 {
			evens = append(evens, el)
		} else {
			odds = append(odds, el)
		}
	}

	if len(evens) > len(odds) {
		return odds[0]
	} else {
		return evens[0]
	}
}
