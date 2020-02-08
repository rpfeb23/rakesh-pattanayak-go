package main

import (
	"fmt"
)

func main()  {
	c := make(chan int)
	fmt.Println("c :",c)
	go send(c)
	receive(c)

}
func send(c1 chan<- int){
	fmt.Println("c1 :",c1)
	for i:=0; i<=100; i++{
		c1 <- i
	}
	close(c1)
}

func receive(c2 <-chan int)  {
	fmt.Println("c2 :",c2)
	for{
		if v, ok := <-c2 ; ok {
			fmt.Println(v)
		}else {
			return
		}

	}
}