package main

//have the function MathChallenge(num) return the string yes if the number given is a part of the Fibonacci sequence.
//This sequence is defined by: Fn = Fn-1 + Fn-2, which means to find Fn you add the previous two numbers up. The first two
//numbers are 0 and 1, then comes 1, 2, 3, 5 etc. If num is not in the Fibonacci sequence, return the string no.
//e.g input: 34 output: yes
//input 54 output: no
func MathChallenge1(num int) string {
	fib := []int{0, 1}
	for i := 2; i < num; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	for _, v := range fib {
		if v == num {
			return "yes"
		}
	}
	return "no"
}
