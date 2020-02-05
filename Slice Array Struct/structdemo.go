//Struct : {a,b,c}
//Slice of Struct : [{a,b,c} {d,e,f}]

package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {

	p1 := person{
		"Rakesh",
		35,
		}

	fmt.Println(p1, p1.age, p1.name)

	p2 := person{
		name: "Rajesh",
		age:  37,
	}

	fmt.Println(p2)

	type cars struct {
		brand	string
		doors   int
	}

	c1 := cars{
		brand: "Audi R8",
		doors: 2,
	}
	fmt.Println(c1, c1.brand, c1.doors)

}
