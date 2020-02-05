package main

import "fmt"

func main()  {
	var array1 [5]int
	array1 = [5]int{11,22,33,44,55}
	fmt.Printf("Type of Array %T\n",array1)

	for i, v := range array1{
		fmt.Println("Index ", i , " Has Value ",v)
	}
}
