package main

import "fmt"

func main()  {
	b1 := bar()
	fmt.Println(b1)  // Returns Address of the function in memory
	fmt.Printf("%T\n", b1)
	fmt.Println(b1()) // b1 is function(0 parm int return) invoke with ()

	fmt.Println(bar()())
}


 // func bar() string{
 //  	return "Hello Bar"
 // }

func bar() func() int {
	return func() int{
		return 451
	}
}


