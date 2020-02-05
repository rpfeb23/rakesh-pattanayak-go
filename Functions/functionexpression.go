package main

import "fmt"

func main()  {
	f := func (){
		fmt.Println("I am blank")
	}
	f()

	f1 := func(i int) {
		fmt.Println(i)
	}
	f1(42)
}
