package main

import "fmt"

/*
given a list and number k, find 2 numbers from the list that sum to k
*/

func kSumPair(k int, data []int) (int, int, error) {
	// create map for all the items in the slice for flaster
	// lookup
	dataMap := make(map[int]bool)
	for _, item := range data {
		dataMap[item] = true
	}

	// for each item search k-item in map
	for _, item := range data {
		_, ok := dataMap[k-item]
		if !ok {
			continue
		} else {
			return item, k-item, nil
		}
	}

	return 0,0, fmt.Errorf("no %d, pair sum present", k)

}
