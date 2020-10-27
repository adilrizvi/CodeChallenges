package main

/*
for a given array (of positive numbers), modify the slice such that for each element, we store the product of all elements except the current
element
 */


func productList(data []int) []int {
	// calculate all product
	if data == nil {
		return nil
	}
	prod := 1
	for _, item := range data {
		prod *= item
	}

	for idx, item := range data{
		data[idx] = prod/item
	}

	return data
}