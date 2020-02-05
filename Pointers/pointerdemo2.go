package main

import "fmt"

func main()  {

	//  IDENTIFIER     TYPE   VALUE
	//    a             int   10
	//   &a            *int   x101010c(Address of a)
	//  *&a             int   10  (Value stored in the adress)
	//    b            *int   Address of a = Pointer
	//
	//   *b             int   Value stored in the address

	// & gives you address
	// * gives value stored in that address

	var a int = 10
	fmt.Println("Value of a :", a)
	fmt.Println("Address of a :",&a)

	var b *int
	fmt.Println(b)
	b = &a           // Pointer is &(Variable Name)
	fmt.Println(b)
	fmt.Println(*b)  // Deference the address
	                 // Value of Pointer *(Pointer Name)
	                 //  *int is a Type
	fmt.Printf("Type of b :%T\n",b)
	fmt.Printf("Type of *b :%T\n",*b)

	*b = 43
	fmt.Println("VALUE CHANGD Via b and see what happens to b and a")
	fmt.Println(a)
	fmt.Println(*b)

	fmt.Println("******* Just for Fun Start ********")
	fmt.Println(&b)  // Address of Pointer
	fmt.Println(*&b) // Value stored in the Pointer (Should be another Address)

	fmt.Println("******* Just for Fun End ********")



}