package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	x := func(){
		fmt.Println("SUCCESS FUNCTION")
	}

	y := func(){
		fmt.Println("FAILURE FUNC")
	}

	fmt.Println(x,y)
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	fmt.Println(random)

	dosomething(random, x,y)
	
}

// Just give function signature not function name
// then assign an identifier lik a, b

func dosomething(random int, a func(), b func())  {
	    fmt.Println("Details passed to dosomething function :")
	    fmt.Printf("\t\t\t\t%v\t%v\t%v\n",random,a,b)
        if random % 2 == 0{
        	a()
		}else{
            b()
		}
}