package main

//Write a program that prints out the numbers 1 to 100 (inclusive). If the number is divisible by 3,
//print Crackle instead of the number. If it's divisible by 5, print Pop. If it's divisible by both 3 and 5, print CracklePop.

func main() {
	//cracklePop()
	n := int64(20)
	p := int64(7)
	result := pthFactor(n, p)
	println(result)
}

func cracklePop() {
	// loop through the numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		// If the number is a multiple of 3 and 5 print "CracklePop"
		if i%3 == 0 && i%5 == 0 {
			println("CracklePop")
			// if the number is a multiple of 3	print "Crackle"
		} else if i%3 == 0 {
			println("Crackle")
			// if the number is a multiple of 5 print "Pop"
		} else if i%5 == 0 {
			println("Pop")
			// for all other numbers print the number
		} else {
			println(i)
		}
	}
}
