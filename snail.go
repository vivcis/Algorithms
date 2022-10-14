package main

//Snail on codewars
func Snail(snailMap [][]int) []int {
	result := []int{}

	if len(snailMap) == 0 {
		return result
	}

	var rowStart int
	rowEnd := len(snailMap) - 1
	var colStart int
	colEnd := len(snailMap[0]) - 1

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			result = append(result, snailMap[rowStart][i])
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			result = append(result, snailMap[i][colEnd])
		}
		colEnd--

		for i := colEnd; i >= colStart; i-- {
			result = append(result, snailMap[rowEnd][i])
		}
		rowEnd--

		for i := rowEnd; i >= rowStart; i-- {
			result = append(result, snailMap[i][colStart])
		}
		colStart++
	}

	return result
}
