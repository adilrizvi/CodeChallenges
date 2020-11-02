package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (ll *linkedList) getHead() *node {
	return ll.head
}

func (ll *linkedList) insertAtTail(val int) {
	if ll.isEmpty() {
		ll.insertAtHead(val)
		return
	}

	toInsert := &node{data: val, next: nil}

	// finding the tail of linkedList
	temp := ll.head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// insertion at tail
	temp.next = toInsert

}

func (ll *linkedList) insertAtHead(val int) {
	// create a node
	temp := &node{data: val, next: nil}

	temp.next = ll.head
	ll.head = temp
}

func (ll *linkedList) search(val int) bool {
	for temp := ll.head; ; {
		if temp == nil {
			return false
		}
		if temp.data == val {
			return true
		}
		temp = temp.next

	}

	return false
}

func (ll *linkedList) isEmpty() bool {
	return ll.head == nil
}

func printLinkedList(ll linkedList) {
	if ll.isEmpty() {
		fmt.Println("Empty List")
		return
	}

	temp := ll.head
	for {
		if temp == nil {
			fmt.Println()
			return
		}
		fmt.Printf("%d->", temp.data)
		temp = temp.next
	}
}

func (ll *linkedList) delete(val int) {
	if ll.isEmpty() {
		fmt.Println("Nothing to delete")
		return
	}

	if ll.head.data == val {
		ll.deleteAtHead()
		return
	}

	temp := ll.head
	var prev *node
	for {
		if temp == nil {
			return
		}
		if temp.data == val {
			fmt.Printf("deleting: %d\n", val)
			prev.next = temp.next
			temp = nil
			return
		}
		prev = temp
		temp = temp.next
	}

}

func (ll *linkedList) deleteAtHead() {
	if ll.isEmpty() {
		fmt.Println("Nothing to delete")
		return
	}

	temp := ll.head
	ll.head = ll.head.next

	// free memory
	temp.next = nil
	temp = nil

}
