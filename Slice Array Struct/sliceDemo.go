package main

import "fmt"

func main()  {

	var x []int
    fmt.Println(x, len(x))
	x = []int{1,4,5}
	fmt.Println(x, len(x))

	for i,v := range x {
		fmt.Println("Index : ",i," Value : ",v)
	}
}
