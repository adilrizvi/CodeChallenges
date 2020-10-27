package main

/*
remove the even numbers from the given slice
 */

func removeEvenNumbers(input []int) []int {

	var result []int
	for _, item := range input {
		if item%2 != 0 {
			result = append(result, item)
		}
	}

	return result


}
