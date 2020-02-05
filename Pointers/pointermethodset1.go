package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length int
	width int
}

func (c circle) area()  {
	fmt.Println("Circle area :", math.Pi*math.Pow(c.radius,2))
}

func (s square) area()  {
	fmt.Println("SQUARE Area : ", s.length*s.width)
}
type shape interface {
	area()
}

func calcarea(sh shape)  {
     sh.area()
}

func main()  {

	c1 := circle{radius:1}
	fmt.Println("After passing non-pointer to non-pointer receiver")
	calcarea(c1)

    fmt.Println("After passing pointer to non-pointer receiver")
	calcarea(&c1)
	fmt.Println()
	fmt.Println()
	s1 := square{
		length: 2,
		width:  2,
	}
	fmt.Println("After passing pointer to non-pointer receiver")
	calcarea(&s1)
	fmt.Println(s1)
	fmt.Println("After passing non-pointer to non-pointer receiver")
	calcarea(s1)
}