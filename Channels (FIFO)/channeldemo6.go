// This will fail there is a deadlock in receiving method ue to for loop
package main

import "fmt"

func main()  {
	values := make(chan int)
	defer close(values)

	i := []int{42,43,44,45,46,47,48,49,50}

	go sendvalues(values, i)
	go receivevalues(values)

	for v := range values {
		fmt.Println(v)
	}
}

func sendvalues(c chan int, i []int)  {

	fmt.Println(" Go Routine STARTED ")
    for _, v := range i{
    	c <- v
	}
    fmt.Println(" Go Routine FINISHED ")
}
func receivevalues(c chan int)  {
	for channelelement := range c{
		fmt.Println(channelelement)
	}

}