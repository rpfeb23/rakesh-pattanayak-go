package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main()  {
	var counter int
	gs := 10
	wg.Add(gs)
	for i:=0; i<gs; i++ {
		fmt.Println("Go Routine # : ", i)
		go func() {
			tempvalue := counter
			runtime.Gosched()
			tempvalue++
			counter = tempvalue
			wg.Done()
		}()
	}
    wg.Wait()
	fmt.Println("Counter : ", counter)
}