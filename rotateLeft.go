package main

func main() {

}

func rotateLeft(d int32, arr []int32) []int32 {
	//iterate through the array
	//it should divide the no at the point of d index
	//return the nos after the division and then append the       rest
	var result []int32

	for i := int32(0); i < int32(len(arr)); i++ {
		//s := int32(i)
		if i+d < int32(len(arr)) {
			result = append(result, arr[i+d])
		} else {
			result = append(result, arr[i+d-int32(len(arr))])
		}
	}
	return result
}
