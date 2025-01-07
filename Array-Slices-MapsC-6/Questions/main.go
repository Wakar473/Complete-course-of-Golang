package main

import (
	"fmt"
)

// How do you access the 4th element of an array or slice?
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[4])


// What is the length of a slice created using:
a:=make([]int, 3, 9) 
fmt.Println(len(a))
}
