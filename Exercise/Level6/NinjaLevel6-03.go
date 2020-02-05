package main

import "fmt"

func main()  {
	foo()
}

func foo()  {
	fmt.Println("I am entering foo")
	defer bar()
	fmt.Println("I am exiting foo")
}
func bar()  {
	fmt.Println(" I am in bar")
}