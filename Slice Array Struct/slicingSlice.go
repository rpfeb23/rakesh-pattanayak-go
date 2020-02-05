package main

import "fmt"

func main()  {

	var x = []string{"A", "B", "C", "D", "E"}
	fmt.Println("Entire Slice : ", x)
	fmt.Println("Entire Slice [:]: ", x[:])
	fmt.Println("Slice a Slice [0:] : ", x[0:])
	fmt.Println("Slice a Slice [3:] : ", x[3:])
	fmt.Println("Slice a Slice [:0] : ", x[:0])
	fmt.Println("Slice a Slice [:3] : ", x[:3])
	fmt.Println("Slice a Slice [1:3] : ", x[1:3])

	for i:=0; i<len(x);i++{
		fmt.Println(x[i])
	}
}
