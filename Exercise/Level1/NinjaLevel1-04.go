package main

import "fmt"

func main()  {
	type custometype int
	var x custometype

	fmt.Println(x)
	fmt.Printf("%T\n",x)

	x = custometype(42)
	fmt.Println(x)
}
