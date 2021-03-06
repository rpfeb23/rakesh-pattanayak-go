// Use go run -race racecondition.go to see if race condition exist
//  you will see at Bottom "Found 2 data race(s)
// which are x and y

package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
var x,y int
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
		y = y+i
		fmt.Println("Increment1 : ", x)

	}
    wg.Done()
}

func increment2()  {
	for i := 0; i <5 ; i++{
		x = x+i
		y = y+i
		fmt.Println("Increment2 : ", x)
	}
	wg.Done()
}