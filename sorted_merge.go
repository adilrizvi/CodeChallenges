package main

/*
merge 2 sorted slices such that result is also sorted
 */

func sortedMerge(data1 []int, data2 []int) []int {

	if data1 == nil {
		return data2
	}

	if data2 == nil {
		return data1
	}

	ctr1, ctr2 := 0, 0
	len1, len2 := len(data1), len(data2)

	var result []int

	for {
		if ctr1 < len1 && ctr2 < len2 {
			if data1[ctr1] < data2[ctr2] {
				result = append(result, data1[ctr1])
				ctr1++
			} else {
				result  = append(result, data2[ctr2])
				ctr2++
			}
		} else {
			break
		}
	}


	// one of the slices are consumed, henc we need to append
	// the rest of the bigger slice at the end of result
	if ctr1 < len1 {
		// data2 is exhausted, hence appending data1 to result
		for _, item := range data1[ctr1:] {
			result = append(result, item)
		}
	}

	if ctr2 < len2 {
		for _, item := range data2[ctr2:]{
			result = append(result, item)
		}
	}

	return result



}