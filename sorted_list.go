package main

/*
rearrange a sorted list so that the first element is largest, 2nd element is smalles
3rd elements is 2nd largest and so on
*/

func sortedList(data []int) []int {
	var result []int
	

	i := 0
	j := len(data) -1
	for
	{
		if i > j {
			return result
		}

		if i < j {
			result = append(result, data[j])
			result = append(result, data[i])
			i++
			j--
		} else if i == j {
			result = append(result, data[i])
			return result

		}

	}

}
