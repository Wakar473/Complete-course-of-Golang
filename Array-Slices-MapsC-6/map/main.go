package main
import(
	"fmt"
)

func main() {
	 a := make(map[string]int)
	a["key"] = 10
	fmt.Println(a["key"])
/*
The map type is represented by the keyword map , fol-
lowed by the key type in brackets and finally the value
type. If you were to read this out loud you would say “ a
is a map of string s to int s.”
	*/

b := make(map[int]int)
b[1] = 10
delete(b,1)
// We can also delete items from a map using the built-in delete function:
 fmt.Println(b[1])
}