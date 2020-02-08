package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	receive(even, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
	close(e)
	close(o)
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println(" From the Even channel : ", v)
		case v := <-o:
			fmt.Println(" From the ODD channel : ", v)
		case v := <-q:
			fmt.Println(" From the Quit channel : ", v)
			return
		}
	}
}
