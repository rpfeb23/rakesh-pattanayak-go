package main

import "fmt"

var a int

type rakeshcreatedtype int
var b rakeshcreatedtype

func main()  {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// error not an identifier its type "int"
	//fmt.Println(rakeshcreatedtype)
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Below not possible as a is type "Int" b is "rakeshcreatedtype"
	//a = b

}