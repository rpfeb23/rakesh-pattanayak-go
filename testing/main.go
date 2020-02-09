package main

import "fmt"

func main()  {
	fmt.Println(" Sum      : ",mysum(1,2,3))
	fmt.Println(" Multiply : ", mymultiply(1,2,3))
}

func mysum(a ...int) int {
	sum :=0
	for _, v := range a{
		sum -= v          // Intentionally kep - so as to fail Testing
	}
	return sum
}

func mymultiply(a ...int) int{
	multiplyresult :=1
	for _, v := range a{
		multiplyresult = multiplyresult+v // Intentionally kept + instead of *
	}
	return multiplyresult
}