package main

func main() {

}

//You're given a two-dimensional array (a matrix) of distinct integers and a target integer. Each row in the matrix is sorted, and each column
//is also sorted; the matrix doesn't necessarily have the same height and width.
//Write a function that returns an array of the row and column indices of the target integer if it's contained in the matrix, otherwise [-1, -1].
//Example:
//matrix = [
//  [1, 4, 7, 12, 15, 1000],
//  [2, 5, 19, 31, 32, 1001],
//  [3, 8, 24, 33, 35, 1002],
//  [40, 41, 42, 44, 45, 1003],
//  [99, 100, 103, 106, 128, 1004],
//]
//target = 44
//Return [3, 3] since the target integer is contained in the matrix.
func SearchMatrix(matrix [][]int, target int) []int {
	if len(matrix) == 0 {
		return []int{-1, -1}
	}
	for i, row := range matrix {
		if len(row) == 0 {
			return []int{-1, -1}
		}
		for j, num := range row {
			if num == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
