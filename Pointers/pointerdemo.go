package main

import "fmt"

func main()  {
	//  IDENTIFIER     TYPE   VALUE
	//    a             int   10
	//   &a            *int   x101010c(Address of a)
	a := 10
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*&a)

	fmt.Printf("Type of a is : %T\n",a)    // int is a TYPE
	fmt.Printf("Type of &a is : %T\n",&a)  // *int is a TYPE
	fmt.Printf("Type of *&a is : %T\n", *&a)
}