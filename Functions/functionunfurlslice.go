package main

import "fmt"

func main()  {

	x := []int{1,2,3}

    sum := variadic(x...)

    fmt.Println(sum)

    sum = variadic() // you can pass nothing
    fmt.Println(sum)
}

func variadic(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}