package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
var x int
func main()  {

	wg.Add(2)
	go increment1()
	go increment2()
	wg.Wait()
}

func increment1()  {
	runtime.Gosched()
	for i := 0; i <5 ; i++{

		x = x+i
		fmt.Println("Increment1 : ", x)

	}
    wg.Done()
}

func increment2()  {
	for i := 0; i <5 ; i++{
		x = x+i
		fmt.Println("Increment2 : ", x)
	}
	wg.Done()
}