// pulling from channel will give i : zero value ok: false if the channel oes not have anything
package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	go send(even, odd, quit)
	receive(even, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//q <- true
	close(e)
	close(o)
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println(" From the Even channel : ", v)
		case v := <-o:
			fmt.Println(" From the ODD channel : ", v)
		case i, ok := <-q:
			if ok{
				fmt.Println(" From the Quit channel : ", i)
			}else{
				fmt.Println(" From the Quit channel NOT OK : ", i, ok)
			}
			return
		}
	}
}
