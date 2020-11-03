package main

import "fmt"

func generateBinaryNumbers(n int) queue {
	q := queue{}

	if n == 1 {
		q.enqueue("1")
		return q
	}

	for {
		if n <= 0 {
			return q
		}
		frontElement, err := q.dequeue()
		if err != nil {
			q.enqueue("1")
			n -= 1
			continue
		}
		q.enqueue(fmt.Sprintf("%s0", frontElement))
		n -=1
		if n > 0 {
			q.enqueue(fmt.Sprintf("%s1", frontElement))
			n -=1
		}
		q.enqueue(frontElement)

	}

	return q
}
