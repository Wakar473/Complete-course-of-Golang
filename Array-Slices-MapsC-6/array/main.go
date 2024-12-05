package main

import("fmt")
/* An array is a numbered sequence of elements of a single type with a fixed length */
func main(){
var x [5] int
x[3] = 50 //(output)  [0 0 0 50 0] print thorugh index value
fmt.Println(x)

//================================================================================
var a [3] float64
a[0] =1
a[1] =2
a[2] =3

var total float64 = 0
for i := 0; i < 3; i++ {
	total = total+ a[i]   //it will also declare like that " total += a[i] "
fmt.Println(total)//    if i close "for" loop before println then it will print total no.and average of total.
	fmt.Println(total / 3)
}

//============================================================================
var b [3]int
    b[0] = 1
    b[1] = 2
    b[2] = 3

    var summ float64 = 0

    for i := 0; i < len(b); i++ {
        summ += float64(b[i]) // Convert b[i] to float64 before adding to summ
    }
/* This is an example of a type conversion. In general to
convert between types you use the type name like a
function.*/
    fmt.Println(summ / float64(len(b))) // Calculate and print the average

//==================================================================================
/*Another change to the program we can make is to use
a special form of the for loop:*/

var c [3] float64
c[0]=2
c[1]=2
c[2]=2

var d float64 =0
for _,value := range c{
	d += value
}
fmt.Println(d /float64(len(c)))
}
