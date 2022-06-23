package main

import "fmt"

type Doubly struct {
	Val  int
	Next *Doubly
	Prev *Doubly
}

// {1} -> nil
func AppendDoubly(head *Doubly, val int) *Doubly {
	newNode := &Doubly{Val: val}
	if head == nil {
		head = newNode
		return head
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	newNode.Prev = current
	current.Next = newNode
	return head
}
func PrintDoublyList(head *Doubly) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
