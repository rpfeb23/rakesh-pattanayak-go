package main

import "fmt"

func main()  {
   x := func() int{
   	fmt.Println(" I am annonymous")
   	return 10
   }
   fmt.Println(x)  // Prints the address

   x()  // Thought it retruns I am not capturing here so it will just PRINT

	fmt.Printf("Type of x is : %T\n",x)

   fmt.Println(x()) // Now Println will capture the INT and Print 10
                    // Also it will Print I am annonymous


}