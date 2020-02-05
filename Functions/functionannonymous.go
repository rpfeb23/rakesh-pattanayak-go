package main

import "fmt"

func main()  {
	foo()

	// annonymous func. First () is your function param and second () is to invoke the function and pass param

	func (){
		fmt.Println("Annonymous Func")
	}()

	func (i int){
		fmt.Println("Meaning of life :", i)
	}(42)
}

func foo()  {
	fmt.Println("I am foo")
}
