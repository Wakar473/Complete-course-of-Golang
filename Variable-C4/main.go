package main

import "fmt"

func main (){
	// var a string = "Apple"
	// println(a)

	var a string
	a = "Apple"
	println (a)



	var x string
	x = "first"
	println(x)
	x = "second"
	println(x)


	//infact you can even do this
	var y string
	y = "one"
	println(y)
	y = y + " two"
	println(y)

	var p string ="hey"
	var q string = "hey"
	println(p==q)


	var m string ="hey"
	var n string = "buddy"
	println(m==n)

/*Notice the : before the = and that no type was speci-
fied. The type is not necessary because the Go compiler
is able to infer the type based on the literal value you
assign the variable. (Since you are assigning a string
literal, x is given the type string ) The compiler can
also do inference with the var statement:*/


fmt.Print("Enter a number: ")
var input float64
fmt.Scanf("%f", &input)
output := input * 2
fmt.Println(output)



	 j := 5
println (j)

println(k)

}
var k string = "Hello World"
/*Notice that I moved the variable outside of the main
function. This means that other functions can access
this variable: */

/*we can use global variable any where in the code.*/ 
