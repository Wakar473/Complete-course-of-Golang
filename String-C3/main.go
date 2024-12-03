package main
import(
)
func main(){
	println("Hello Wakar")
/*String Indexing: Strings in Go are treated as a sequence of bytes, not characters. Each character in the string is represented by its ASCII (or UTF-8 encoded) value.
Byte Value: "Hello World"[1] accesses the second byte (index 1) of the string "Hello World". The character at index 1 is e.
ASCII Code: The ASCII value of the character e is 101.
NOTE=>>Go automatically prints the byte value of a character when accessing it by index.*/
	println ("Hey wakar"[1])
    println (string("Hey wakar"[1]))
	println("hello " + "wakar")
	/*A space is also considered a character, so the
string's length is 11 not 10 and the 3 rd line has
"Hello " instead of "Hello" .*/
}