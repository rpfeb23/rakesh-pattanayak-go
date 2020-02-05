package main

import "fmt"

func main()  {

	pos := adder()
	fmt.Println(pos) // OBSERVER CAREFULLY RETURN did not run anything so Print statement inside RETURN did not execute here

	fmt.Println("Execute pos function which takes int as param and returns int : ", pos(10))   // OBSERVER the order first pos(10) got executed then the Println . Thats why the Print statements are out of order

	fmt.Println("Execute pos function second time which takes int as param and returns int : ")
	fmt.Println("\t\t\t\t\t\t\t\t",pos(20))

	 fmt.Println()
	 fmt.Println()
	 fmt.Println(" A New Memory created for neg which has Sum = 5000 from the returned func. ")
     neg := adder()
     fmt.Println(neg(-1*10))
     fmt.Println(neg(-2*10))
}

func adder() func(i int) int  {
	sum :=5000
	fmt.Println("I am inside adder")
	return func(i int) int {
		fmt.Println("\tI am inside annonymous func which is getting returned")
		sum = sum + i
		return sum
	}
}
