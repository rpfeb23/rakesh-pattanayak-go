package main

import "fmt"

func main(){
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	var y [5]int
	fmt.Println(y)
	y = [5]int{1,2,3,4,5}
	fmt.Println(y)
}
