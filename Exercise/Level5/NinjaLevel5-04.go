package main

import "fmt"

func main()  {
	p1 := struct {
		name string
		age int
	}{
		name:"Rakesh",
		age:32,
	}
	fmt.Println(p1)
}
