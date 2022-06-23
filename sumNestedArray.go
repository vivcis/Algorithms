package main

func main() {

}

//You've been tasked with writing a function that returns the sum of all the numbers in an array. However, the array can contain nested arrays.
//Write a function that will return the sum of all the numbers in an array, including any nested arrays.
//Example:
//arr = [1, 2, [3, 4], [5, [6]]]
//Return 21
func SumNestedArray(arr []interface{}) int {
	sum := 0
	for _, v := range arr {
		switch v.(type) {
		case []interface{}:
			sum += SumNestedArray(v.([]interface{}))
		case int:
			sum += v.(int)
		}
	}
	return sum
}
