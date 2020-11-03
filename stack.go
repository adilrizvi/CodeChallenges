package main

import (
	"errors"
)

var emptyStackErr error = errors.New("empty stack")

type stack struct {
	top int
	data []int
}

func (st *stack) push(val int) {
	st.top += 1
	st.data = append(st.data, val)
}

func (st *stack) topElement() (int, error){
	if st.isEmpty() {
		return -1, emptyStackErr
	}
	return st.data[st.top], nil

}

func (st *stack) pop() (int, error) {
	if st.isEmpty() {
		return -1, emptyStackErr
	}
	data := st.data[st.top]
	st.top -= 1
	st.data = st.data[:st.size()]
	return data, nil
}

func (st *stack) isEmpty() bool {
	return st.top == -1
}

func (st *stack) size() int {
	return st.top+1
}
