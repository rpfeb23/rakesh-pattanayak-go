package main

import "fmt"

func main()  {

	foo()
	defer too()
	defer bar()

}

func foo()  {
	fmt.Println("foo")
}

func bar()  {
	fmt.Println("bar")
}

func too()  {
	fmt.Println("too")
}
