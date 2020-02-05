package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main()  {
	t1 := truck{
		vehicle:   vehicle{doors: 2,color:"white"},
		fourwheel: false,
	}

	fmt.Println(t1)
	fmt.Println(t1.vehicle.color,t1.doors,t1.fourwheel)

	c1 := sedan{
		vehicle: vehicle{doors:4,color:"black"},
		luxury:  true,
	}
	fmt.Println(c1)
	fmt.Println(c1.vehicle.doors,c1.color,c1.luxury)
}
