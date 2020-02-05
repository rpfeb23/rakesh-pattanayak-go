package main

import "fmt"

func main()  {

	a := increment()
	fmt.Println("First Increment call for a : ", a)
	a = increment()
	fmt.Println("First Increment call for a : ", a)
	a = increment()
	fmt.Println("Third Increment call for a : ", a)

	y := incrementreturnfunc()
	fmt.Println("Address of function y : ", y," But y has it own address : ", &y)
	fmt.Println("Value of z after first func execution : ", y())
	fmt.Println("Value of z after second func execution : ", y())
	fmt.Println("Value of z after Third func execution : ", y())

	fmt.Println()
	y2 := incrementreturnfunc()
	fmt.Println("Address of function y2 : ", y2," But y2 has it own address : ", &y2)
	fmt.Println("Value of z after first func execution : ", y2())
	fmt.Println("Value of z after second func execution : ", y2())
	fmt.Println("Value of z after Third func execution : ", y2())


}

func increment() int {
	var x int
	x++
	return x
}

func incrementreturnfunc() func() int{
	var z int
	z++
	return func() int {
		z = z+2
		return z
	}
}
