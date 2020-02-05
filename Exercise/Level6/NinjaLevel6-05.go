package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width float64
}

func (s SQUARE) area()  float64{
	return (s.width*s.length)
}
type CIRCLE struct {
	radius float64
}

func (c CIRCLE) area() float64 {
	return (math.Pi*math.Pow(c.radius,2))
}
type SHAPE interface {
	area() float64   // signature has to match with types implementing
}

func INFO(sh SHAPE)  {
	var shapearea float64
	switch sh.(type) {
	case SQUARE:
		fmt.Println("SHAPE Type : SQUARE")
	case CIRCLE:
		fmt.Println("SHAPE Type : CIRCLE")
	}
	shapearea= sh.area()
	fmt.Println("SHAPE AREA IS : ", shapearea)
}

func main()  {
	s := SQUARE{
		length: 2,
		width:  3,
	}
	c := CIRCLE{radius:1}
	fmt.Println(s.area())
	fmt.Println(c.area())

	fmt.Println("*****************************")
	INFO(s)
	INFO(c)
}