package main

import (
	"fmt"
	"sort"
)

// How do you access the 4th element of an array or slice?
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[4])

	// What is the length  & capacity of a slice created using:
	a := make([]int, 3, 9)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// Given the array:
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	// what would x[2:5] give you?
	fmt.Println(x[2:5])

	/*. Write a program that finds the smallest number in this list:*/
	y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17}
	// Sort the slice in ascending order
	sort.Ints(y)
	smallestNumber := y[0]
	fmt.Println(smallestNumber)

}
