package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		fmt.Println(" Put to channel C successful")
	}()

	fmt.Println("Inside Main ")
	fmt.Println(<-c)

	fmt.Println("-------------1---------------")
	c1 := make(chan int, 1)

	c1 <- 42

	go func() {
		fmt.Println("   2nd Way     ",<-c1)
	}()

	fmt.Println("-------------2---------------")
	c2 := make(chan int)
	go func() {
		for {
			fmt.Println(<-c2)
		}
	}()
	c2 <- 42
	c2 <- 43
	c2 <- 44
	fmt.Println("-------------3---------------")

}
