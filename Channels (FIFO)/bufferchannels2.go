// This will not run as only Buffer is there and we trying to put 2
package main
import "fmt"
func main()  {
	c := make(chan int, 1)
	c <- 42
	// fmt.Println(<- c) if you take out what you put in channel only 1 element will be left and then you can put one more. try uncommenting this line and see the result
	c <- 43
	fmt.Println(<- c)
}
