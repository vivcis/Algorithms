package main

//Number line jumps on hackerrank
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here

	if v1 > v2 {
		if (x2-x1)%(v1-v2) == 0 {
			return "YES"
		}
	}
	return "NO"
}
