package main

import"fmt"

func main(){
	for i:= 1; i<=10; i++{
		if i%2 ==0{
			fmt.Println(i,"even")

		}else{
			fmt.Println(i,"odd")
			
		}
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
		}
}}
/*A switch statement starts with the keyword switch fol-
lowed by an expression (in this case i ) and then a se-
ries of case s. The value of the expression is compared
to the expression following each case keyword. If they
are equivalent then the statement(s) following the : is
executed */