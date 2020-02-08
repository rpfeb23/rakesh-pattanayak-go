package main

import "fmt"

func main()  {
	values := make(chan int)
	defer close(values)

	go sendvalues(values)
    fmt.Println(" Mid way ")
	value := <- values
	fmt.Printf("Type of value : %T  Type of Values : %T ", value, values)
	fmt.Println(values) // Does not Print Anything
	fmt.Println()
	fmt.Println(value)
	fmt.Println(<- values)
	fmt.Println(<- values)
}

func sendvalues(c chan int)  {
	fmt.Println("Start of send Values")
     c <- 42
     c <- 43
     c <- 44
     fmt.Println("End of send Values")
}