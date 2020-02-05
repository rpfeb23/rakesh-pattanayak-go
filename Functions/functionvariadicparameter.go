package main

import "fmt"

func main()  {
	sum1 := variadic(1,2,3)
	fmt.Println(sum1)
	sum2 := variadic(99,100)
	fmt.Println(sum2)
}
// variadic parameter is SLICE of Type
func variadic(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	sum :=0
	for _,v := range x{
		sum = sum + v
	}
	return sum
}
