package main

import "container/list"

/*
implement 2 stacks with the same underlying list
 */




type MyStack struct {
	top1 int
	top2 int
	// solutions doesn't work for duplicate data in Stack
	data list.List
}

func (st *MyStack) push1(val interface{}) {
	st.top1 += 1
	st.data.PushFront(val)
}

func (st *MyStack) push2(val int) {
	st.top2 += 1
	st.data.PushBack(val)
}


func (st *MyStack) pop1() (interface{}, error) {
	if st.top1 == -1 {
		return -1, emptyStackErr
	}
	st.top1 -= 1
	data := st.data.Front()
	// remove the element from the data
	st.data.Remove(data)
	return data.Value, nil
}

func (st *MyStack) pop2() (interface{}, error) {
	if st.top2 == -1 {
		return -1, emptyStackErr
	}
	st.top2 -= 1
	data := st.data.Back()
	// remove the element from the data
	st.data.Remove(data)
	return data.Value, nil
}


