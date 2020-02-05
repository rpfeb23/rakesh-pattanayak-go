package main

import "fmt"

func main() {
	string1 := "HELLLLLLLO"
	b := foo
	fmt.Println("Address of foo : ", b)
	fmt.Printf("Type of b is : %T\n", b)
	b1 := b()
	fmt.Println("*** FOO Now Executed and Returned a Function to b1 ***")
	fmt.Println("Address of b1 : ", b1)
	fmt.Printf("Type of b is : %T\n", b1)

	fmt.Println("Executed b and Result is : ")
	fmt.Println(b1(string1))

	fmt.Println("Other way of execution : ")
	fmt.Println(b()(string1))
}

func foo() func(s string) int {
	return bar
}

func bar(s string) int {
	fmt.Println(s)
	return 100
}
