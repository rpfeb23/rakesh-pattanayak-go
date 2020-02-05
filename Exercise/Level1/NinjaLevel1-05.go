package main

import "fmt"

type customtype int
var x customtype //x is of TYPE customtype with UNERLYING Type int
var y int
func main(){
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)

    y = int(x)
    fmt.Println(y)
    fmt.Printf("%T\n", y)

}
