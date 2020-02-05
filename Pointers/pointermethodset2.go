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

func (c *circle) area()  {
	fmt.Println("Value Received :", c)
	fmt.Println("Circle area :", math.Pi*math.Pow(c.radius,2))
}

func (s *square) area()  {
	fmt.Println("SQUARE Area : ", s.length*s.width)
	s.width = 4
	s.length =4
	fmt.Println(s)
}
type shape interface {
	area()
}

func calcarea(sh shape)  {
	fmt.Println(sh)
	sh.area()
}

func main()  {

	c1 := circle{radius:1}
    fmt.Println(c1,&c1)
	fmt.Println("----------------")
	fmt.Printf("%T\t%T\n",c1,&c1)
	c1.area()
	(&c1).area() // You need a pointer to access area
	fmt.Println("----------------")
	calcarea(&c1)    // Receiver is  Pointer so it can only get Pointer
    // calcarea(c1) not possible
    s1 := square{
		length: 2,
		width:  2,
	}
    calcarea(&s1)
    fmt.Println("I have changed Length and Width after calc in area func." +
    	"lets see what comes out")
    fmt.Println(s1)
    calcarea(&s1)
}