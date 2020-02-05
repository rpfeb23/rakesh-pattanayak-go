package main

import (
	"fmt"
)

func main()  {
	fmt.Println(add)
	fmt.Println(multiply)
	fmt.Println(math)
	// At place of calling just give function name which will serve as address of that function . No parameter or signature needed
	math(2,3,add)
}

func add(x, y int)  {
	fmt.Println(x + y  )
}

func multiply(x, y int){
	fmt.Println(x * y)
}

// Just give function signature not function name
// then assign an identifier lik a, b

func math(x int, y int, f func(a,b int))  {
	  fmt.Println("Pass to Math function :", x,y,f)
      f(x,y)
}
