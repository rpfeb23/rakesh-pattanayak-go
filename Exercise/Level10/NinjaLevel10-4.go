package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	fmt.Println("q : ", q)
	c1 := gen(q)
	fmt.Println("c1 : " , c1)
	receive(c1, q)

	fmt.Println("about to exit")
}

func gen(q1 chan<- int) <-chan int {
	fmt.Println("q1 : ", q1)
	c := make(chan int)
	fmt.Println("c :",c)
	fmt.Println("c :", c)
	go func() {
		for i := 1; i <= 3; i++ {
			c <- i
		}
		for i := 25; i <= 27; i++ {
			q1 <- i
		}
		close(c)
		close(q1)
	}()

	return c
}

func receive(c2 <-chan int, q2 <-chan int) {
	fmt.Println("c2 : ",c2)
	fmt.Println("q2 :",q2)
     for {
		 select {
		 case v,ok := <-c2 :
		 	if ok {
				fmt.Println("---c---",v)
			}else {
				return
			}
		 case v,ok := <-q2 :
		 	if ok {
				fmt.Println("***q***",v)
			}else{
				return
			}
		}

	 }
}
