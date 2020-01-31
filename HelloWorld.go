package main

import "fmt"

func main() {

	// Signature
	fmt.Println("hello World")
	fmt.Println("String 1", 42 , 1.0) //variadic empty interface

	bytes, error := fmt.Print("Rakesh")

	fmt.Println("Number of bytes :", bytes)
	fmt.Println("Error if any : ", error)
}
