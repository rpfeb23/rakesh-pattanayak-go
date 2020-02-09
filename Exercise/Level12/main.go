// This Package is to Calculate Dog Years
package main

import (
	"fmt"
	"github.com/rpfeb23/rakesh-pattanayak-go/Exercise/Level12/dog"
)

type doginfo struct {
	name string
	age  int
}

func main()  {
	humanyears := 2

	d1 := doginfo{
		name: "Oli",
		age:  dog.Dogyearsrakesh(humanyears),
	}

	fmt.Println(" First Dog : ", d1)

}