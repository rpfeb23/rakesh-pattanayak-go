package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
var wg sync.WaitGroup

func main()  {
	var counter int64
	gs := 10
	wg.Add(gs)
	for i:=0; i<gs; i++ {
		fmt.Println("Go Routine # : ", i)
		go func() {
		    atomic.AddInt64(&counter,1)
			wg.Done()
		}()
	}
    wg.Wait()
	fmt.Println("Counter : ", counter)
}