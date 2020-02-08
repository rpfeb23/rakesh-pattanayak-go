package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main()  {
	values := make(chan int)
	i := []int{42,43,44}
    wg.Add(2)
	go sendvalues(values, i)
	go receivevalues(values)
    wg.Wait()
}

func sendvalues(c chan<- int, i []int)  {
	fmt.Println(" Sender Go Routine STARTED ")
    for _, v := range i{      // Channels have to be closed for this to work
    	c <- v
    	fmt.Printf("Sent %v to the channel \n", v)
	}
	close(c)
    fmt.Println(" Sender Go Routine FINISHED ")
    wg.Done()
}
func receivevalues(c <-chan int)  {
	fmt.Println(" Receiver Go Routine STARTED ")
	for channelelement := range c{
		fmt.Println(channelelement)
	}
    wg.Done()
	fmt.Println(" Receiver Go Routine FINISHED ")
}