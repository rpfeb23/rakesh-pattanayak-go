package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) speak()  {
	fmt.Println(" I can speak")
	fmt.Println("Name :", p.name, " Age : ", p.age)
}

type human interface {
	speak()
}

func saysomething(h human)  {
	h.speak()
}

func main()  {
	p := person{
		name: "Rakesh",
		age:  32,
	}
	p.speak()
	(&p).speak()


	//saysomething(p)  // This will not work
	saysomething(&p)
}