package main

import "fmt"

func main()  {
	a := 42
	fmt.Printf("Decimal : %d Binary : %b Hex : %#x \n",a,a,a)

	b := a << 1

	fmt.Printf("Decimal : %d Binary : %b Hex : %#x \n",b,b,b)
}
