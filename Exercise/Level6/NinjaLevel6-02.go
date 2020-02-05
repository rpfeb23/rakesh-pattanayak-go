package main

import (
	"fmt"
)

func main()  {
	x := []int{1,2,3,4,5}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(i ...int) int {
	sum :=0
	for _,v := range i {
		sum += v
	}
	return sum
}

func bar(i []int) int {
	sum := 0
	for _,v := range i{
		sum = sum+v
	}
	return sum
}