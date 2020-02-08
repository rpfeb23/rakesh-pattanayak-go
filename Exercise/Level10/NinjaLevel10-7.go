package main

import "fmt"

func main() {
	c := make(chan int)
	send(c)
	receive(c)
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		go func() {
			for j := 21; j <= 30; j++ {
				value := (i * 100) + j
				c <- value
			}
		}()
	}
}

func receive(c <-chan int) {
	fmt.Println("Enetered Receive")
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}
