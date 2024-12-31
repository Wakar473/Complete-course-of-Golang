/*
map[KeyType]ValueType

KeyType: Keys ka data type (jaise string, int, etc.).
ValueType: Values ka data type (jaise int, string, float64, etc.).

using make function:
myMap := make(map[string]int)

myMap := map[string]int{
    "Alice": 25,
    "Bob": 30,
}



*/

package main

import (
	"fmt"
)

func main() {

	// Use Case: Jab aap map ko blank initialize karna chahte ho aur values baad me add karoge.
	a := make(map[string]int)
	a["key"] = 10
	a["key"] = 13
	fmt.Println(a["key"])
	/*
	   The map type is represented by the keyword map , fol-
	   lowed by the key type in brackets and finally the value
	   type. If you were to read this out loud
	   you would say “ 'a'
	   is a map of strings to ints.”
	*/

	b := make(map[int]int)
	b[1] = 10
	delete(b, 1)
	b[1] = 2
	// We can also delete items from a map using the built-in delete function:
	fmt.Println(b[1])
	//========================================================================================================================

	// Use Case: Jab aap map ko initialize karte waqt hi kuch values set karna chahte ho.
	myMap := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	myMap["charlie"] = 30
	delete(myMap, "Alice")
	fmt.Println(myMap)
	//========================================================================================================================

	x := map[string]string{
		"rahul": "4 year",

		"amit": "3 year",
	}
	// fmt.Println(x)
	x["rahul"] = "5 year"
	x["amit"] = "10 year"
	fmt.Println(x)
	//========================================================================================================================

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}

	name, ok := elements["Un"]
	fmt.Println(name, ok) // Output:  false

	/*    name, ok := elements["Un"]
	    fmt.Println(name, ok)
	Map ke andar "Un" key ko access karne ki koshish ho rahi hai.

	Agar "Un" key map ke andar exist karti hai, toh uski corresponding value return hogi.

	Agar "Un" key exist nahi karti, toh default zero value return hoti hai.
	2. name, ok :=:

	name: Yeh variable key "Un" ki value ko store karega (agar key exist karti hai).

	ok: Yeh ek boolean variable hai jo indicate karta hai ki key exist karti hai ya nahi.
	true: Key exist karti hai.
	false: Key exist nahi karti.
	*/
	//========================================================================================================================

	element := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}

	// Check for key "Un"
	if name, ok := element["Un"]; ok {
		fmt.Println(name, ok) // Yeh tabhi chalega jab key exist kare
	} else {
		fmt.Println("Key does not exist")
	}
	/*
		Iska use tab hota hai jab aapko map me kisi key ki existence ko check karte hue uski
		value retrieve karni ho, aur agar key exist nahi karti ho toh safe handling karni ho.
	*/
	//========================================================================================================================

	ele := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "Gas",
		},
		"He": {
			"name":  "Helium",
			"state": "Gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "Solid",
		},
	}

	// Accessing nested map
	fmt.Println(ele["H"]["name"])   // Output: Hydrogen
	fmt.Println(ele["He"]["state"]) // Output: Gas
	/*
		ye ek nested map ka example hai. Isme ek map ki values ke andar doosra
		map nest kiya gaya hai.

	*/

	ele["H"]["atomicNumber"] = "1"
	fmt.Println(ele["H"]) // Output: map[name:Hydrogen state:Gas atomicNumber:1]

	// Aap nested map ke andar naye keys aur values dynamically add kar sakte ho:

	//========================================================================================================================

	// Hierarchical structure using nested maps
	organization := map[string]map[string]map[string]string{
		"IT": {
			"Emp1": {
				"name": "Alice",
				"age":  "30",
			},
			"Emp2": {
				"name": "Bob",
				"age":  "25",
			},
		},
		"HR": {
			"Emp1": {
				"name": "Charlie",
				"age":  "35",
			},
			"Emp2": {
				"name": "Diana",
				"age":  "28",
			},
		},
	}

	// Accessing hierarchical data
	fmt.Println("IT Employee 1 Name:", organization["IT"]["Emp1"]["name"])
	fmt.Println("HR Employee 2 Age:", organization["HR"]["Emp2"]["age"])
}
