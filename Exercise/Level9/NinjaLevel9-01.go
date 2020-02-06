package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main()  {
    wg.Add(2)
	go routine1()
	go routine2()
	wg.Wait()
}
func routine1()  {
	fmt.Println("I am in routine 1")
	wg.Done()
}
func routine2()  {
	fmt.Println("I am in routine 2")
	wg.Done()
}