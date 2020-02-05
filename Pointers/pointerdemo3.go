package main

import (
	"fmt"
)

func main()  {
	a :=42
	foo(a)
	fmt.Println("After Foo call a is :,",a)
	bar(&a)
	fmt.Println("After Bar call a is : ", a ," a got changed")

}
func foo(a int)  {
	a++
	fmt.Println("Inside Foo a is :",a)
}
func bar(a *int)  {
	*a = *a +10
}