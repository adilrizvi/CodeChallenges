package main

/*
evaluate the given postfix expression
*/

// evaluate returns the result of op1 op op2
func evaluate(op string, op1, op2 int) int {

	var result int

	switch op {
	case "*":
		result = op1 * op2
	case "-":
		result = op1 - op2
	case "+":
		result = op1 + op2
	case "/":
		result = op1 / op2
	}

	return result

}

func evaluatePostFix(data []interface{}) int {

	st := stack{top: -1}
	for _, item := range data {
		val, ok := item.(int)
		if !ok {
			op1, _ := st.pop()
			op2, _ := st.pop()

			result := evaluate(item.(string), op2, op1)
			st.push(result)
		} else {
			st.push(val)
		}
	}

	res, _ := st.topElement()

	return res

}
