package main

import "fmt"

func main()  {
	var slice1 []int
	slice1 = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

    fmt.Println(slice1)
	fmt.Println(slice1[:5])
	fmt.Println(slice1[5:])
	fmt.Println(slice1[2:7])
	fmt.Println(slice1[1:6])
}
