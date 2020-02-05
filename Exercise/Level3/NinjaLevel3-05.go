package main

import "fmt"

func main()  {

	for i:=10;i<=100;i++{
		fmt.Printf("%v modulus 4 is : %v\n", i, i%4)
	}
}
