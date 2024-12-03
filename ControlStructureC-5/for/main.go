package main
import(
	// "fmt"
)

// func main (){
// 	i := 1
// 	for i<= 10{
// 		fmt.Println(i)
// 		i = i+1
// 	}
// }
/*First we create a variable called i that we use to store the number we want to print. Then we create a for
loop by using the keyword for , providing a conditional expression which is either true or false and finally
supplying a block to execute */
func main (){
	for i :=1; i<=10; i++{
		println(i)
	}
}