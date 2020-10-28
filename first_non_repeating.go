package main

import "math"

/*
find first non repeating number from given array
 */

type info struct {
	val             int
	firstOccurrence int
	frequency       int
}

func firstNonRepeatingNumber(data []int) (int, error) {

	// construct frequency map (char, firstOccurrence, frequency) for the input data
	freqInfo := make(map[int]*info)

	for idx, item := range data {
		fi, ok := freqInfo[item]
		if !ok {
			// case: firstOccurrence
			freqInfo[item] = &info{
				val: item, firstOccurrence: idx, frequency: 1,
			}
		} else {
			// item already exists increase the frequency counter
			freqInfo[item].frequency = fi.frequency + 1
		}
	}

	firstNonRepeatingIndex := math.MaxInt64
	for _, itemInfo := range freqInfo {
		if itemInfo.frequency == 1 {
			if itemInfo.firstOccurrence < firstNonRepeatingIndex {
				firstNonRepeatingIndex = itemInfo.firstOccurrence
			}

		}

	}

	return data[firstNonRepeatingIndex], nil
}
