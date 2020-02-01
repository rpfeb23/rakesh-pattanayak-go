package main

import "fmt"

var s string = "Hi, There"
var z string = `Rakesh said, "Do it by Monday"`
var i int
var b bool = true
func main()  {
    fmt.Println(s)
    fmt.Println(z)
	fmt.Printf("%T\n%T\n%T\n",s,i,b)
    fmt.Println(b)
    fmt.Printf("%t\n",b)

    c := 911 //Porshe 911

    fmt.Printf("%v\n",c)  //Value
	fmt.Printf("%b\n",c)  //Value in Binary
	fmt.Printf("%x\n",c)  //value in hex
	fmt.Printf("%#x\n",c) //value in hex with 0X infront

	sprintreturn := fmt.Sprint("Hi ","Its 2020")
	fmt.Println(sprintreturn)
}
