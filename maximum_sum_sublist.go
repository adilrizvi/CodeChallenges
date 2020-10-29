package main

/*
from a given list find the maximum sum sublist
*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxSumSublist(data []int) int {

	maxEndingHere := 0
	maxSoFar := 0

	for _, item := range data {

		maxEndingHere = max(0, maxEndingHere+item)
		maxSoFar = max(maxEndingHere, maxSoFar)

	}

	return maxSoFar

}
