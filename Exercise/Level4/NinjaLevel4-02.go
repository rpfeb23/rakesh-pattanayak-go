package main

import "fmt"

func main()  {
	var slice1 []int
	slice1 = []int{10,20,30,40,50,60,70,80,90,100}
	fmt.Printf("Type of Slice %T\n",slice1)

	for i, v := range slice1{
		fmt.Println("Index ", i , " Has Value ",v)
	}
}
