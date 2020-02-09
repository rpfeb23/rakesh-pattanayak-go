package main

import "fmt"

func main()  {
	fmt.Println(" Using Recur : ",factorial(-4))
	fmt.Println(" Using Loop  : ",factusingloop(-4))
}

func factorial(i int) int {
	var factorialOutput int = 1
	if (i == 0) {
		factorialOutput = 1
	} else if (i > 0){
		factorialOutput = i * factorial(i-1)
	}else{
		fmt.Println("Factorial of negetive number can not be calculated")
		return 0
	}
	return factorialOutput
}

func factusingloop(i int) int {

	factorial_output := 1

	if (i == 0){

	} else if (i>0){
		for ; i > 0; i-- {
			factorial_output = factorial_output * i
		}
	}else{
		fmt.Println("Factorial of negetive number can not be calculated")
		return 0
	}
	return factorial_output
}