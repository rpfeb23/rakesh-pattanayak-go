// functions are first-class citizen
// you can pass function as param - Call Back
// you can assign function to a variable - Func Expression
// you can return a function

package main

import "fmt"

func main()  {
	foo()
	fmt.Println("----------------------------------")
	bar("I am from Main")
	fmt.Println("----------------------------------")
	mainstring := "I am from Main"
	fmt.Println("Before Woo function call : ", mainstring)
	mainstring = woo(mainstring)
    fmt.Println("After Woo function call : ", mainstring)
	fmt.Println("----------------------------------")
	s, b := too(mainstring, 35)
	fmt.Println(s,"            " , b)
	fmt.Println("----------------------------------")
}

// func (r receiver) identifer(parameter) (return){...}

func foo()  {
	fmt.Println("Hello, foo")
}

// Everything is PASS BY VALUE

func bar(s string)  {
	fmt.Println("You passed to bar function : ", s)
	s = "I got changed in bar"
}

func woo(s string) (string) {
	fmt.Println("You passed to woo function : ", s)
	s = "I got changed in woo"
	return s
}

func too(s string, i int) (string, bool) {
	s = s+" Again changed in too"
	b := true
	return s, b
}
