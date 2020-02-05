package main

import (
	"fmt"
	"runtime"
)

func main()  {
    fmt.Println("OS   :", runtime.GOOS)
    fmt.Println("ARCH :", runtime.GOARCH)
    fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("GO Routines :",runtime.NumGoroutine())

    // There could be scenario where your go routine is running but main routine gets finished leaving unfinished work. Use MUTEX/ WAITGROUP from Sync package so as to ensure your go routine runs and main waits for its completion. Demo in next

	go foo()
	fmt.Println("*************************")
	bar()
	fmt.Println("*************************")
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("GO Routines :",runtime.NumGoroutine())

}

func foo()  {
	for i :=0; i<=5; i++{
		fmt.Println("foo : ",i)
	}
}

func bar()  {
	for i :=0; i<=5; i++{
		fmt.Println("bar : ",i)
	}
}