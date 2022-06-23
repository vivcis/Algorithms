package main

import "fmt"

//bon appetit on hackerrank
func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here

	var annaVal int32
	var change int32
	var totalBill int32
	var dutch int32

	for i := 0; i < len(bill); i++ {
		totalBill = totalBill + bill[i]
		annaVal = bill[k]
	}
	dutch = (totalBill - annaVal) / 2
	if dutch != b {
		change = b - dutch
		fmt.Println(change)
	} else {
		fmt.Println("Bon Appetit")
	}
}
