package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
var mu sync.Mutex
func main()  {
	var counter int
	gs := 10
	wg.Add(gs)
	for i:=0; i<gs; i++ {
		fmt.Println("Go Routine # : ", i)
		go func() {
			mu.Lock()
			tempvalue := counter
			tempvalue++
			counter = tempvalue
			mu.Unlock()
			wg.Done()
		}()
	}
    wg.Wait()
	fmt.Println("Counter : ", counter)
}