package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// Write your code here
	if llist == nil {
		return llist
	}
	if position == int32(0) {
		return llist.next
	}
	current := llist
	var count int32
	for current != nil { //10,1, 2, 12 , 15
		if count == position-1 {
			current.next = current.next.next
			//newNode = current.next
		}
		count++
		current = current.next
	}
	return llist
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{data: data}
	if position == 0 {
		newNode.next = llist
		llist = newNode
	}
	var count int32
	current := llist
	for current != nil {
		if position == position-1 {
			newNode.next = current.next
			current.next = newNode
			break
		}
		count++
		current = current.next
	}

	return llist

}

//reverse a linked list
func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	var behind *SinglyLinkedListNode
	for llist != nil {
		next := llist.next
		llist.next = behind
		behind = llist
		llist = next
	}
	return behind

}

//Print in REVERSE
func reversePrint(llist *SinglyLinkedListNode) {
	if llist == nil {
		return
	}
	reversePrint(llist.next)
	fmt.Println(llist.data)
}

func insertNodeAtPosition2(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	// Write your code here
	newNode := &SinglyLinkedListNode{data: data}
	if position == 0 {
		newNode.next = llist
		llist = newNode
		return llist
	}
	var count int32
	current := llist
	for current.next != nil {
		if count == position-1 {
			newNode.next = current.next
			current.next = newNode
			break
		}
		count++
		current = current.next
	}
	return llist
}
