package main

func main() {
	println(MathChallenge(19))
	println(MathChallenge(110))
}

// MathChallenge - have the function MathChallenge(num) take the num parameter being passed and return the string true if the parameter is a prime number,
//otherwise return the string false. The range will be between 1 and 2^16.
//e.g input: 19 output: true
//input 110 output: false
func MathChallenge(num int) string {
	// Check if `num` is divisible by any number between 2 and `num`-1
	for i := 2; i < num; i++ {
		// If `num` is divisible by `i`, return "false"
		if num%i == 0 {
			return "false"
		}
	}

	// If `num` is not divisible by any number between 2 and `num`-1, return "true"
	return "true"
}
