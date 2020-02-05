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
		lastname:     "Pattnaik",
		favicecreams: []string{"Vannila","Strawberry"},
	}

	p2 := person{
		firstname:    "Rajesh",
		lastname:     "Pattanayak",
		favicecreams: []string{"Choc","Blackcurrent"},
	}

    m := map[string]person{
    	p1.lastname : p1,
    	p2.lastname : p2,
	}
    fmt.Println(m)

	for k,v := range m {
		fmt.Println(k,v)
		fmt.Println(v.lastname,v.firstname)
		for _,v1 := range v.favicecreams{
			fmt.Println(v1)
		}
	}

}
