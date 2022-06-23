package main

func main() {

}

func saveThePrisoner1(n int32, m int32, s int32) int32 {
	// Write your code here
	var result int32
	for i := s; i < m+s; i++ {
		result = i % n
	}
	return result
}

//OR

func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	var result int32
	var zero int32
	for i := zero; i < m; i++ {
		if s > n {
			s = 1
		}

		result = s
		s++
	}
	return result

}
