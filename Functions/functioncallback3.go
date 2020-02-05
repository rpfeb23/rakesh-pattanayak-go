package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	x := func() []int{
		fmt.Println("SUCCESS FUNCTION")
		r := []int{0}
		return r
	}

	y := func() []int{
		fmt.Println("FAILURE FUNC")
		return []int{1}
	}

	fmt.Println(x,y)
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	fmt.Println(random)

	dosomething(random, x,y)

}

func dosomething(random int, a func() []int, b func() []int)  {
	fmt.Println("Details passed to dosomething function :")
	fmt.Printf("\t\t\t\t%v\t%v\t%v\n",random,a,b)
	if random % 2 == 0{
		a()
		fmt.Println(a())
	}else{
		b()
		fmt.Println(b())
	}
}