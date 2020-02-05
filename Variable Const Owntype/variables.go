//This program have errors intentionally left to demo

package main

import "fmt"

a := 42  // will throw error

var b = 43  //Accesible in entire Package to all functions

//Declares a variable with IDENTIFIER as "c"
//The TYPE of the IDENTIFIER "c" is INT
//It assigns a ZERO VALUE of type INT to "c"
//     false for booleans,
//     0 for numeric types,
//     "" for strings, and
//     nil for pointers, functions, interfaces, slices, channels, and maps.

var c int //Accessible in entire Package to all functions

func main()  {
	x := 42
	var y = 43
	fmt.Println(x, y)
	fmt.Println(a, b)
	fmt.Println(c)
	foo()
}

func foo()  {
	fmt.Println(x, y, a, b) // Only b accessible in foo()
}
