package main

import "fmt"

const c = 42

const (
	x = 43
	y string = "james bond"
)

func main()  {
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(x,"\t",y)
	fmt.Printf("%T\t%T\n", x, y)
}
