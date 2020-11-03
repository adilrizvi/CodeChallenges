package main

/*
sort the elements in the given stack
*/

func sortStack(st *stack) (stack, error) {

	if st.isEmpty() {
		return stack{}, nil
	}

	result := stack{top: -1}

	for {
		if st.isEmpty() {
			return result, nil
		}
		currentElement, _ := st.pop()

		if result.isEmpty() {
			result.push(currentElement)
		} else {
			// pop the result stack  till currentElement > resultTopElement
			for {
				resultTopElement, err := result.topElement()
				if err != nil {
					break
				}
				if currentElement > resultTopElement {
					poppedElement, _ := result.pop()
					st.push(poppedElement)

				} else {
					break
				}

			}
			result.push(currentElement)
		}

	}
}
