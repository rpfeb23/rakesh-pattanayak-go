package main

import "fmt"

func main()  {
	foo("Hmm",bar)
}

func foo(foostring string, f func(s string))  {
	fmt.Println(foostring)
	f("I am getting passed to bar from foo")
}

func bar(s string)  {
	fmt.Println(s)
}