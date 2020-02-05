package main

import "fmt"

type person struct {
	firstname string
	lastname string
	age int
}

func (p person) speak()  {
	fmt.Println("Name : ", p.firstname,p.lastname)
	fmt.Println("Age  : ", p.age)
}

func main()  {
	person1 := person{
		firstname: "Rakesh",
		lastname:  "Pattanayak",
		age:       35,
	}

	person1.speak()

}