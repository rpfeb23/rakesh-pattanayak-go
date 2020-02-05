package main

import "fmt"

func main()  {
	var x = []int{10,20,30,40,50}
	fmt.Println(x)
	// delete 3rd element 30

	x = append(x[:2],x[3:]...)
	fmt.Println(x)
}
