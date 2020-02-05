package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	p1 := person{
		name: "Rakesh",
		age:  35,
	}

	fmt.Println("Person Struct at begining : ", p1)

	foo(&p1)

    fmt.Println("Person Struct after foo call : ", p1)
}

func foo(p *person)  {
	fmt.Println("Entire Person Passed to foo is an ADDRESS : ", p)
	fmt.Println("Person Passed to foo Value at that Address: ", *p)
	fmt.Println("Name : ",p.name,"  Age : ",p.age)
	p.name = "James Bond"
	(*p).age  = 40   //another way to do
}