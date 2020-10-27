package main

import "fmt"

func main() {
	fmt.Println("inside main")
	fmt.Printf("result : %v\n", removeEvenNumbers([]int{1,2,3,5,6}))
	fmt.Printf("sorted merge: %v\n", sortedMerge([]int{-10,100,200}, []int{}))



}
