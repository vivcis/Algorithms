package main

import "fmt"

// Modify the order of a linked list in the following pattern, adding the current node to the result list after every step:
// Start at the head
// Take two steps forward
// Take one step back
// Take three steps forward
// Go to step 3 unless outside end of list
// Add unvisited element at end of list to result, if any

//type Node struct {
//	Val int
//	Next *Node
//}

func (head *Node) Stagger() {
	var result []int
	var current *Node = head
	var next *Node
	var prev *Node
	var nextNext *Node
	var nextNextNext *Node

	for current != nil {
		next = current.Next
		nextNext = next.Next
		nextNextNext = nextNext.Next

		current.Next = prev
		prev = current
		current = nextNextNext
		result = append(result, next.Val)
	}
	for _, val := range result {
		fmt.Println(val)
	}
}

func (head *Node) Stagger1() {
	for head != nil {
		tmp := head.Next
		if tmp != nil && tmp.Next != nil {
			head.Next = head.Next.Next
			tmp.Next = head.Next.Next
			head.Next.Next = tmp
		}
		head = tmp
	}
}
