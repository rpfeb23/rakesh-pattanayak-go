package main

import "fmt"

type person struct {
   firstname string
   lastname string
}

func changeMe(a *person)  {
	a.lastname = "Bond"
	(*a).firstname = "James"
}
func main()  {
	p1 := person{
		firstname: "Rakesh",
		lastname:  "Pattanayak",
	}
	fmt.Println("Before changeMe Func call : ", p1)
	changeMe(&p1)
	fmt.Println("After changeMe Func call : ", p1)

}