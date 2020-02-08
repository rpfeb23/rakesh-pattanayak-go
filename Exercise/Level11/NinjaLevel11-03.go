package main

import "fmt"

type customErr struct {
	erorcode int
}

func (ce customErr) Error() string  {
	fmt.Println(" I am inside the function that implements error Interface")
    return fmt.Sprintf("Custom Message : Error Code : %v", ce.erorcode)
}

func (ce customErr) Error2() string  {
	ce.erorcode = 10000 + ce.erorcode
	return fmt.Sprintf("Custom Message 2 : Error Code : %v", ce.erorcode)
}

func foo(e error){
	// e can only access info for the function it got implemented.
	// so Error2 will not be invoked

	// WHY this is INVOKING the Error() function but in Interfacedemo.go it prints the Struct itself not invoke the Method it implements
	fmt.Println(e)
}
func main()  {

	ce1 := customErr{erorcode:999}
	// customError Implements error Interface
	// so ce1 is of type  1) "customErr" 2) builtin "error"

	foo(ce1)

}