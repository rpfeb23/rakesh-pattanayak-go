package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := gen()
	go receive(c)

	fmt.Println("about to exit")
	wg.Wait()
}

func gen() <-chan int {
	c := make(chan int)
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}
