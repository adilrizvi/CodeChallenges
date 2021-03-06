package main

import (
	"fmt"
)

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

	var prev *node
	for current := ll.head; ; {
		if current == nil {
			return
		}
		if current.data == val {
			fmt.Printf("deleting: %d\n", val)
			if prev == nil {
				//case : first element duplicated
				ll.deleteAtHead()
				// restart the process
				current = ll.head
			} else {
				prev.next = current.next
				current = current.next
				//TODO: free memory for deleted node
			}

		} else {
			prev = current
			current = current.next
		}

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

func length(ll linkedList) int {
	var ctr int

	for temp := ll.head; ; {
		if temp == nil {
			return ctr
		}
		temp = temp.next
		ctr++
	}
}

func findMiddleNode(ll linkedList) (int, error) {

	if ll.isEmpty() {
		return 0, fmt.Errorf("Empty List")
	}

	if length(ll) <= 2 {
		return ll.head.data, nil
	}

	// case length(ll) > 2
	slow := ll.head
	fast := ll.head.next

	for {
		if fast == nil || fast.next == nil {
			return slow.data, nil
		}
		slow = slow.next
		fast = fast.next.next

	}

}

func reverse(ll linkedList) *node {

	if ll.isEmpty() {
		fmt.Println("Empty List")
		return nil
	}

	var prev, next, current *node
	current = ll.head
	for {
		if current == nil {
			//ll.head = prev
			return prev
		}
		next = current.next
		current.next = prev
		prev = current
		current = next

	}

	return prev

}

func (ll *linkedList) deleteAtKIndex(index int) {

	if ll.isEmpty() {
		// TODO: possibly return error
		fmt.Println("Empty List")
		return
	}

	if (index + 1) > length(*ll) {
		fmt.Println("Index (0 based) greater than length")
		return

	}

	var current, prev *node
	var ctr int
	current = ll.head

	for {
		if current == nil {
			return
		}

		if ctr == index {
			prev.next = current.next
			current = nil
			return
		}

		prev = current
		current = current.next
		ctr++

	}

}

func (ll *linkedList) removeDuplicates() {

	for current := ll.head; ; {

		if current == nil {
			return
		}

		// deleting the duplicates from remaining linkedList
		prev := current
		for temp := current.next; ; {
			if temp == nil {
				break
			}
			if temp.data == current.data {
				prev.next = temp.next
			} else {
				prev = temp
			}
			temp = temp.next
		}

		current = current.next

	}
}


func union(ll1 linkedList, ll2 linkedList) linkedList {
	ll := linkedList{}

	valMap := make(map[int]int)

	if ll1.isEmpty() {
		return ll2
	}
	 if ll2.isEmpty() {
	 	return ll1
	 }

	 // processing ll1
	 for current:=ll1.head;; {
	 	if current == nil {
	 		break
		}
		_, ok := valMap[current.data]
		if !ok {
			valMap[current.data] = 1
		}
		current = current.next
	 }

	 // processing ll2
	for current:=ll2.head;; {
		if current == nil {
			break
		}
		_, ok := valMap[current.data]
		if !ok {
			valMap[current.data] = 1
		}
		current = current.next
	}


	for val, _ := range valMap{
		ll.insertAtHead(val)
	}

	return ll


}


func intersection(ll1 linkedList, ll2 linkedList) linkedList {
	ll := linkedList{}

	valMap := make(map[int]int)

	if ll1.isEmpty() || ll2.isEmpty(){
		return ll
	}

	// processing ll1
	for current:=ll1.head;; {
		if current == nil {
			break
		}
		_, ok := valMap[current.data]
		if !ok {
			valMap[current.data] = 1
		}
		current = current.next
	}

	// processing ll2 and performing intersection
	for current:=ll2.head;; {
		if current == nil {
			break
		}
		_, ok := valMap[current.data]
		if ok {
			valMap[current.data] = 2
		}
		current = current.next
	}


	for val, ctr := range valMap{
		if ctr == 2{
			ll.insertAtHead(val)
		}

	}

	return ll


}


func (ll *linkedList) nthNodeFromEnd(n int) int {
	if ll.isEmpty() {
		return -1
	}

	if length(*ll) == 1 {
		return ll.head.data
	}

	var ctr int
	llLength := length(*ll)
	for current:=ll.head ;; {
		if ctr == llLength-n {
			return current.data
		}
		current = current.next
		ctr++
	}


}

func detectLoop(ll linkedList) bool {
	visitedMap := make(map[*node]bool)

	for current:=ll.head;;{
		if current == nil {
			return false
		}

		_, ok := visitedMap[current]
		if ok {
			return true

		} else{
			visitedMap[current] = true

		}
		current = current.next
	}
}
