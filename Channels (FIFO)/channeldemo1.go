// This code will not Run. Channel Blocks

package main
import "fmt"
func main()  {
	c := make(chan int)
    fmt.Printf("Type of Channel : %T \n", c)
	fmt.Println("Value of Channel : ", c)
	c <- 42

	fmt.Println(<- c)


}