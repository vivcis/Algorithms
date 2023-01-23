package main

// MathChallenge - have the function MathChallenge(num) take the num parameter being passed and return the string true if the parameter is a prime number,
//otherwise return the string false. The range will be between 1 and 2^16.
//e.g input: 19 output: true
//input 110 output: false
func MathChallenge(num int) string {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return "false"
		}
	}
	return "true"
}
