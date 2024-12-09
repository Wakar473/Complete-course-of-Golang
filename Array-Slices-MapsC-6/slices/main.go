/* A slice is a segment of an array. Like arrays slices are
indexable and have a length. Unlike (not like) arrays this length
is allowed to change.*/

package main

import(
"fmt"
)

func main(){
	b := make([]float64, 4, 5)    //The make function  allows a 3rd parameter:
    a := make([]float64, 5)    //If you want to create a slice you should use the built-in make function:
println(a,b)
// ===================================================================================================
/*  Another way to create slices is to use the [low : high] expression:*/
arr := [] float64{1,2,3,4,5}
x := arr[0:5] //it count from index zero to element 5th i.e "5"
y:= arr[1:5]  //it count from index one to 5th element (element is 1,2,3,4,5 & index value is 0,1,2,3,4)
fmt.Println(x)
fmt.Println(y)


/* Go includes two built-in functions to assist with slices:
append and copy . Here is an example of append :*/

slice1:= []int {1,2,3}
slice2:= append(slice1, 4,5)
fmt.Println(slice2)

sli1:= [] int {1,2,3}
sli2:= make()

}