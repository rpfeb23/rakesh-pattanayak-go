package main

import "fmt"

type person struct {
	firstname string
	lastname string
	favicecreams []string
}

func main()  {
	p1 := person{
		firstname:    "Rakesh",
		lastname:     "Pattanayak",
		favicecreams: []string{"Vannila","Strawberry"},
	}

	p2 := person{
		firstname:    "Rajesh",
		lastname:     "Pattanayak",
		favicecreams: []string{"Choc","Blackcurrent"},
	}

	for _, v := range p1.favicecreams{
		fmt.Println(v)
	}

	for _, v := range p2.favicecreams{
		fmt.Println(v)
	}
}
