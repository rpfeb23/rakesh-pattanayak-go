package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main()  {
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("GO Routines :",runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	fmt.Println("*************************")
	bar()
	fmt.Println("*************************")
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("GO Routines :",runtime.NumGoroutine())
    wg.Wait()
}

func foo()  {
	for i :=0; i<=5; i++{
		fmt.Println("foo : ",i)
	}
	wg.Done()
}

func bar()  {
	for i :=0; i<=5; i++{
		fmt.Println("bar : ",i)
	}
}