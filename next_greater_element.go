package main

import "fmt"

/*
find the next greater element on the right side of the given element
*/


func nextGreaterElement(data []int) {

	result := make(map[int]int)

	st := stack{top: -1}

	if len(data) <= 1 {
		fmt.Println(data)
		return
	}

	for _, item := range data {
		if st.isEmpty() {
			st.push(item)
			continue
		}

		if topEl, _ := st.topElement(); item > topEl {

			// pop the elements from stack till item <= topEl
			for {
				if st.isEmpty() || item <= topEl {
					break
				} else {
					result[topEl] = item
					st.pop()
					topEl, _ = st.topElement()
				}
			}

		}
		st.push(item)

	}

	fmt.Println(result)

}
