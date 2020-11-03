package main

/*
reverse first k elements for the given queue
*/

func reverseKElements(q *queue, n int) (queue, error) {
	result := queue{}
	st := stack{top: -1}

	// dequeue n elements from queue and push those elements to
	// stack
	for i := 0; i < n; i++ {
		data, err := q.dequeue()
		if err != nil {
			return result, err
		}
		st.push(data.(int))

	}

	// enqueue all elements of stack into new result queue
	for i := 0; i < n; i++ {
		data, err := st.pop()
		if err != nil {
			return result, err
		}
		result.enqueue(data)
	}

	// enqueue the remaining elements of the original queue to the
	// result queue
	for {
		if q.isEmpty() {
			return result, nil
		}
		data, err := q.dequeue()
		if err != nil {
			return result, err

		}
		result.enqueue(data)
	}

	return result, nil

}
