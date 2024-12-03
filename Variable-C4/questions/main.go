// What are two ways to create a new variable?

package main

import "fmt"

func main() {
	var a string
	a = "Apple"
	println(a)

	b := "orange"
	println(b)

	var c string = "graps"
	println(c)


	// What is the value of x after running:x := 5; x += 1 ?
	x := 5
	x += 1
	println(x)

	y := 5
	y = y + 1
	println(y)



	//Using the example program as a starting point,write a program that converts from Fahrenheit into Celsius. ( C = (F - 32) * 5/9 )
	var fahrenheit float64
	println("enter fahrenhite")
	fmt.Scanf("%f", &fahrenheit)

	celcius:= fahrenheitToCelcius(fahrenheit)
	fmt.Printf("%2f is %f2f celcius\n",fahrenheit,celcius)


//Write another program that converts from feet into meters. ( 1 ft = 0.3048 m )
	
var feet float64
println("enter feet")
fmt.Scanf("%f", &feet)

meter:=feetToMeters(feet)
fmt.Printf(" %f meter\n", feet,meter)

}
func fahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9

}


func feetToMeters(feet float64) float64 {
	return   feet *0.3048 
}
