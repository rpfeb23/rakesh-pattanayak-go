package main

import "fmt"

type car struct {
	brand string
	price int
}

func main()  {
    c1 := car{
    	brand:"Audi",
    	price:35000,
	}
    c2 := car{
    	brand:"Merc",
    	price:36000,
	}

    slicecars := []car{}
    slicecars = append(slicecars, c1, c2)
    fmt.Println(slicecars)

    slicecars2 := []car{
    	car{
    		brand:"Honda",
    		price:20000,
		},
		car{
			brand:"Toyota",
			price:22000,
		},
	}
    fmt.Println(slicecars2)

    for _, v := range slicecars2{
    	fmt.Println(v)
	}
}
