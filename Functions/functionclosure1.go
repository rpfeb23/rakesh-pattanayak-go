package main

import "fmt"

var x int

func main()  {
	fmt.Println("Initial Value of x :", x)
	x++
	fmt.Println("x value after x++ inside main : ", x)
	foo()
}
func foo(){
	x++
	fmt.Println("x value after x ++ inside foo : ", x)
}
