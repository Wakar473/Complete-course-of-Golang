package main

import (
	"fmt"
)

/* An array is a numbered sequence of elements of a single type with a fixed length */
func main() {
	var x [5]int
	x[3] = 50 //(output)  [0 0 0 50 0] print thorugh index value
	fmt.Println(x)

	// =====================================================================================================================
	var a [3]float64
	a[0] = 1
	a[1] = 2
	a[2] = 3

	var total float64 = 0
	for i := 0; i < 3; i++ {
		total = total + a[i] //it will also declare like that " total += a[i] "
		fmt.Println(total)   //    if i close "for" loop before println then it will print total no.and average of total.
		fmt.Println(total / 3)
	}

	// ========================================================================================================================
	var b [3]int
	b[0] = 1
	b[1] = 2
	b[2] = 3

	var summ float64 = 0

	for i := 0; i < len(b); i++ {
		summ += float64(b[i]) // Convert b[i] to float64 before adding to summ
	}

	/*
	   This is an example of a type conversion*. In general to
	   convert between types you use the type name like a
	   function.

	   Note---Here we can see that for loop is closed before the print statment thats why it will print average of all no only once */

	fmt.Println(summ / float64(len(b))) // Calculate and print the average

	//============================================================================================================================
	/*Another change to the program we can make is to use
	  a special form of the for loop:*/

	var c [3]float64
	c[0] = 2
	c[1] = 2
	c[2] = 2

	var d float64 = 0
	for _, value := range c { //A single _ (underscore) is used to tell the compiler that we don't need this.
		// (In this case we don't need the iterator variable)

		// The range keyword is used to iterate over a collection such as arrays, slices, maps, or strings.[3]element 3 (0,1,2,3)
		d += value //Or d = d + value

		// value stores the current element of the array.

	}
	fmt.Println(d / float64(len(c)))

	// ============================================================================================================================

	// Go also provides a shorter syntax for creating arrays:

	e := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(e)

	// Sometimes arrays like this can get too long to fit on one line, so Go allows you to break it up like this:
	f := [5]float64{
		1,
		2,
		3,
		4,
		// 5,
		// After commenting '5' it will give output at element 5 index 4 zero because there is no value at index 4 and array of element[5]
	}
	fmt.Println(f)
}
